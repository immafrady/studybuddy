package marx

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/helpers/colorhelper"
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/service/quiz"
	"strings"
)

type AnswerSheet struct {
	features.ExamTaker // 指定继承了这个
	*model.Question
	correct  bool                            // 是否正确
	ans      string                          // 用户自己写的答案
	options  []promptuihelper.Option[string] // 下拉选项
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
	a.options = a.pipeline.ParseOption(a.Detail)
	a.ans, a.correct = a.pipeline.DoTask(a.options, a)

	quiz.MarkQuestionDone(a.Question, a.correct)
	a.DisplayResult()
	return a.correct
}

func (a *AnswerSheet) GetLabel() string {
	// (1/10:单选题) 我是问题我是问题
	return fmt.Sprintf("(%v/%v:%v) %v ", a.index+1, a.total, model.QuestionTypeLabelMap[a.Type], a.Q)
}

func (a *AnswerSheet) GetAnswer() string {
	return a.A
}

func (a *AnswerSheet) DisplayResult() {
	fmt.Println("\n-------------")
	var marker *color.Color
	if a.correct {
		marker = color.New(color.FgGreen)
		colorhelper.RightColor.Print("✓ ")
	} else {
		marker = color.New(color.BgRed)
		colorhelper.WrongColor.Print("✗ ")
	}
	fmt.Print(a.GetLabel(), "\n")
	for _, v := range a.options {
		display := " " + v.Label
		selected := strings.Contains(a.ans, v.Value)
		if selected {
			display = "●" + display
		} else {
			display = " " + display
		}
		if a.correct {
			if selected {
				marker.Println(display)
			} else {
				fmt.Println(display)
			}
		} else {
			if strings.Contains(a.A, v.Value) {
				marker.Println(display)
			} else {
				fmt.Println(display)
			}
		}
	}
	fmt.Println("-------------")

}
