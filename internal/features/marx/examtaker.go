package marx

import (
	"fmt"
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/service/quiz"
)

type AnswerSheet struct {
	features.ExamTaker // 指定继承了这个
	*model.Question
	Correct  bool
	index    int
	total    int
	pipeline features.Pipeline[string]
}

func NewExamTaker(question *model.Question, curr int, total int) features.ExamTaker {
	var pipeline features.Pipeline[string]
	switch question.Type {
	case model.QuestionSingle:
		pipeline = singlePipeline{}
	case model.QuestionMultiple:
		pipeline = multiplePipeline{}
	case model.QuestionJudge:
		pipeline = judgePipeline{}
	}
	return &AnswerSheet{
		Question: question,
		index:    curr,
		total:    total,
		pipeline: pipeline,
	}
}

func (a *AnswerSheet) TakeExam() bool {
	label := a.FormatLabel()

	options := a.pipeline.ParseOption(a.Detail)
	correct := a.pipeline.DoTask(options, a)
	quiz.MarkQuestionDone(a.Question, correct)
	if correct {
		fmt.Println(label, "对了！")
	} else {
		fmt.Println(label, "错了！")
	}
	return correct
}

func (a *AnswerSheet) FormatLabel() string {
	// (1/10:单选题) 我是问题我是问题
	return fmt.Sprintf("(%v/%v:%v) %v ", a.index, a.total, model.QuestionTypeLabelMap[a.Type], a.Q)
}

func (a *AnswerSheet) GetAnswer() string {
	return a.A
}
