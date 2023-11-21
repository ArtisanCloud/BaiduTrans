package response

import "github.com/ArtisanCloud/BaiduTrans/kernel/response"

type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type ResponseGeneralTrans struct {
	response.ResponseBase
	From         string         `json:"from"`
	To           string         `json:"to"`
	TransResults []*TransResult `json:"trans_result"`
}
