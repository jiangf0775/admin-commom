package web

const OK = 0
const FAIL = -1

const (
	TOKEN_PAYLOAD        string = "payload"
	TOKEN_PAYLOAD_SPLIT         = "|"
	TOKEN_PAYLOAD_SPLIT2        = ","
)

const (
	USER_NAME uint8 = iota
	USER_ID
	USER_ROLE_IDS
)
