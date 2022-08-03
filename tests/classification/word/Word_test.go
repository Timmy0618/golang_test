package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	wordModel "myapp/model/classification/word"
	wordRouter "myapp/router/classification/word"
	testDefault "myapp/tests"
)

type word struct {
	Word    string
	Weight  int64
	GroupID int64
}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestCreate(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(wordRouter.GetRoute(test.Route, test.DB, test.Rmq, test.Rdb), t)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"}
			},
			"required": ["Code", "Msg"]
	}`

	fake := word{
		Word:    "test",
		Weight:  1,
		GroupID: 5}

	createResult := e.POST("/classification/word").
		WithJSON(fake).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	createResult.Schema(schema)
	var w wordModel.Word
	test.DB.
		Where("word = ? AND weight = ? AND group_id = ?", fake.Word, fake.Weight, fake.GroupID).
		Find(&w)

	assert.Equal(t, fake.Word, w.Word, "they should be equal")
	assert.Equal(t, fake.Weight, w.Weight, "they should be equal")
	assert.Equal(t, fake.GroupID, w.GroupID, "they should be equal")
}

func TestList(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(wordRouter.GetRoute(test.Route, test.DB, test.Rmq, test.Rdb), t)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"},
				"Data":		{"type": "array"}
			},
			"required": ["Code", "Msg", "Data"]
	}`

	things := e.GET("/classification/word").
		WithQuery("page", 1).
		Expect().
		Status(http.StatusAccepted).JSON()

	things.Schema(schema)
}

func TestRead(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(wordRouter.GetRoute(test.Route, test.DB, test.Rmq, test.Rdb), t)

	fake := wordModel.Word{
		Word:    "test",
		Weight:  1,
		GroupID: 5}
	test.DB.Create(&fake)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"},
				"Data":		{"type": "array"}
			},
			"required": ["Code", "Msg", "Data"]
	}`

	things := e.GET("/classification/word/{id}", fake.ID).
		Expect().
		Status(http.StatusOK).JSON()

	things.Schema(schema)
}

func TestUpdate(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(wordRouter.GetRoute(test.Route, test.DB, test.Rmq, test.Rdb), t)

	fake := wordModel.Word{
		Word:    "test",
		Weight:  1,
		GroupID: 5}

	test.DB.Create(&fake)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"}
			},
			"required": ["Code", "Msg"]
	}`

	wordInput := word{
		Word:    "testUpdate",
		Weight:  1,
		GroupID: 5}

	updateResult := e.PATCH("/classification/word/{id}", fake.ID).
		WithJSON(wordInput).
		Expect().
		Status(http.StatusOK).JSON()

	updateResult.Schema(schema)

	w := wordModel.Word{
		ID: fake.ID,
	}
	test.DB.Find(&w)

	assert.Equal(t, wordInput.Word, w.Word, "they should be equal")
	assert.Equal(t, wordInput.Weight, w.Weight, "they should be equal")
	assert.Equal(t, wordInput.GroupID, w.GroupID, "they should be equal")
}

func TestDelete(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(wordRouter.GetRoute(test.Route, test.DB, test.Rmq, test.Rdb), t)

	fake := wordModel.Word{
		Word:    "test",
		Weight:  1,
		GroupID: 5}

	test.DB.Create(&fake)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"}
			},
			"required": ["Code", "Msg"]
	}`

	deleteResult := e.DELETE("/classification/word/{id}", fake.ID).
		Expect().
		Status(http.StatusOK).JSON()

	deleteResult.Schema(schema)

	w := wordModel.Word{
		ID: fake.ID,
	}
	result := test.DB.First(&w)
	assert.Equal(t, result.Error.Error(), "record not found", "they should be equal")
}
