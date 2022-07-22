package word

import (
	"fmt"
	wordModel "myapp/model/classification/word"
	"myapp/pkg/gorm"
)

func GetWordList() []wordModel.Word {
	db, _ := gorm.New()

	var wordList []wordModel.Word
	result := db.Limit(10).Offset(0).Find(&wordList)
	if result.Error != nil {
		fmt.Println("List fail")
		panic(result.Error)
	}

	return wordList
}
