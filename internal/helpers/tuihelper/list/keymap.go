package list

import "github.com/charmbracelet/bubbles/key"

var (
	keyUp = key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑/k", "向右"),
	)
	keyDown = key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓/j", "向下"),
	)
	keyLeft = key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←/h", "左移"),
	)
	keyRight = key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→/l", "右移"),
	)
	keyQuit = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "退出"),
	)
	keyHelp = key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	)
	keyLike = key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "标记为关注对象"),
	)
	keyEnter = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("回车", "确认选择"),
	)
	keySpace = key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("空格", "选择"),
	)
)

type keyMap struct {
	showLike bool
	multiple bool
}

func (k keyMap) getCustomKeys() []key.Binding {
	var list []key.Binding
	if k.multiple {
		list = append(list, keySpace, keyEnter)
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
		{keyUp, keyDown, keyLeft, keyRight},
	}
}
