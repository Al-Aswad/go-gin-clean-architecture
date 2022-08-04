package services

type AuthSevice interface {
	VerifyCredential(email string, password string) interface{}
}
