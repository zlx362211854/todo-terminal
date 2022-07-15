package handler

import (
	"strconv"
	"terminal/prompt"
	"terminal/service/ctx"
	"terminal/util"

	"github.com/fatih/color"
)

func HandleCommand(ctx *ctx.Ctx, args util.StrSlice) {
	var token, _ = ctx.Auth.GetAuthToken()
	var argLen = len(args)
	var command = args[0]

	if argLen > 0 {
		if command == "login" {
			HandleLogin(ctx)
		} else if command == "ls" {
			HandleList(ctx)
		} else {
			handleCRUD(ctx, args, command)
		}
	} else {
		selectedCommandType := prompt.StartPrompt(token)
		if selectedCommandType == "Login" {
			HandleLogin(ctx)
			HandleList(ctx)
		}
		if selectedCommandType == "ShowTODOS" {
			token, _ := ctx.Auth.GetAuthToken()
			if token == "" {
				color.Red("please login first")
				return
			}
			HandleList(ctx)
		}
	}
}

func handleCRUD(ctx *ctx.Ctx, args util.StrSlice, command string) {
	var title string = util.GetValueFromArgs(args, "-t")
	var description string = util.GetValueFromArgs(args, "-d")
	var status string = util.GetValueFromArgs(args, "-s")

	if command == "create" {
		createSuccessfully := HandleCreate(ctx, title, description, status)
		if createSuccessfully {
			HandleList(ctx)
		}
	}

	if command == "delete" {
		index, _ := strconv.Atoi(args[1])
		HandleDelete(ctx, index)
		HandleList(ctx)
	}

	if command == "done" {
		index, _ := strconv.Atoi(args[1])
		updateSuccessfully := HandleDone(ctx, index)
		if updateSuccessfully {
			HandleList(ctx)
		}
	}
	if command == "modify" {
		index, _ := strconv.Atoi(args[1])
		var current = ctx.Data.GetData()[index]
		title, description, status := prompt.ModifyPrompt(current)
		modifySuccessfully := HandleModify(ctx, index, title, description, status)
		if modifySuccessfully {
			HandleList(ctx)
		}
	}
}
