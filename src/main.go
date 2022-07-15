package main

import (
	"terminal/handler"
	"terminal/service/auth"
	"terminal/service/ctx"
	"terminal/service/data"
)

func main() {

	auth := auth.NewAuth()
	data := data.NewTodoData()
	ctx := &ctx.Ctx{
		Auth: auth,
		Data: data,
	}
	handler.HandleCommand(ctx)
}
