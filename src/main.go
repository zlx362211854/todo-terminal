package main

import (
	"fmt"
	"os"
	"terminal/handler"
	"terminal/service/auth"
	"terminal/service/ctx"
	"terminal/service/data"
	"terminal/util"
)

func main() {
	var args util.StrSlice = os.Args[1:]
	i := args.FindIndex("-t")
	fmt.Println(i)
	// args =["create", "-t", "title", "-d", "deded", "-s", "TODO"]
	auth := auth.NewAuth()
	data := data.NewTodoData()
	ctx := &ctx.Ctx{
		Auth: auth,
		Data: data,
	}

	handler.HandleCommand(ctx, args)

	// if argLen > 0 {
	// 	if args[0] == "login" {
	// 		handler.HandleLogin(ctx)
	// 	}
	// 	if args[0] == "ls" {
	// 		handler.HandleList(ctx)
	// 	}
	// } else {
	// 	selectedCommandType := prompt.StartPrompt(token)
	// 	if selectedCommandType == "Login" {
	// 		handler.HandleLogin(ctx)
	// 		handler.HandleList(ctx)
	// 	}
	// 	if selectedCommandType == "ShowTODOS" {
	// 		token, _ := auth.GetAuthToken()
	// 		if token == "" {
	// 			color.Red("please login first")
	// 			return
	// 		}
	// 		handler.HandleList(ctx)
	// 	}
	// }
}
