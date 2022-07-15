package auth

import (
	"log"
	"os"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}
func (a *Auth) SetAuthToken(token string) error {
	f, err := os.Create("auth.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	if err != nil {
		println(err)
	}
	_, err2 := f.WriteString(token)

	if err2 != nil {
		log.Fatal(err2)
	}
	return err2
}

func (a *Auth) GetAuthToken() (authToken string, err error) {
	dat, err := os.ReadFile("auth.txt")
	if err != nil {
		println(err)
	}
	token := string(dat)
	return token, err
}
