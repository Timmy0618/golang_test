package tests

import (
	"encoding/json"
	"fmt"
	"myapp/services/classification"
	wordService "myapp/services/classification/word"
	testDefault "myapp/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassify(t *testing.T) {
	fake := classification.Sentence{
		UserId:   123,
		Sentence: "test",
	}
	in, _ := json.Marshal(fake)

	test := testDefault.New()
	wordService := wordService.New(test.DB)
	classification := classification.New(test.DB)
	result := classification.Classify(in, wordService.GetWordList())
	fmt.Println(result)

	test.DB.Find(&result)

	assert.Equal(t, result.UserId, fake.UserId, "they should be equal")
	assert.NotEqual(t, result.GroupId, int64(0), "they should not be equal")

}
