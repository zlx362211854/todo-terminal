package prompt

import (
	"terminal/types"
	"terminal/validate"

	"github.com/manifoldco/promptui"
)

type todoType struct {
	Name string
}

func ModifyPrompt(item types.TaskDto) (title, description, status string) {
	titlePrompt := promptui.Prompt{
		Label:       "Title",
		Default:     item.Title,
		HideEntered: true,
		Validate: func(input string) error {
			_, err := validate.ValidateTitle(input)
			if err != nil {
				return err
			}
			return nil
		},
	}
	descPrompt := promptui.Prompt{
		Label:       "Description",
		Default:     item.Description,
		HideEntered: true,
		Validate: func(input string) error {
			_, err := validate.ValidateTitle(input)
			if err != nil {
				return err
			}
			return nil
		},
	}

	statusList := []todoType{
		{Name: "TODO"},
		{Name: "DONE"},
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Name | green }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "> {{ .Name | cyan }}",
	}
	statusPrompt := promptui.Select{
		Label:     "Status",
		Items:     statusList,
		Templates: templates,
		Size:      4,
	}

	title, _ = titlePrompt.Run()
	description, _ = descPrompt.Run()

	i, _, _ := statusPrompt.Run()
	status = statusList[i].Name
	return title, description, status
}
