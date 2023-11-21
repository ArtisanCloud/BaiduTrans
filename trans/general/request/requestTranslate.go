package request

type RequestGeneralTrans struct {
	Query  string `form:"q" json:"q"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	AppId  string `form:"appid" json:"appid"`
	Action string `form:"action,omitempty" json:"action,omitempty"`
}
