package list

import (
	"github.com/charmbracelet/bubbles/list"
)

// Option 单个选项
type Option struct {
	list.Item
	list.DefaultItem
	Label   string // 标签
	Value   string // 值
	Desc    string // 描述
	Checked bool   // 是否被选中
}

func (o Option) FilterValue() string {
	return o.Label
}

func (o Option) Title() string {
	if o.Checked {
		return "◉ " + o.Label
	} else {
		return "  " + o.Label
	}
}
func (o Option) Description() string {
	return "  " + o.Desc
}

// ToggleCheck 切换是否被选中
func (o Option) ToggleCheck() {
	o.Checked = !o.Checked
}
