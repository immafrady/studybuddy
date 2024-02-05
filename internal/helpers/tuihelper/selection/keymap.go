package selection

import "github.com/charmbracelet/bubbles/key"

var (
	keyUp = key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "向上"),
	)
	keyDown = key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "向下"),
	)
	keyQuit = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "退出"),
	)
	keyHelp = key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "展示帮助信息"),
	)
	keyLike = key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "标记为关注对象"),
	)
	keyEnter = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("回车(enter)", "提交"),
	)
	keySpace = key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("空格(space)", "选择当前项"),
	)
)

type keyMap struct {
	showLike bool
	multiple bool
}

func (k keyMap) getCustomKeys() []key.Binding {
	var list []key.Binding
	list = append(list, keyEnter)
	if k.multiple {
		list = append(list, keySpace)
	}
	if k.showLike {
		list = append(list, keyLike)
	}
	return list
}

func (k keyMap) ShortHelp() []key.Binding {
	return append(k.getCustomKeys(), keyHelp)
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.getCustomKeys(),
		{keyHelp, keyQuit},
		{keyUp, keyDown},
	}
}
