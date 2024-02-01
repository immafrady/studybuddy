package promptuihelper

import (
	"github.com/manifoldco/promptui"
)

type MultipleOption[T any] struct {
	Label   string
	Value   T
	Detail  string
	Checked bool
}

func (m *MultipleOption[T]) onCheck() {
	m.Checked = !m.Checked
}

func MultipleChoiceSelect[T any](options []MultipleOption[T], config SelectConfig) (ret []T) {
	options = append([]MultipleOption[T]{{Label: "Confirm"}}, options...)
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
