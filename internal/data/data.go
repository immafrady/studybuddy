package data

import _ "embed"

//go:embed marx/judge.json
var MarxJudgeJson []byte

//go:embed marx/multiple.json
var MarxMultipleJson []byte

//go:embed marx/single.json
var MarxSingleJson []byte
