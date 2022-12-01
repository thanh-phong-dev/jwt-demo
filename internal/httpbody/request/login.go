package request

type Login struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}
