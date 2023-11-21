package main

import (
	"context"
	"github.com/ArtisanCloud/BaiduTrans/config"
	"github.com/ArtisanCloud/BaiduTrans/trans/general"
	"github.com/ArtisanCloud/BaiduTrans/trans/general/request"
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
