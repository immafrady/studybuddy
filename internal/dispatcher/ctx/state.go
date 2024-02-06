package ctx

type State uint

const (
	DoNothing State = iota
	DoAgain
	StartAllOver
)
