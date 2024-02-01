package promptuihelper

import (
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/manifoldco/promptui"
)

type SingleOption[T any] struct {
	Label  string
	Value  T
	Detail string
}

func SingleChoiceSelect[T any](options []SingleOption[T], config SelectConfig) (ret T) {
	prompt := promptui.Select{
		Label: config.Label,
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    config.LabelTemplate,
			Active:   "● {{ .Label }}",
			Inactive: "  {{ .Label }}",
			Selected: config.Selected,
			Details:  config.Details,
		},
		HideSelected: config.Selected == "",
		Size:         5,
	}
	i, _, err := prompt.Run()
	errorhelper.LogError(err)
	if i < len(options) {
		option := options[i]
		ret = option.Value
	}
	return
}
