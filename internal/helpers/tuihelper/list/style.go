package list

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
)

var (
	docStyle   = lipgloss.NewStyle().Margin(1, 2)
	itemStyles = list.NewDefaultItemStyles()
)

func init() {
	//itemStyles.NormalTitle.PaddingLeft(3)
	//itemStyles.NormalDesc.PaddingLeft(3)
	//itemStyles.SelectedTitle.PaddingLeft(2)
	//itemStyles.SelectedDesc.PaddingLeft(2)
	//itemStyles.DimmedTitle.PaddingLeft(3)
	//itemStyles.DimmedDesc.PaddingLeft(3)
}

type OptionDelegate struct {
	ShowDescription bool
	Styles          list.DefaultItemStyles
	height          int
	spacing         int
}

func NewOptionDelegate(options []Option) OptionDelegate {
	showDesc := false
	for _, v := range options {
		if v.Description() != "" {
			showDesc = true
			break
		}
	}
	return OptionDelegate{
		ShowDescription: showDesc,
		Styles:          itemStyles,
		height:          2,
		spacing:         1,
	}
}

func (d OptionDelegate) Height() int                             { return 1 }
func (d OptionDelegate) Spacing() int                            { return 1 }
func (d OptionDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d OptionDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {

	var (
		title, desc string
		checked     bool
		selected    = index == m.Index()
	)

	if i, ok := item.(Option); ok {
		title = i.Title()
		desc = i.Description()
		checked = i.Checked
	} else {
		return
	}

	if checked {
		title = "◉ " + title
	} else {
		title = "◎ " + title
	}
	desc = "  " + desc

	if selected {
		title = itemStyles.SelectedTitle.Render(title)
		desc = itemStyles.SelectedDesc.Render(desc)
	} else {
		title = itemStyles.NormalTitle.Render(title)
		desc = itemStyles.NormalDesc.Render(desc)
	}

	if d.ShowDescription {
		fmt.Fprintf(w, "%s\n%s", title, desc)
	} else {
		fmt.Fprintf(w, "%s", title)
	}

}
