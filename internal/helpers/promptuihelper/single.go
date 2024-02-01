package promptuihelper

import (
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/manifoldco/promptui"
)

type SingleOption[T any] struct {
	Label string
	Value T
}

func SingleChoiceSelect[T any](options []SingleOption[T], label string) (ret T) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}：",
			Active:   "● {{ .Label }}",
			Inactive: "  {{ .Label }}",
		},
		HideSelected: true,
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
