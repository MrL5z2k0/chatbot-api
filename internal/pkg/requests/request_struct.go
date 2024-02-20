/**
 * @Author haochen
 * @Email: lizekun01@bilibili.com
 * @Date: 2024/2/20 21:00
 * @Description: 描述
**/

package requests

import (
	"time"

	xtime "go-common/library/time"

	"github.com/go-resty/resty/v2"

	"helloworld/ecode"
)

type Scheme string
type ClientParam func(o *ClientConfig)

const (
	HTTP                 Scheme = "http"
	HTTPs                Scheme = "https"
	Discovery            Scheme = "discovery"
	defaultRetryTime            = 3
	defaultRetryWaitTime        = 500 * time.Millisecond
)

var ErrAppIDEmpty = ecode.RequestAppIdEmpty
var ErrSchemeType = ecode.RequestSchemeEmpty
var ErrHttpRespCode = ecode.RequestError

var defaultRetryCondition = func(r *resty.Response, err error) bool {
	return err != nil || !r.IsSuccess()
}

type ClientConfig struct {
	xSecretid  string
	xSignature string
	appid      string
	dial       xtime.Duration
	timeout    xtime.Duration
	keepAlive  xtime.Duration
	scheme     Scheme // http or https or discovery
	host       string
	retryConf  RetryConfig
}

type RetryConfig struct {
	retryTime     int
	retryWaitTime time.Duration
	condition     []resty.RetryConditionFunc
}

func (s *ClientConfig) Apply(params ...ClientParam) *ClientConfig {
	for _, param := range params {
		param(s)
	}
	return s
}

type defaultHttpRespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl"`
	Data    interface{} `json:"data"`
}

func (s *defaultHttpRespData) IsSuccess() bool {
	return s.Code == 200 || s.Code == 0
}

func WithRequestNecessaryParam(appid, host, xSecretid, xSignature string, scheme Scheme) ClientParam {
	return func(cfg *ClientConfig) {
		cfg.appid = appid
		cfg.host = host
		cfg.xSecretid = xSecretid
		cfg.xSignature = xSignature
		cfg.scheme = scheme
	}
}

func WithRequestTimeoutParam(timeout xtime.Duration) ClientParam {
	return func(cfg *ClientConfig) {
		cfg.timeout = timeout
	}
}

func WithRequestRetryParam(retryTime int, retryWaitTime time.Duration) ClientParam {
	return func(cfg *ClientConfig) {
		cfg.retryConf.retryTime = retryTime
		cfg.retryConf.retryWaitTime = retryWaitTime
	}
}

func WithRequestRetryConditionParam(conditions ...resty.RetryConditionFunc) ClientParam {
	return func(cfg *ClientConfig) {
		cfg.retryConf.condition = conditions
	}
}

func fixConfig(cfg *ClientConfig) error {
	if cfg.timeout <= 0 {
		cfg.timeout = xtime.Duration(5 * time.Second)
	}

	if cfg.scheme == "" {
		cfg.scheme = HTTP
	} else if cfg.scheme == Discovery {
		if cfg.appid == "" {
			return ErrAppIDEmpty
		}
	} else if cfg.scheme != HTTP && cfg.scheme != Discovery && cfg.scheme != HTTPs {
		return ErrSchemeType
	}

	if cfg.scheme == Discovery {
		cfg.host = cfg.appid
	}

	fixRetryConfig(&cfg.retryConf)
	return nil
}

func fixRetryConfig(cfg *RetryConfig) {
	if cfg.retryTime <= 0 {
		cfg.retryTime = defaultRetryTime
	}
	if cfg.retryWaitTime <= 0 {
		cfg.retryWaitTime = defaultRetryWaitTime
	}
	if len(cfg.condition) == 0 {
		cfg.condition = append(cfg.condition, defaultRetryCondition)
	}
}
