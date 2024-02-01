package marx

import (
	"github.com/immafrady/studybuddy/internal/model"
)

// Quiz 问题的类型
type Quiz interface {
	Evaluate(answer string) bool
}

// NewQuiz 构造问题
func NewQuiz(question *model.Question) Quiz {
	base := baseQ{
		Question: question.Q,
		Answer:   question.A,
		Like:     question.Like,
		Count:    question.Count,
	}
	if question.Type == model.QuestionJudge {
		return &JudgeQuiz{
			base,
		}
	} else {
		s := parseSelection(question.Detail)
		switch question.Type {
		case model.QuestionSingle:
			return &SingleQuiz{
				baseQ:     base,
				selection: *s,
			}
		case model.QuestionMultiple:
			return &MultipleQuiz{
				baseQ:     base,
				selection: *s,
			}
		}
	}
	return nil
}

/* --- 基本类型 --- */

type baseQ struct {
	Question string
	Answer   string
	Like     bool
	Count    uint
}

/* --- 基本类型 end --- */

/* --- 具体类型 --- */

// MultipleQuiz 多选题
type MultipleQuiz struct {
	baseQ
	selection
}

func (m MultipleQuiz) Evaluate(answer string) bool {
	return answer == m.Answer
}

// SingleQuiz 单选题
type SingleQuiz struct {
	baseQ
	selection
}

func (s SingleQuiz) Evaluate(answer string) bool {
	return answer == s.Answer
}

// JudgeQuiz 判断题
type JudgeQuiz struct {
	baseQ
}

func (j JudgeQuiz) Evaluate(answer string) bool {
	return answer == j.Answer
}

/* --- 具体类型 end --- */
