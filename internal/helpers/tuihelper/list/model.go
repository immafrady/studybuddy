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
	options    []Option
	keyMap     keyMap
	isMultiple bool
	liked      bool
	likeFn     func(liked bool)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyUp):
			m.list.CursorUp()
		case key.Matches(msg, keyDown):
			m.list.CursorDown()
		case key.Matches(msg, keyLeft):
			m.list.PrevPage()
		case key.Matches(msg, keyRight):
			m.list.NextPage()
		case key.Matches(msg, keyLike):
			if m.likeFn != nil {
				m.liked = !m.liked
				m.likeFn(m.liked)
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
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetHeight(msg.Height - h)
		m.list.SetWidth(msg.Width - v)
	}
	return m, nil
}

func (m Model) View() string {
	helpView := m.help.View(m.keyMap)
	return docStyle.Render(m.list.View(), helpView)
}

func (m Model) toggleCheck() {
	index := m.list.Index()
	m.options[index].ToggleCheck()
}

func (m Model) GetSelectedOptions() []Option {
	var newOs []Option
	for _, o := range m.options {
		if o.Checked {
			newOs = append(newOs, o)
		}
	}
	return newOs
}
