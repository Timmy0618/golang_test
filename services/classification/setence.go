package classification

import (
	"encoding/json"
	"fmt"
	wordModel "myapp/model/classification/word"
	"strings"
)

type Sentence struct {
	UserId   int64  `json:UserId`
	Sentence string `json:Sentence`
}

func Classify(input []byte, wordList []wordModel.Word) {
	var sentence Sentence
	err := json.Unmarshal(input, &sentence)
	if err != nil {
		panic(err)
	}
	fmt.Println(sentence.Sentence)
	for _, word := range wordList {
		fmt.Println(word.Word)
		if strings.Contains(sentence.Sentence, word.Word) {
			fmt.Println("比對成功")
			fmt.Println(word.Word)
			fmt.Println(sentence.UserId)
			fmt.Println(sentence.Sentence)
		} else {
			fmt.Println("未比對到")
			fmt.Println(word.Word)
			fmt.Println(sentence.UserId)
			fmt.Println(sentence.Sentence)
		}

	}
}
