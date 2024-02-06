package screens

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/viewport"
	"github.com/immafrady/studybuddy/internal/model"
	"strings"
)

type QuizResultRunArgs struct {
	Classify *model.Classify
	History  []tuihelper.ResultOutput
	RedoFn   func() // 重新开始的方法
}

func QuizResultRun(args QuizResultRunArgs) {
	title := "查看结果 - " + args.Classify.Name

	strs := make([]string, len(args.History))
	for i, h := range args.History {
		strs[i] = h.ResultView()
	}
	content := strings.Join(strs, "\n\n")

	viewport.Run(viewport.Config{
		Title:   string(title),
		Content: content,
		//KeyMap:  keyMap{},
	})
}

type keyMap struct {
}

func (k keyMap) ShortHelp() []key.Binding {
	//TODO implement me
	panic("implement me")
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}

func (k keyMap) Update(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	//TODO implement me
	panic("implement me")
}
