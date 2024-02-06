package screens

import (
	"errors"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/immafrady/studybuddy/internal/dispatcher/ctx"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/viewport"
	"strings"
)

type QuizResultRunArgs struct {
	Ctx *ctx.Context
}

func QuizResultRun(args QuizResultRunArgs) {
	title := "查看结果 - " + args.Ctx.Classify.Name

	strs := make([]string, len(args.Ctx.History))
	for i, h := range args.Ctx.History {
		strs[i] = h.ResultView()
	}
	content := strings.Join(strs, "\n\n")

	viewport.Run(viewport.Config{
		Title:   string(title),
		Content: content,
		KeyMap: quizResultKeyMap{
			Restart: tuihelper.KeyPair{
				Binding: key.NewBinding(
					key.WithKeys("r"),
					key.WithHelp("r", "重新开始"),
				),
				Callback: func(model tea.Model) (tea.Model, tea.Cmd) {
					args.Ctx.State = ctx.DoAgain
					return model, tea.Quit
				},
			},
			StartAllOver: tuihelper.KeyPair{
				Binding: key.NewBinding(
					key.WithKeys("enter"),
					key.WithHelp("回车(enter)", "回到首页"),
				),
				Callback: func(model tea.Model) (tea.Model, tea.Cmd) {
					args.Ctx.State = ctx.StartAllOver
					return model, tea.Quit
				},
			},
			Quit: tuihelper.KeyPair{
				Binding: key.NewBinding(
					key.WithKeys("ctrl+c", "q", "esc"),
					key.WithHelp("q", "退出"),
				),
				Callback: func(model tea.Model) (tea.Model, tea.Cmd) {
					args.Ctx.State = ctx.DoNothing
					return model, tea.Quit
				},
			},
		},
	})
}

type quizResultKeyMap struct {
	Restart      tuihelper.KeyPair
	StartAllOver tuihelper.KeyPair
	Quit         tuihelper.KeyPair
}

func (k quizResultKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Restart.Binding,
		k.StartAllOver.Binding,
		k.Quit.Binding,
	}
}

func (k quizResultKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}

func (k quizResultKeyMap) GetKeyPair(msg tea.KeyMsg) (kp tuihelper.KeyPair, err error) {
	switch {
	case key.Matches(msg, k.Restart.Binding):
		return k.Restart, nil
	case key.Matches(msg, k.StartAllOver.Binding):
		return k.StartAllOver, nil
	case key.Matches(msg, k.Quit.Binding):
		return k.Quit, nil
	}
	return kp, errors.New("KeyPair not found")
}
