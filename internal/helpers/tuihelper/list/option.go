package list

import "github.com/charmbracelet/bubbles/list"

// Option 单个选项
type Option struct {
	list.Item
	Label   string // 标签
	Value   string // 值
	Desc    string // 描述
	Checked bool   // 是否被选中
}

func (o Option) FilterValue() string {
	return o.Label
}

func (o Option) Title() string {
	return o.Label
}

// ToggleCheck 切换是否被选中
func (o Option) ToggleCheck() {
	o.Checked = !o.Checked
}

// Options 选项组
type Options []Option

func (os Options) FilterValue() string {
	return os[0].Label
}

// FilterChecked 过滤选中的
func (os Options) FilterChecked() Options {
	var newOs Options
	for _, o := range os {
		if o.Checked {
			newOs = append(newOs, o)
		}
	}
	return newOs
}
