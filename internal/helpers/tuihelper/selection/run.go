package selection

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
)

type Config struct {
	Options      []*Option   // 下拉选项
	Title        string      // 标题
	Multiple     bool        // 是否为多选
	Liked        bool        // 是否为收藏
	ToggleLikeFn func() bool // 标记收藏的方法
	Index        int         // 起始下标
	ShowResult   bool
}

func Run(config Config) (m Model) {
	p := tea.NewProgram(Model{
		help: help.New(),
		handler: &handler{
			idx:        config.Index,
			l:          len(config.Options),
			showResult: config.ShowResult,
		},
		itemStyles:   NewItemStyles(),
		options:      config.Options,
		liked:        config.Liked,
		toggleLikeFn: config.ToggleLikeFn,
		multiple:     config.Multiple,
		title:        config.Title,
	}, tea.WithAltScreen())
	ret, err := p.Run()
	errorhelper.ExitOnError(err)
	if m, ok := ret.(Model); ok {
		return m
	}
	return
}
