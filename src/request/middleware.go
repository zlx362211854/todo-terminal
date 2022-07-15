package request

import (
	"os"
	"terminal/service/ctx"

	"github.com/fatih/color"
	"github.com/mozillazg/request"
)

func Authentication(ctx *ctx.Ctx, resp *request.Response) {
	if resp.StatusCode == 400 {
		color.Yellow("The token has expired. Please login again!")
		ctx.Auth.SetAuthToken("")
		os.Exit(3)
	}
}
