/**
 * @Author haochen
 * @Email: lizekun01@bilibili.com
 * @Date: 2024/2/20 21:01
 * @Description: 描述
**/

package ecode

import (
	"net/http"

	"go-common/library/ecode"
)

func init() {
	ecode.Register(map[int]map[string]string{
		InvalidArgument.Code():      {"default": "参数不合法"},
		Unknown.Code():              {"default": "未知错误"},
		NotFound.Code():             {"default": "目标不存在"},
		PermissionDenied.Code():     {"default": "权限不足"},
		ThirdPartFailed.Code():      {"default": "三方依赖错误"},
		BlockError.Code():           {"default": "审批中无法修改"},
		JsonMarshalFailed.Code():    {"default": "json marshal failed"},
		SQLErr.Code():               {"default": "DAO错误"},
		SchemaValidateFailed.Code(): {"default": "参数schema校验未通过"},
		SchemaEmpty.Code():          {"default": "参数schema为空"},
		HasHistoryRecord.Code():     {"default": "存在历史发布单未完成"},
		RequestAppIdEmpty.Code():    {"default": "配置文件appid为空"},
		RequestSchemeEmpty.Code():   {"default": "配置文件Scheme为空"},
		RequestError.Code():         {"default": "第三方请求失败"},
	})
}

var (
	InvalidArgument      = ecode.New(-http.StatusBadRequest)
	SchemaValidateFailed = ecode.New(-4001003)
	SchemaEmpty          = ecode.New(-4001005)
	Unknown              = ecode.New(-http.StatusInternalServerError)
	ThirdPartFailed      = ecode.New(-5001001)
	JsonMarshalFailed    = ecode.New(-5001002)
	SQLErr               = ecode.New(-11000001000000)
	NotFound             = ecode.New(-http.StatusNotFound)
	PermissionDenied     = ecode.New(-http.StatusForbidden)
	BlockError           = ecode.New(-http.StatusLocked)
	HasHistoryRecord     = ecode.New(-11000004000002)

	RequestAppIdEmpty  = ecode.New(-10000001000001)
	RequestSchemeEmpty = ecode.New(-10000001000002)
	RequestError       = ecode.New(-10000001000003)
)
