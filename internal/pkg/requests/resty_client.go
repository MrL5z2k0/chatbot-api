/**
 * @Author haochen
 * @Email: lizekun01@bilibili.com
 * @Date: 2024/2/20 21:01
 * @Description: 描述
**/

package requests

import (
	"context"
	"fmt"
	"time"

	"go-common/library/log"
	"go-common/library/net/http/bmc"
	"go-common/library/net/http/bmc/xresty"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type restyClient struct {
	cfg    *ClientConfig
	resCli *resty.Client
}

func newRestyClient(cfg *ClientConfig) *restyClient {

	// 默认校验将非2xx 状态码认为错误的
	resCli := xresty.New(&bmc.Config{
		Timeout: cfg.timeout,
		DialConfig: bmc.DialConfig{
			cfg.dial,
			cfg.keepAlive,
		},
	})

	return &restyClient{
		cfg:    cfg,
		resCli: resCli,
	}
}

func newRestyClientWithRetry(cfg *ClientConfig) *restyClient {

	// 默认校验将非2xx 状态码认为错误的
	resCli := xresty.New(&bmc.Config{
		Timeout: cfg.timeout,
		DialConfig: bmc.DialConfig{
			cfg.dial,
			cfg.keepAlive,
		},
	})

	resCli.SetRetryCount(cfg.retryConf.retryTime).
		SetRetryWaitTime(cfg.retryConf.retryWaitTime).
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			log.Error("请求重试url:%s, resp :%d", client.HostURL, resp.StatusCode())
			return 0, nil
		})
	for i := 0; i < len(cfg.retryConf.condition); i++ {
		resCli.AddRetryCondition(cfg.retryConf.condition[i])
	}

	return &restyClient{
		cfg:    cfg,
		resCli: resCli,
	}
}

/*
不推荐
*/
func (s *restyClient) httpBase() *resty.Request {
	return s.resCli.R().SetHeader("Content-Type", "application/json").
		SetHeader("x-secretid", s.cfg.xSecretid).
		SetHeader("x-signature", s.cfg.xSignature)
}

/*
推荐
*/
func (s *restyClient) httpBaseWithContext(ctx context.Context) *resty.Request {
	return s.resCli.R().
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetHeader("x-secretid", s.cfg.xSecretid).
		SetHeader("x-signature", s.cfg.xSignature)
}

func (s *restyClient) SentPostUrl(ctx context.Context, path string, body interface{}, result interface{}) error {
	url := s.getUrl(path)

	resp, err := s.httpBaseWithContext(ctx).SetBody(body).SetResult(result).Post(url)
	log.Info("SentPostUrl url:%s,body:%+v,resp: %+v", url, body, resp)
	if err != nil {
		log.Error("SentPostUrl url:%s,body:%+v,resp: %+v,err:%s", url, body, resp, err.Error())
		return err
	}

	return nil
}

func (s *restyClient) SentGetUrl(ctx context.Context, path string, queryData interface{}, result interface{}) error {

	queryParams, err := query.Values(queryData)
	if err != nil {
		return err
	}

	url := s.getUrl(path)

	resp, err := s.httpBaseWithContext(ctx).SetQueryParamsFromValues(queryParams).SetResult(result).Get(url)
	log.Info("SentGetUrl url:%s,queryParams:%+v,resp: %+v", url, queryParams, resp)
	if err != nil {
		log.Error("SentGetUrl url:%s,queryParams:%+v,resp: %+v,err:%s", url, queryParams, resp, err.Error())
		return err
	}

	return nil
}

func (s *restyClient) SentDeleteUrl(ctx context.Context, path string, body interface{}, result interface{}) error {
	url := s.getUrl(path)

	resp, err := s.httpBaseWithContext(ctx).SetBody(body).SetResult(result).Delete(url)
	log.Info("SentDeleteUrl url:%s,body:%+v,resp: %+v", url, body, resp)
	if err != nil {
		log.Error("SentDeleteUrl url:%s,body:%+v,resp: %+v,err:%s", url, body, resp, err.Error())
		return err
	}

	return nil
}

func (s *restyClient) getUrl(path string) string {
	return fmt.Sprintf(`%s://%s%s`, s.cfg.scheme, s.cfg.host, path)
}
