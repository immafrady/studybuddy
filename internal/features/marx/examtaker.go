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
	return fmt.Sprintf("(%v/%v:%v) %v ", a.index+1, a.total, model.QuestionTypeLabelMap[a.Type], a.Q)
}

func (a *AnswerSheet) GetAnswer() string {
	return a.A
}

func (a *AnswerSheet) DisplayResult(options []promptuihelper.Option[string], ans string, correct bool) {
	fmt.Println("\n-------------")
	var marker *color.Color
	if correct {
		marker = color.New(color.FgGreen)
		colorhelper.RightColor.Print("✓ ")
	} else {
		marker = color.New(color.BgRed)
		colorhelper.WrongColor.Print("✗ ")
	}
	fmt.Print(a.GetLabel(), "\n")
	for _, v := range options {
		display := " " + v.Label
		selected := strings.Contains(ans, v.Value)
		if selected {
			display = "●" + display
		} else {
			display = " " + display
		}
		if correct {
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
