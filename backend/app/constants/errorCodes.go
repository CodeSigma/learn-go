package constants

const (
	CD_SUCCESS        = 200
	CD_ERROR          = 500
	CD_INVALID_PARAMS = 400
)

var MsgFlags = map[int]string{
	CD_SUCCESS: "ok",
	CD_ERROR:   "fail",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[CD_ERROR]
}
