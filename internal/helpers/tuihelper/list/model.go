package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list       list.Model
	help       help.Model
	Options    Options
	isMultiple bool
	liked      bool
	likeFn     func()
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyLike):
			if m.likeFn != nil {
				m.liked = !m.liked
				m.likeFn()
			}
		case key.Matches(msg, keyEnter):
			if m.isMultiple { //
				// 多选下仅提交
			} else {
				// 单选下是选中并提交
				m.toggleCheck()
				return m, tea.Quit
			}
		case key.Matches(msg, keySpace):
			if m.isMultiple {
				// 多选下是选中
				m.toggleCheck()
			}
		case key.Matches(msg, keyHelp):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, nil
}

func (m Model) View() string {
	//TODO implement me
	panic("implement me")
}

func (m Model) toggleCheck() {
	idx := m.list.Index()
	m.Options[idx].ToggleCheck()
}
