package promptuihelper

import (
	"github.com/manifoldco/promptui"
)

type MultipleOption[T any] struct {
	Label   string
	Value   T
	Checked bool
}

func (m *MultipleOption[T]) onCheck() {
	m.Checked = !m.Checked
}

func MultipleChoiceSelect[T any](options []MultipleOption[T], label string) (ret []T) {
	options = append([]MultipleOption[T]{{Label: "Confirm"}}, options...)
	l := len(options)
	index := -1
	for index != 0 {
		prompt := promptui.Select{
			Label: label,
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}：",
				Active:   "● {{ if .Value }}[{{ if .Checked }}x{{ else }} {{ end}}] {{ .Label }}{{ else }}{{ .Label }}{{ end }}",
				Inactive: "  {{ if .Value }}[{{ if .Checked }}x{{ else }} {{ end}}] {{ .Label }}{{ else }}{{ .Label }}{{ end }}",
			},
			Size:         l,
			HideSelected: true,
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
