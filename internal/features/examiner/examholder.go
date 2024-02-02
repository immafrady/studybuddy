package examiner

import (
	"fmt"
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/service/quiz"
	"gorm.io/gorm"
)

type Holder struct {
	db *gorm.DB
	*model.Classify
	*model.Record
	getExamTaker
	Types []model.QuestionType
	Limit int
}

// NewExamHolder 开始考试
func NewExamHolder(classify *model.Classify, types []model.QuestionType, limit int) *Holder {
	db, _ := database.Get()
	record := model.Record{
		ClassifyId:       classify.ID,
		QuestionIds:      make([]uint, 0),
		WrongQuestionIds: make([]uint, 0),
	}
	db.Create(&record)

	getExamTaker := examTakerMap[classify.Name]

	return &Holder{
		db:           db,
		Classify:     classify,
		Record:       &record,
		getExamTaker: getExamTaker,
		Types:        types,
		Limit:        limit,
	}
}

func (h Holder) Start() {
	fmt.Println("答题开始！")
	l := len(h.Classify.Questions)
	for i, q := range h.Classify.Questions {
		taker := h.getExamTaker(&q, i, l)
		correct := taker.TakeExam()
		quiz.AddAnswerRecord(h.Record, q.ID, correct)
	}
	fmt.Println("答题结束！")
	wc := len(h.WrongQuestionIds)
	tc := len(h.QuestionIds)
	sc := tc - wc
	fmt.Printf("正确: %v; 错误: %v; 得分：%v", sc, wc, int(float64(sc)/float64(tc)*100))
}
