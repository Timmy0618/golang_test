package classification

import (
	"encoding/json"
	"fmt"
	classificationModel "myapp/model/classification"
	wordModel "myapp/model/classification/word"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type classify struct {
	db *gorm.DB
}
type Sentence struct {
	UserId   int64  `json:UserId`
	Sentence string `json:Sentence`
}

type Result struct {
	UserId  int64
	GroupId int64
	Weight  int64
}

func New(db *gorm.DB) *classify {
	return &classify{db}
}

func (p *classify) Classify(input []byte, wordList []wordModel.Word) classificationModel.Classification {
	var sentence Sentence
	err := json.Unmarshal(input, &sentence)
	if err != nil {
		panic(err)
	}
	fmt.Println(sentence.UserId)
	fmt.Println(sentence.Sentence)
	fmt.Println(wordList)
	for _, word := range wordList {
		fmt.Println("比對字：", word.Word)
		if strings.Contains(sentence.Sentence, word.Word) {
			fmt.Println("比對成功")
			fmt.Println("權重：", word.Weight)
			fmt.Println("GroupId:", word.GroupID)
			input := classificationModel.Classification{UserId: sentence.UserId, GroupId: word.GroupID, Score: word.Weight}
			p.UpdateUser(&input)
			return input
		} else {
			fmt.Println("未比對到")
		}
	}

	return classificationModel.Classification{UserId: 0, GroupId: 0}
}

func (p *classify) UpdateUser(classification *classificationModel.Classification) *classificationModel.Classification {
	fmt.Println(classification)
	result := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "group_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"score": gorm.Expr("score + ?", classification.Score)}),
	}).Create(&classification)

	if result.Error != nil {
		panic(result.Error.Error())
	}

	return classification
}
