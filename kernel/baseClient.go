package kernel

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/BaiduTrans/config"
	request2 "github.com/ArtisanCloud/BaiduTrans/kernel/request"
	response2 "github.com/ArtisanCloud/BaiduTrans/kernel/response"
	"github.com/ArtisanCloud/PowerLibs/v3/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	logger2 "github.com/ArtisanCloud/PowerLibs/v3/logger"
	contract2 "github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"net/http"
	"net/url"
	"strings"
)

type BaseClient struct {
	HttpHelper *helper.RequestHelper
	Config     *config.BaiduTranslateConfig
	BaseUri    string

	GetMiddlewareOfLog func(logger contract2.LoggerInterface) contract.RequestMiddleware
}

func NewBaseClient(config *config.BaiduTranslateConfig) (*BaseClient, error) {
	baseUri := "https://fanyi-api.baidu.com"
	if config.BaseUri != "" {
		baseUri = config.BaseUri
	}
	h, err := helper.NewRequestHelper(&helper.Config{
		BaseUrl: baseUri,
	})
	if err != nil {
		return nil, err
	}

	client := &BaseClient{
		HttpHelper: h,
		Config:     config,
		BaseUri:    baseUri,
	}

	client.OverrideGetMiddlewares()
	client.RegisterHttpMiddlewares()

	return client, nil
}

func (client *BaseClient) OverrideGetMiddlewares() {
	client.OverrideGetMiddlewareOfLog()
}

func (client *BaseClient) RegisterHttpMiddlewares() {

	// log
	logMiddleware := client.GetMiddlewareOfLog

	logger, err := logger2.NewLogger(nil, &object.HashMap{
		"env":        client.Config.Env,
		"outputPath": client.Config.Log.InfoPath,
		"errorPath":  client.Config.Log.ErrPath,
	})
	if err != nil {
		panic(err)
	}
	client.HttpHelper.WithMiddleware(
		logMiddleware(logger),
		helper.HttpDebugMiddleware(client.Config.HttpDebug),
	)
}

func (client *BaseClient) HttpPost(ctx context.Context, endpoint string, data interface{}, outHeader interface{}, outBody interface{}) (interface{}, error) {

	return client.Request(
		ctx,
		endpoint,
		http.MethodPost,
		&object.HashMap{
			"form_params": data,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) OverrideGetMiddlewareOfLog() {
	client.GetMiddlewareOfLog = func(logger contract2.LoggerInterface) contract.RequestMiddleware {
		return contract.RequestMiddleware(func(handle contract.RequestHandle) contract.RequestHandle {
			return func(request *http.Request) (response *http.Response, err error) {

				// 前置中间件
				request2.LogRequest(logger, request)

				response, err = handle(request)
				if err != nil {
					return response, err
				}

				// 后置中间件
				response2.LogResponse(logger, response)

				return
			}
		})
	}
}
func (client *BaseClient) Request(ctx context.Context, uri string, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (*http.Response, error) {

	// http client request
	df := client.HttpHelper.Df().WithContext(ctx).Uri(uri).Method(method)

	// 检查是否需要有请求参数配置
	if options != nil {
		// set query key values
		if (*options)["query"] != nil {
			queries := (*options)["query"].(*object.StringMap)
			if queries != nil {
				for k, v := range *queries {
					df.Query(k, v)
				}
			}
		}

		// set body json
		if (*options)["form_params"] != nil {
			if formParams, ok := (*options)["form_params"].(*object.HashMap); ok {
				// 准备要传递的参数
				params := url.Values{}

				// 假设 formParams 是 map[string]interface{} 类型
				for key, value := range *formParams {
					params.Add(key, fmt.Sprintf("%v", value))
				}
				body := strings.NewReader(params.Encode())
				df.Body(body).
					Header("Content-Type", "application/x-www-form-urlencoded")
			} else {
				return nil, errors.New("缺少提交数据")
			}
		}
	}

	ctxQuery := ctx.Value("query")
	if ctxQuery != nil {
		queries := ctxQuery.(*object.StringMap)
		if queries != nil {
			for k, v := range *queries {
				df.Query(k, v)
			}
		}
	}

	response, err := df.Request()
	if err != nil {
		return response, err
	}

	// decode response body to outBody
	if outBody != nil {
		err = client.HttpHelper.ParseResponseBodyContent(response, outBody)
		if err != nil {
			return nil, err
		}
	}

	// decode response header to outHeader
	if outHeader != nil {
		strHeader, err := object.JsonEncode(response.Header)
		if err != nil {
			return nil, err
		}
		err = object.JsonDecode([]byte(strHeader), outHeader)
		if err != nil {
			return nil, err
		}
	}

	return response, err

}

func AuthSignRequest() {

}
