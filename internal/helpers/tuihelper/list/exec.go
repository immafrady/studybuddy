package list

import "github.com/charmbracelet/bubbles/list"

type Config struct {
	Options []list.Item
	Title   string
	Liked   bool
}

func Run(config Config) {
	for _, o := range config.Options {
		if _, ok := o.(Option); ok {

		}
	}
	m := Model{list: list.New(config.Options, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = config.Title

	// todo 待删除
	list.New([]list.Item{
		Option{
			Item:    nil,
			Label:   "",
			Value:   "",
			Desc:    "",
			Checked: false,
		},
	}, list.NewDefaultDelegate(), 0, 0)
}
