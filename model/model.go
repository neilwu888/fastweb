package model

type TsPassword struct {
	Id       int64
	Name     string
	Password string
	Details  string
}

type TsPasswordList struct {
	Name  string
	Index int64
	List  []TsPassword
}
