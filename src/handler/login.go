package handler

import (
	"os"
	"terminal/request"
	"terminal/service/ctx"
	"terminal/validate"

	"github.com/Luzifer/go-openssl"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var secretKey string = "394812730425442A472D2F423F452848"

func HandleLogin(ctx *ctx.Ctx) {
	color.Green("Please input your username")
	usernamePrompt := promptui.Prompt{
		Label:       "Username",
		HideEntered: true,
		Validate: func(input string) error {
			_, err := validate.ValidateName(input)
			if err != nil {
				return err
			}
			return nil
		},
	}
	usernameResult, _ := usernamePrompt.Run()

	color.Green("Please input your password")
	passwordPrompt := promptui.Prompt{
		Label:       "Password",
		HideEntered: true,
		Mask:        '*',
		Validate: func(input string) error {
			_, err := validate.ValidateName(input)
			if err != nil {
				return err
			}
			return nil
		},
	}
	passwordResult, _ := passwordPrompt.Run()
	o := openssl.New()
	encodePassword, _ := o.EncryptBytes(secretKey, []byte(passwordResult))
	resp := request.Login(usernameResult, string(encodePassword))

	if resp.Response.Code == 200 {
		color.Green("Login successful!")
		ctx.Auth.SetAuthToken(resp.Result.AccessToken)
	}
	if resp.Response.Code == 100 {
		color.Red(resp.Response.Message)
		os.Exit(3)
	}
}
