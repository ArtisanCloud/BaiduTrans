package response

type ResponseBase struct {
	ErrCode string `json:"error_code,omitempty"`
	ErrMsg  string `json:"error_msg,omitempty"`
}
