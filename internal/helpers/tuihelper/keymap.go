package tuihelper

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap interface {
	help.KeyMap
	Update(msg tea.KeyMsg) (tea.Model, tea.Cmd)
}

//type KeyPair struct {
//	Binding key.Binding
//	Update func()
//}
