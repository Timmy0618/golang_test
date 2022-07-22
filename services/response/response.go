package response

type response struct {
	Code int
	Msg  string
}
type response1 struct {
	Code int
	Msg  string
	Data interface{}
}

func Responser(code int, msg string, data ...interface{}) interface{} {
	var res interface{}
	if data == nil {
		res = response{Code: code, Msg: msg}
	} else {
		res = response1{Code: code, Msg: msg, Data: data}
	}

	return res
}
