package general

import (
	"context"
	"github.com/ArtisanCloud/BaiduTrans/kernel"
	"github.com/ArtisanCloud/BaiduTrans/trans/general/request"
	"github.com/ArtisanCloud/BaiduTrans/trans/general/response"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// https://api.fanyi.baidu.com/doc/21
func (comp *Client) Translate(ctx context.Context, param *request.RequestGeneralTrans) (*response.ResponseGeneralTrans, error) {

	result := &response.ResponseGeneralTrans{}

	if param.AppId == "" {
		param.AppId = comp.BaseClient.Config.AppID
	}

	data, _ := object.StructToHashMap(param)
	salt := kernel.GenerateSalt()
	(*data)["sign"] = kernel.GenerateSignature(param.AppId, param.Query, salt, comp.BaseClient.Config.AppSecret)
	(*data)["salt"] = salt
	_, err := comp.BaseClient.HttpPost(ctx, "api/trans/vip/translate", data, nil, result)

	return result, err
}
