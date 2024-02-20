/**
 * @Author haochen
 * @Email: lizekun01@bilibili.com
 * @Date: 2024/2/20 20:59
 * @Description: 描述
**/

package requests

import (
	"context"

	"go-common/library/log"
)

/*
WaveClient 使用方法
client, err := requests.NewDefaultClient(

	requests.WithRequestNecessaryParam(conf.Appid, conf.Url, conf.XSecretId, conf.XSignature, requests.Scheme(conf.Scheme)),

)
可选择
WithRequestNecessaryParam
WithRequestTimeoutParam
WithRequestRetryParam
WithRequestRetryConditionParam
*/
type WaveClient struct {
	client      *restyClient
	retryClient *restyClient
}

func NewDefaultClient(params ...ClientParam) (*WaveClient, error) {
	cfg := new(ClientConfig).Apply(params...)

	if err := fixConfig(cfg); err != nil {
		return nil, err
	}

	client := newRestyClient(cfg)

	retryClient := newRestyClientWithRetry(cfg)

	return &WaveClient{
		client:      client,
		retryClient: retryClient,
	}, nil
}

func (s *WaveClient) SentPostUrl(ctx context.Context, path string, body interface{}, result interface{}) error {

	resp := &defaultHttpRespData{
		Data: result,
	}

	if err := s.getClient(false).SentPostUrl(ctx, path, body, resp); err != nil {
		log.Error("WaveClient SentPostUrl err:%s, resp:%+v", err.Error(), resp)
		return err
	}

	if !resp.IsSuccess() {
		log.Error("WaveClient SentPostUrl err url:%s,body:%+v,resp: %+v", path, body, resp)
		return ErrHttpRespCode
	}

	return nil
}

func (s *WaveClient) SentPostUrlWithRetry(ctx context.Context, path string, body interface{}, result interface{}) error {

	resp := &defaultHttpRespData{
		Data: result,
	}

	if err := s.getClient(true).SentPostUrl(ctx, path, body, resp); err != nil {
		log.Error("WaveClient SentPostUrlWithRetry err:%s, resp:%+v", err.Error(), resp)
		return err
	}

	if !resp.IsSuccess() {
		log.Error("WaveClient SentPostUrlWithRetry err url:%s,body:%+v,resp: %+v", path, body, resp)
		return ErrHttpRespCode
	}

	return nil
}

/*
SentGetUrl queryData 参考格式

	type Test struct {
		AppId string `url:"app_id"`
	}
*/
func (s *WaveClient) SentGetUrl(ctx context.Context, path string, queryData interface{}, result interface{}) error {

	resp := &defaultHttpRespData{
		Data: result,
	}

	if err := s.getClient(false).SentGetUrl(ctx, path, queryData, resp); err != nil {
		log.Error("WaveClient SentGetUrl err:%s, resp:%+v", err.Error(), resp)
		return err
	}

	if !resp.IsSuccess() {
		log.Error("WaveClient SentGetUrl data err url:%s,queryData:%+v,resp: %+v", path, queryData, resp)
		return ErrHttpRespCode
	}

	return nil
}

func (s *WaveClient) SentGetUrlWithRetry(ctx context.Context, path string, queryData interface{}, result interface{}) error {

	resp := &defaultHttpRespData{
		Data: result,
	}

	if err := s.getClient(true).SentGetUrl(ctx, path, queryData, resp); err != nil {
		log.Error("WaveClient SentGetUrlWithRetry err:%s, resp:%+v", err.Error(), resp)
		return err
	}

	if !resp.IsSuccess() {
		log.Error("WaveClient SentGetUrlWithRetry err url:%s,queryData:%+v,resp: %+v", path, queryData, resp)
		return ErrHttpRespCode
	}

	return nil
}

func (s *WaveClient) SentDeleteUrl(ctx context.Context, path string, body interface{}, result interface{}) error {

	resp := &defaultHttpRespData{
		Data: result,
	}

	if err := s.getClient(false).SentDeleteUrl(ctx, path, body, resp); err != nil {
		log.Error("WaveClient SentDeleteUrl err:%s, resp:%+v", err.Error(), resp)
		return err
	}

	if !resp.IsSuccess() {
		log.Error("WaveClient SentDeleteUrl err url:%s,body:%+v,resp: %+v", path, body, resp)
		return ErrHttpRespCode
	}

	return nil
}

func (s *WaveClient) SentDeleteUrlWithRetry(ctx context.Context, path string, body interface{}, result interface{}) error {

	resp := &defaultHttpRespData{
		Data: result,
	}

	if err := s.getClient(true).SentDeleteUrl(ctx, path, body, resp); err != nil {
		log.Error("WaveClient SentDeleteUrlWithRetry err:%s, resp:%+v", err.Error(), resp)
		return err
	}

	if !resp.IsSuccess() {
		log.Error("WaveClient SentDeleteUrlWithRetry err url:%s,body:%+v,resp: %+v", path, body, resp)
		return ErrHttpRespCode
	}

	return nil
}

func (s *WaveClient) getClient(needRetry bool) *restyClient {
	if needRetry {
		return s.retryClient
	}
	return s.client
}
