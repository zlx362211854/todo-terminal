package handler

import (
	"fmt"
	"os"
	"terminal/service/ctx"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func statusTransformer(val interface{}) string {
	str := fmt.Sprintf("%v", val)
	if val == "DONE" {
		return color.GreenString(str)
	}
	return color.YellowString(str)
}
func ShowTodoTable(ctx *ctx.Ctx) {
	data := ctx.Data.GetData()

	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:        "Status",
			Transformer: statusTransformer,
		},
	})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Description", "Status"})

	var rows []table.Row
	for index, row := range data {
		r := table.Row{
			index, row.Title, row.Description, row.Status,
		}
		rows = append(rows, r)
	}
	t.AppendRows(rows)
	t.AppendSeparator()
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}
