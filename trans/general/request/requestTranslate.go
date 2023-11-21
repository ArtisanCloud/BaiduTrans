package request

type RequestGeneralTrans struct {
	Query  string `form:"q"`
	From   string `form:"from"`
	To     string `form:"to"`
	AppId  string `form:"appid"`
	Action string `form:"action,omitempty"`
}
