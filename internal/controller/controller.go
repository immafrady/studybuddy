package controller

import "github.com/immafrady/studybuddy/internal/helpers/promptuihelper"

type entryItem struct {
	Label    string
	Detail   string
	Callback func()
}

func (e entryItem) toSingleOption() promptuihelper.SingleOption[entryItem] {
	return promptuihelper.SingleOption[entryItem]{
		Label:  e.Label,
		Value:  e,
		Detail: e.Detail,
	}
}

func Navigate() {
	entries := []entryItem{
		startEntry,
		examEntry,
		reviewEntry,
		historyEntry,
	}
	var options = make([]promptuihelper.SingleOption[entryItem], len(entries))
	for i, entry := range entries {
		options[i] = entry.toSingleOption()
	}

	item := promptuihelper.SingleChoiceSelect(options, promptuihelper.SelectConfig{
		Label: "请选择模式",
		Details: `
------模式解析-----
{{ .Detail }}
`,
	})

	item.Callback()
}
