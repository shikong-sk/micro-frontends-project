package Error

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	DatabaseConnectError = Error{
		Code: 500,
		Msg:  "数据库连接失败",
	}
)
