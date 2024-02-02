package promptuihelper

import (
	"github.com/manifoldco/promptui"
)

func MultipleChoiceSelect[T any](options []Option[T], config SelectConfig) (ret []T) {
	var confirmText string
	if config.ConfirmText == "" {
		confirmText = "提交"
	} else {
		confirmText = config.ConfirmText
	}
	options = append([]Option[T]{{Label: confirmText}}, options...)
	l := len(options)
	index := -1
	for index != 0 {
		prompt := promptui.Select{
			Label: config.Label,
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label:    config.LabelTemplate,
				Active:   "● {{ if .Value }}[{{ if .Checked }}x{{ else }} {{ end}}] {{ .Label }}{{ else }}{{ .Label }}{{ end }}",
				Inactive: "  {{ if .Value }}[{{ if .Checked }}x{{ else }} {{ end}}] {{ .Label }}{{ else }}{{ .Label }}{{ end }}",
				Selected: config.Selected,
				Details:  config.Details,
			},
			Size:         l,
			HideSelected: config.Selected == "",
		}
		index, _, _ = prompt.RunCursorAt(index, 0)
		if index != 0 && index < l {
			options[index].onCheck()
		}
	}

	for i := 1; i < l; i++ {
		if options[i].Checked {
			ret = append(ret, options[i].Value)
		}
	}
	return
}
