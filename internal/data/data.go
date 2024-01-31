package data

import _ "embed"

// MarxJudgeJson 马克思判断题
//
//go:embed marx/judge.json
var MarxJudgeJson []byte

// MarxMultipleJson 马克思多选题
//
//go:embed marx/multiple.json
var MarxMultipleJson []byte

// MarxSingleJson 马克思单选题
//
//go:embed marx/single.json
var MarxSingleJson []byte
