package response

type Response struct {
	Code int
	Msg  string
	Data interface{}
}
type response struct {
	Code int
	Msg  string
}

func Responser(in Response) interface{} {
	if in.Data == nil {
		return response{
			Code: in.Code,
			Msg:  in.Msg,
		}
	}

	return in
}
