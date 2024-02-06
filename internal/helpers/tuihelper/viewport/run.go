package viewport

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper"
)

type Config struct {
	Title   string
	Content string
	KeyMap  tuihelper.KeyMap
}

func Run(config Config) (m Model) {
	p := tea.NewProgram(Model{
		KeyMap:  config.KeyMap,
		title:   config.Title,
		content: config.Content,
	},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion())
	ret, err := p.Run()
	errorhelper.ExitOnError(err)
	if m, ok := ret.(Model); ok {
		return m
	}
	return
}
