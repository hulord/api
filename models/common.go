package models
type Result struct {
    Code    int
    Msg 	string
    Data    interface{}
}
const (
    SUCCESS = iota
    FAIL
)

func SendResponse(code int, Meg string, Data interface{}) (result Result) {
    result.Code = code
    result.Msg = Meg
	result.Data = Data
    return 
}