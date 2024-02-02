package marx

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
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
	options := a.pipeline.ParseOption(a.Detail)
	ans, correct := a.pipeline.DoTask(options, a)
	quiz.MarkQuestionDone(a.Question, correct)
	a.DisplayResult(options, ans, correct)
	return correct
}

func (a *AnswerSheet) GetLabel() string {
	// (1/10:单选题) 我是问题我是问题
	return fmt.Sprintf("(%v/%v:%v) %v ", a.index, a.total, model.QuestionTypeLabelMap[a.Type], a.Q)
}

func (a *AnswerSheet) GetAnswer() string {
	return a.A
}

func (a *AnswerSheet) DisplayResult(options []promptuihelper.Option[string], ans string, correct bool) {
	fmt.Println("\n-------------")
	fmt.Println(a.GetLabel())
	for _, v := range options {
		fmt.Println(v.Label)
	}
	if correct {
		color.Set(color.FgGreen)
		fmt.Println("回答正确：", ans)
		color.Unset()
	} else {
		fmt.Print("回答错误：")
		color.Set(color.FgRed)
		fmt.Print(ans, "\n")
		color.Unset()
		fmt.Printf("正确答案：%v\n", a.A)
	}
	fmt.Println("-------------")

}
