package e

import "fmt"

var msgFlags = map[int]string{
	Success:                 "ok",
	Error:                   "fail",
	ErrorDatabaseConnection: "failed to connect a database",
}

func getMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[Error]
}

// PrintError dumps error message
func PrintError(code int) {
	fmt.Printf("[Error] code=%d, msg=%s\n", code, getMsg(code))
}
