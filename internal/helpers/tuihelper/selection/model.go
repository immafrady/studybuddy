package selection

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type Model struct {
	tea.Model
	help       help.Model
	itemStyles DefaultItemStyles
	*handler
	options  []*Option
	liked    bool
	likeFn   func(bool)
	multiple bool
	title    string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyUp):
			m.handler.MoveUp()
		case key.Matches(msg, keyDown):
			m.handler.MoveDown()
		case key.Matches(msg, keyLike):
			if m.likeFn != nil {
				m.liked = !m.liked
				m.likeFn(m.liked)
			}
		case key.Matches(msg, keyEnter):
			if m.multiple { //
				// 多选下仅提交
				return m, tea.Quit
			} else {
				// 单选下是选中并提交
				m.toggleCheck()
				return m, tea.Quit
			}
		case key.Matches(msg, keySpace):
			if m.multiple {
				// 多选下是选中
				m.toggleCheck()
			}
		case key.Matches(msg, keyHelp):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, keyQuit):
			fmt.Println("Bye!")
			os.Exit(1)
			return m, nil
		}
	case tea.WindowSizeMsg:
		//h, v := docStyle.GetFrameSize()
		//m.list.SetHeight(msg.Height - h)
		//m.list.SetWidth(msg.Width - v)
	}
	return m, nil
}

func (m Model) View() string {
	var str string
	if m.handler.onResultView {
		str = m.resultView()
	} else {
		str = m.selectView()
	}
	return docStyle.Render(str + "\n\n" + m.help.View(keyMap{
		showLike: m.likeFn != nil,
		multiple: m.multiple,
	}))
}

func (m Model) GetSelectedValues() []interface{} {
	var values []interface{}
	for _, o := range m.options {
		if o.IsChecked {
			values = append(values, o.Value)
		}
	}
	return values
}

func (m Model) toggleCheck() {
	m.options[m.idx].ToggleCheck()
}

func (m Model) selectView() string {
	var (
		str      string
		showDesc bool
	)
	str += titleBlockStyle.Render(m.title) + "\n"

	for _, o := range m.options {
		if o.Desc != "" {
			showDesc = true
			break
		}
	}
	for i, o := range m.options {
		str += o.display(i == m.idx, showDesc, m.multiple, m.itemStyles) + "\n"
	}
	return str
}

func (m Model) resultView() string {
	// todo
	return ""
}
