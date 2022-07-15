package request

import (
	"encoding/json"
	"io"
	"net/http"
	"terminal/service/ctx"
	"terminal/types"

	"github.com/mozillazg/request"
)

func Login(username string, password string) types.LoginResponse {
	client := new(http.Client)
	req := request.NewRequest(client)
	req.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	req.Json = map[string]string{
		"username": username,
		"password": password,
	}
	resp, err := req.Post(LoginUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	jsonString := string(b)
	var jsonMap types.LoginResponse
	json.Unmarshal([]byte(jsonString), &jsonMap)
	return jsonMap
}
func Create(ctx *ctx.Ctx, token string, title, description, status string) types.CreateResponse {
	client := new(http.Client)
	req := request.NewRequest(client)
	req.Json = map[string]string{
		"title":       title,
		"description": description,
		"status":      status,
	}
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}
	resp, err := req.Post(TaskUrl)
	Authentication(ctx, resp)
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var jsonMap types.CreateResponse
	json.Unmarshal(b, &jsonMap)
	return jsonMap

}

func GetList(ctx *ctx.Ctx, token string) types.ListResponse {
	client := new(http.Client)
	req := request.NewRequest(client)
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}
	resp, err := req.Get(TaskUrl)
	Authentication(ctx, resp)
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var jsonMap types.ListResponse
	json.Unmarshal(b, &jsonMap)
	return jsonMap
}

func DeleteItem(ctx *ctx.Ctx, token string, id string) {
	client := new(http.Client)
	req := request.NewRequest(client)
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}
	resp, err := req.Delete(TaskUrl + "/" + id)
	Authentication(ctx, resp)
	io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
}

func Update(ctx *ctx.Ctx, id, title, description, status string) types.UpdateResponse {
	token, _ := ctx.Auth.GetAuthToken()
	client := new(http.Client)
	req := request.NewRequest(client)
	req.Json = map[string]string{
		"title":       title,
		"description": description,
		"status":      status,
	}
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}
	resp, err := req.Put(TaskUrl + "/" + id)
	Authentication(ctx, resp)
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var jsonMap types.UpdateResponse
	json.Unmarshal(b, &jsonMap)
	return jsonMap
}
