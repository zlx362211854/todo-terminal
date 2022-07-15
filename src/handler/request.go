package handler

import (
	"encoding/json"
	"os"
	"terminal/request"
	"terminal/service/ctx"

	"github.com/fatih/color"
)

func HandleList(ctx *ctx.Ctx) {
	color.Green("Query...")
	token, _ := ctx.Auth.GetAuthToken()
	resp := request.GetList(ctx, token)
	if resp.Response.Code == 200 {
		s, err := json.Marshal(resp.Result)
		if err != nil {
			println(err)
		} else {
			ctx.Data.SetData(string(s))
			ShowTodoTable(ctx)
		}
	}
	if resp.Response.Code == 100 {
		println(resp.Response.Message)
	}
}

func HandleCreate(ctx *ctx.Ctx, title, description, status string) bool {
	color.Green("Create...")
	token, _ := ctx.Auth.GetAuthToken()
	resp := request.Create(ctx, token, title, description, status)
	if resp.Response.Code == 200 {
		return true
	}
	if resp.Response.Code == 100 {
		color.Yellow(resp.Response.Message)
		os.Exit(3)
	}
	return false
}

func HandleDelete(ctx *ctx.Ctx, index int) {
	data := ctx.Data.GetData()
	l := len(data)
	if index >= l || index < 0 {
		panic("index does not exist!")
	}
	current := data[index]
	print(current.Id)
	token, _ := ctx.Auth.GetAuthToken()
	request.DeleteItem(ctx, token, current.Id)
}

func HandleModify(ctx *ctx.Ctx, index int, title, description, status string) bool {
	data := ctx.Data.GetData()
	l := len(data)
	if index >= l || index < 0 {
		panic("index does not exist!")
	}
	resp := request.Update(ctx, data[index].Id, title, description, status)
	if resp.Response.Code == 200 {
		return true
	}
	if resp.Response.Code == 100 {
		color.Yellow(resp.Response.Message)
		os.Exit(3)
	}
	return false

}

func HandleDone(ctx *ctx.Ctx, index int) bool {
	data := ctx.Data.GetData()
	l := len(data)
	if index >= l || index < 0 {
		panic("index does not exist!")
	}
	current := data[index]
	id := current.Id
	title := current.Title
	description := current.Description
	resp := request.Update(ctx, id, title, description, "DONE")
	if resp.Response.Code == 200 {
		return true
	}
	if resp.Response.Code == 100 {
		color.Yellow(resp.Response.Message)
		os.Exit(3)
	}
	return false
}
