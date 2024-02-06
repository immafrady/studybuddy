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
	options       []*Option
	liked         bool // todo
	toggleLikeFn  func() bool
	multiple      bool
	title         string
	beforeLeaveFn func(m Model) // todo 回调函数（防止闪屏）
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
			if m.toggleLikeFn != nil {
				m.liked = m.toggleLikeFn()
			}
		case key.Matches(msg, keyEnter):
			if m.onResultView { // 在结果页就直接离开
				return m, tea.Quit
			} else {
				if !m.multiple {
					// 单选场景下，选中值
					m.toggleCheck()
				}
				if m.showResult { // 如果要展示结果，就跳去结果页
					m.handler.onResultView = true
				} else { // 不需要展示结果就直接
					return m, tea.Quit
				}
			}
		case key.Matches(msg, keySpace):
			if !m.onResultView && m.multiple {
				// 非结果页，多选下是选中
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
	var km help.KeyMap
	if m.onResultView {
		km = keyMapResultView{
			showLike: m.toggleLikeFn != nil,
		}
	} else {
		km = keyMap{
			showLike: m.toggleLikeFn != nil,
			multiple: m.multiple,
		}
	}

	return docStyle.Render(str + "\n\n" + m.help.View(km))
}

// GetSelectedValues 获取选择的值
func (m Model) GetSelectedValues() []interface{} {
	var values []interface{}
	for _, o := range m.options {
		if o.IsChecked {
			values = append(values, o.Value)
		}
	}
	return values
}

// AllSelectMatched 所有勾选都匹配上答案？
func (m Model) AllSelectMatched() bool {
	for _, o := range m.options {
		if o.IsChecked != o.IsCorrect {
			return false
		}
	}
	return true
}

func (m Model) ResultView() string {
	return m.resultView()
}

func (m Model) showDesc() bool {
	showDesc := false
	for _, o := range m.options {
		if o.Desc != "" {
			showDesc = true
			break
		}
	}
	return showDesc
}

func (m Model) toggleCheck() {
	m.options[m.idx].ToggleCheck()
}

func (m Model) selectView() string {
	str := m.formattedTitle()

	for i, o := range m.options {
		str += o.display(m.multiple, i == m.idx, m.showDesc(), m.itemStyles) + "\n"
	}
	return str
}

func (m Model) resultView() string {
	str := m.formattedTitle()

	for _, o := range m.options {
		str += o.displayResult(m.multiple, m.showDesc(), m.itemStyles, m.AllSelectMatched()) + "\n"
	}
	return str
}

// formattedTitle 格式化Title
func (m Model) formattedTitle() string {
	symbol := " "
	if m.onResultView {
		symbol = symbolStatus
		if m.AllSelectMatched() {
			symbol = m.itemStyles.CorrectFg.Render(symbol)
		} else {
			symbol = m.itemStyles.WrongFg.Render(symbol)
		}
	}
	return titleBlockStyle.Render(symbol+" "+m.title) + "\n"
}
