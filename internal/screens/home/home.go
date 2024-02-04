package home

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/immafrady/studybuddy/internal/dispatcher"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper"
	"os"
)

func Run() {
	items := []list.Item{
		item{title: "开始做题", desc: "随便刷刷，从做的少的题目开始", cb: func() {
			dispatcher.Start()
		}},
		item{title: "考试模式", desc: "来一场痛痛快快都决斗吧！", cb: func() {

		}},
		item{title: "复习模式", desc: "查漏补缺", cb: func() {

		}},
		item{title: "历史记录", desc: "看看曾经都做过些什么", cb: func() {

		}},
	}
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "请选择模式"

	ret, err := tuihelper.NewProgram(m).Run()
	errorhelper.ExitOnError(err)
	if m, ok := ret.(model); ok {
		if m.quitting {
			fmt.Println("Bye!")
			os.Exit(1)
		}
		if i, ok := m.list.SelectedItem().(item); ok {
			i.cb()
		}
	}

}
