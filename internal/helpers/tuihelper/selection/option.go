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

func (o *Option) display(multiple bool, onHover bool, showDesc bool, itemStyle DefaultItemStyles) string {
	var (
		titleStyle, descStyle lipgloss.Style
	)
	title, desc := o.format(onHover, multiple)

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

func (o *Option) displayResult(multiple bool, showDesc bool, itemStyle DefaultItemStyles, isSuccess bool) string {
	title, desc := o.format(false, multiple)
	title = itemStyle.NormalTitle.Render(title)
	desc = itemStyle.NormalDesc.Render(desc)
	if o.IsCorrect {
		var bgStyle lipgloss.Style
		if isSuccess {
			bgStyle = itemStyle.CorrectBg
		} else {
			bgStyle = itemStyle.WrongBg
		}
		title = bgStyle.Render(title)
		desc = bgStyle.Render(desc)
	}

	if showDesc {
		return strings.Join([]string{title, desc}, "\n")
	} else {
		return title
	}
}

func (o *Option) format(onHover bool, multiple bool) (title string, desc string) {
	prefixer := func(str string) string {
		if multiple {
			return "[" + str + "] "
		} else {
			return " " + str + "  "
		}
	}
	if o.IsChecked {
		title = prefixer(symbolSolid) + o.Label
	} else if onHover {
		title = prefixer(symbolHollow) + o.Label
	} else {
		title = prefixer(" ") + o.Label
	}
	desc = strings.Repeat(" ", len(prefixer(" "))) + o.Desc
	return
}

func (o *Option) ToggleCheck() {
	o.IsChecked = !o.IsChecked
}
