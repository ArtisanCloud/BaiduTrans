package general

import (
	"github.com/ArtisanCloud/BaiduTrans/config"
	"github.com/ArtisanCloud/BaiduTrans/kernel"
)

type GeneralTrans struct {
	*Client
}

func NewGeneralTrans(config *config.BaiduTranslateConfig) (*GeneralTrans, error) {
	var err error
	baseClient, err := kernel.NewBaseClient(config)
	if err != nil {
		return nil, err
	}
	// init app
	app := &GeneralTrans{
		Client: &Client{
			baseClient,
		},
	}

	return app, err
}
