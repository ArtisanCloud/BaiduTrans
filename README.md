# BaiduTrans



官方文档：  
https://api.fanyi.baidu.com/  
https://api.fanyi.baidu.com/doc/21


通用翻译API HTTPS 地址：  
https://fanyi-api.baidu.com/api/trans/vip/translate


```shell
go get -u github.com/ArtisanCloud/BaiduTrans
```

```go

package main

import (
	"github.com/ArtisanCloud/BaiduTrans/config"
	"github.com/ArtisanCloud/BaiduTrans/trans/general"
	"github.com/ArtisanCloud/BaiduTrans/trans/general/request"
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
)

func main() {

	conf := config.Get()

	var err error
	gApp, err := general.NewGeneralTrans(conf)
	if err != nil || gApp == nil {
		panic(err)
	}

	params := &request.RequestGeneralTrans{
		Query: "室内，暗淡，高品质，桌椅，大床",
		From:  "zh",
		To:    "en",
	}
	res, err := gApp.Translate(context.Background(), params)
	if err != nil || gApp == nil {
		panic(err)
	}
	if res.ErrCode != "" {
		panic(res)
	}

	fmt.Dump(res)

	return
}
```

## 更多产品：
PowerWechat: https://github.com/ArtisanCloud/PowerWeChat  
PowerX: https://github.com/ArtisanCloud/PowerX

