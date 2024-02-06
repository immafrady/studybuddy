package examiner

import (
	"github.com/immafrady/studybuddy/internal/dispatcher/ctx"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/screens"
	"github.com/immafrady/studybuddy/internal/service/quiz"
	"gorm.io/gorm"
)

type Examiner struct {
	Ctx        *ctx.Context
	ShowResult bool
}

func (e Examiner) StartExam() {
	// 构建
	quiz.FetchQuestionList(e.Ctx.Classify, func(db *gorm.DB) *gorm.DB {
		return db.Where("type IN (?)", e.Ctx.Types).Order("count").Limit(e.Ctx.Limit)
	})
	e.Ctx.Record = &model.Record{
		ClassifyId:       e.Ctx.Classify.ID,
		QuestionIds:      make([]uint, 0),
		WrongQuestionIds: make([]uint, 0)}
	e.Ctx.History = make([]tuihelper.ResultOutput, 0)
	e.Ctx.DB.Create(e.Ctx.Record)

	getExamTaker := examTakerMap[e.Ctx.Classify.Name]

	// 做题
	l := len(e.Ctx.Classify.Questions)
	for i, q := range e.Ctx.Classify.Questions {
		examTaker := getExamTaker(&q, i, l)
		m := examTaker.DoExam(e.ShowResult)
		e.Ctx.History = append(e.Ctx.History, &m)
		// 标记记录
		correct := m.AllSelectMatched()
		quiz.AddAnswerRecord(e.Ctx.Record, q.ID, correct)
		quiz.MarkQuestionDone(&q, correct)
	}
	// 结果
	screens.QuizResultRun(screens.QuizResultRunArgs{
		Classify: e.Ctx.Classify,
		History:  e.Ctx.History,
		RedoFn:   nil,
	})
}
