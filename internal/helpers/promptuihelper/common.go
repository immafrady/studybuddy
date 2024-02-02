package promptuihelper

// Option 下拉选项
type Option[T any] struct {
	Label   string
	Value   T
	Detail  string
	Checked bool
}

func (m *Option[T]) onCheck() {
	m.Checked = !m.Checked
}

type SelectConfig struct {
	Label         any
	LabelTemplate string
	Selected      string
	Details       string
	Like          bool   // 标记为喜欢
	OnLikeClicked func() // 标记为喜欢的方法 todo
	ConfirmText   string // 确认按钮的文字
}
