package examiner

import (
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/features/marx"
	"github.com/immafrady/studybuddy/internal/model"
)

type getExamTaker func(question *model.Question, curr int, total int) features.ExamTaker2

var examTakerMap = map[model.ClassName]getExamTaker{
	model.ClassMarx: marx.NewExamTaker2,
}
