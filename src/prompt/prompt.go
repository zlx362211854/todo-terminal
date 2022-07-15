package prompt

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type pepper struct {
	Name string
	Desc string
}

func StartPrompt(token string) string {
	peppers := []pepper{
		{Name: "ShowTODOS", Desc: "Show todo list"},
	}
	if token == "" {
		peppers = append(peppers, pepper{
			Name: "Login",
			Desc: "Login todo app",
		})
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Name | green }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "> {{ .Name | cyan }}",
		Details: `
--------- Description ----------
{{ .Desc }}`,
	}
	searcher := func(input string, index int) bool {
		pepper := peppers[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Switch:",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	return peppers[i].Name
}
