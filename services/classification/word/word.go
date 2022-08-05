package word

import (
	"fmt"
	wordModel "myapp/model/classification/word"

	"gorm.io/gorm"
)

type wordService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *wordService {
	return &wordService{db}
}
func (p *wordService) GetWordList() []wordModel.Word {

	var wordList []wordModel.Word
	result := p.db.Limit(10).Offset(0).Find(&wordList)
	if result.Error != nil {
		fmt.Println("List fail")
		panic(result.Error)
	}
	fmt.Println(wordList)

	return wordList
}
