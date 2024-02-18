package selection

import "github.com/charmbracelet/bubbles/help"

func (m Model) View() string {
	var str string
	if m.handler.onResultView {
		str = m.ResultView()
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

func (m Model) ResultView() string {
	str := m.formattedTitle()

	for _, o := range m.options {
		str += o.displayResult(m.multiple, m.showDesc(), m.itemStyles, m.AllSelectMatched()) + "\n"
	}
	return str
}

func (m Model) selectView() string {
	str := m.formattedTitle()

	for i, o := range m.options {
		str += o.display(m.multiple, i == m.idx, m.showDesc(), m.itemStyles) + "\n"
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
