package selection

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Option struct {
	Label     string
	Value     interface{}
	Desc      string
	IsChecked bool
	IsCorrect bool // 是否为正确的
}

func (o *Option) display(onHover bool, showDesc bool, multiple bool, itemStyle DefaultItemStyles) string {
	var (
		title, desc           string
		titleStyle, descStyle lipgloss.Style
	)
	if multiple {
		if o.IsChecked {
			title = "[X] " + o.Label
		} else {
			title = "[ ] " + o.Label
		}
		desc = "    " + o.Desc
	} else {
		title = o.Label
		desc = o.Desc
	}

	if onHover {
		titleStyle = itemStyle.SelectedTitle
		descStyle = itemStyle.SelectedDesc
	} else {
		titleStyle = itemStyle.NormalTitle
		descStyle = itemStyle.NormalDesc
	}

	title = titleStyle.Render(title)
	if showDesc {
		desc = descStyle.Render(desc)
		return strings.Join([]string{title, desc}, "\n")
	} else {
		return title
	}
}

func (o *Option) displayResult() {

}

func (o *Option) ToggleCheck() {
	o.IsChecked = !o.IsChecked
}
