package selection

type handler struct {
	idx          int // 当前选择下标
	l            int
	onResultView bool // 是否在结果页
	showResult   bool // 是否展示结果
}

func (h *handler) MoveUp() {
	if h.idx > 0 {
		h.idx--
	}
}

func (h *handler) MoveDown() {
	if h.idx < h.l-1 {
		h.idx++
	}
}
