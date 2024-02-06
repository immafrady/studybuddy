package tuihelper

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap interface {
	help.KeyMap
	GetKeyPair(msg tea.KeyMsg) (KeyPair, error)
}

type KeyPair struct {
	Binding  key.Binding
	Callback func(model tea.Model) (tea.Model, tea.Cmd)
}
