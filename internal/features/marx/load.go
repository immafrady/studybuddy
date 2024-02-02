package marx

import (
	"encoding/json"
	"github.com/immafrady/studybuddy/internal/data"
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/immafrady/studybuddy/internal/model"
	"gorm.io/gorm"
)

// LoadData 加载题目
func LoadData() {
	db, _ := database.Get()
	classify := model.Classify{Name: model.ClassMarx}
	// 如果找到一条，说明有初始化了
	if db.Limit(1).Find(&classify).RowsAffected != 1 {
		// 添加分类
		result := db.Create(&classify)
		errorhelper.ExitOnError(result.Error)
		// 添加具体题目
		tasks := map[model.QuestionType][]byte{
			model.QuestionSingle:   data.MarxSingleJson,
			model.QuestionMultiple: data.MarxMultipleJson,
			model.QuestionJudge:    data.MarxJudgeJson,
		}

		for t, jsonBytes := range tasks {
			var qs []rawQuiz
			errorhelper.ExitOnError(json.Unmarshal(jsonBytes, &qs))
			var questions = make([]*model.Question, len(qs))
			for i, quiz := range qs {
				questions[i] = quiz.toQuestion(t, &classify)
			}
			result := db.Create(questions)
			errorhelper.ExitOnError(result.Error)
		}
	}
}

// rawQuiz 数据源的结构
type rawQuiz struct {
	Problem string `json:"problem"`
	Answer  string `json:"answer"`
	A       string `json:"A,omitempty"`
	B       string `json:"B,omitempty"`
	C       string `json:"C,omitempty"`
	D       string `json:"D,omitempty"`
}

func (q rawQuiz) toQuestion(questionType model.QuestionType, classify *model.Classify) *model.Question {
	selection := fromRawQuiz(q)
	return &model.Question{
		Model:      gorm.Model{},
		ClassifyId: classify.ID,
		Q:          q.Problem,
		A:          q.Answer,
		Detail:     selection.jsonStringify(),
		Like:       false,
		Count:      0,
		Type:       questionType,
	}
}
