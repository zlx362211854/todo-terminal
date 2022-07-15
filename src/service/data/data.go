package data

import (
	"encoding/json"
	"log"
	"os"
	"terminal/types"
)

type TODOData struct {
}

func NewTodoData() *TODOData {
	return &TODOData{}
}

func (todo *TODOData) SetData(data string) {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	if err != nil {
		println(err)
	}
	_, err2 := f.WriteString(data)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func (todo *TODOData) GetData() []types.TaskDto {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		println(err)
	}
	var jsonMap []types.TaskDto
	json.Unmarshal(data, &jsonMap)
	return jsonMap
}
