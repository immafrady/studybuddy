package selection

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle        = lipgloss.NewStyle().Margin(1, 2)
	titleBlockStyle = lipgloss.NewStyle().Padding(0, 0, 2, 0)
	itemBlockStyle  = lipgloss.NewStyle().PaddingBottom(2)
)

type DefaultItemStyles struct {
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style
	NormalTitle   lipgloss.Style
	NormalDesc    lipgloss.Style
	CorrectBg     lipgloss.Style
	CorrectFg     lipgloss.Style
	WrongBg       lipgloss.Style
	WrongFg       lipgloss.Style
}

func NewItemStyles() (s DefaultItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		PaddingLeft(3)

	s.NormalDesc = s.NormalTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
		PaddingLeft(2)

	s.SelectedDesc = s.SelectedTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	sc := lipgloss.Color("46")
	s.CorrectBg = lipgloss.NewStyle().Background(sc)
	s.CorrectFg = lipgloss.NewStyle().Foreground(sc)

	wc := lipgloss.Color("160")
	s.WrongBg = lipgloss.NewStyle().Background(wc)
	s.WrongFg = lipgloss.NewStyle().Foreground(wc)
	return
}