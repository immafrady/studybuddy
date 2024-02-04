package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
)

type Config struct {
	Options  []Option
	Title    string
	Liked    bool
	LikeFn   func()
	Multiple bool
}

func Run(config Config) Model {
	var items = make([]list.Item, len(config.Options))
	for i, o := range config.Options {
		items[i] = o
	}

	m := Model{
		list: list.New(items, NewOptionDelegate(config.Options), 0, 0),
		help: help.New(),
		keyMap: keyMap{
			showLike: config.LikeFn != nil,
			multiple: config.Multiple,
		},
		options:    config.Options,
		isMultiple: config.Multiple,
	}
	m.list.Title = config.Title
	m.list.Styles.Title = lipgloss.NewStyle()
	//m.list.SetShowStatusBar(false)
	m.list.SetShowHelp(false)
	model, err := tea.NewProgram(m, tea.WithAltScreen()).Run()
	if err != nil {
		errorhelper.LogError(err)
	}
	m, _ = model.(Model)
	return m
}
