package tuihelper

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// DocStyle 同一的文档样式
var DocStyle = lipgloss.NewStyle().Margin(1, 2)

// NewProgram 同一的程序启动
func NewProgram(model tea.Model) *tea.Program {
	return tea.NewProgram(model, tea.WithAltScreen())
}

// CalculateWindowSizeMsg 计算宽高
func CalculateWindowSizeMsg(msg tea.Msg) (h int, v int) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		h, v = DocStyle.GetFrameSize()
		h = msg.Width - h
		v = msg.Height - v
	}
	return
}
