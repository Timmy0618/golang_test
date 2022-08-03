package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	groupModel "myapp/model/classification/group"
	groupRouter "myapp/router/classification/group"
	testDefault "myapp/tests"
)

type group struct {
	Name string
}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestCreate(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(groupRouter.GetRoute(test.Route, test.DB), t)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"}
			},
			"required": ["Code", "Msg"]
	}`

	fake := group{
		Name: "騎空團"}

	createResult := e.POST("/classification/group").
		WithJSON(fake).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	createResult.Schema(schema)
	var g groupModel.Group
	test.DB.
		Where("name = ?", fake.Name).
		Find(&g)

	assert.Equal(t, fake.Name, g.Name, "they should be equal")
}

func TestList(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(groupRouter.GetRoute(test.Route, test.DB), t)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"},
				"Data":		{"type": "array"}
			},
			"required": ["Code", "Msg", "Data"]
	}`

	things := e.GET("/classification/group").
		WithQuery("page", 1).
		Expect().
		Status(http.StatusAccepted).JSON()

	things.Schema(schema)
}

func TestRead(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(groupRouter.GetRoute(test.Route, test.DB), t)

	fake := groupModel.Group{
		Name: "戰隊戰"}
	test.DB.Create(&fake)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"},
				"Data":		{
					"type": "object",
					"properties": {
						"Name": {"type": "string"}
					}
				}
			},
			"required": ["Code", "Msg", "Data"]
	}`

	things := e.GET("/classification/group/{id}", fake.ID).
		Expect().
		Status(http.StatusOK).JSON()

	things.Schema(schema)
}

func TestUpdate(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(groupRouter.GetRoute(test.Route, test.DB), t)

	fake := groupModel.Group{
		Name: "戰隊戰"}

	test.DB.Create(&fake)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"}
			},
			"required": ["Code", "Msg"]
	}`

	groupInput := group{
		Name: "戰隊戰測試"}

	updateResult := e.PATCH("/classification/group/{id}", fake.ID).
		WithJSON(groupInput).
		Expect().
		Status(http.StatusOK).JSON()

	updateResult.Schema(schema)

	g := groupModel.Group{
		ID: fake.ID,
	}
	test.DB.Find(&g)

	assert.Equal(t, groupInput.Name, g.Name, "they should be equal")
}

func TestDelete(t *testing.T) {
	test := testDefault.New()
	e := testDefault.IrisTester(groupRouter.GetRoute(test.Route, test.DB), t)

	fake := groupModel.Group{
		Name: "戰隊戰測試"}

	test.DB.Create(&fake)

	schema := `{
			"type": "object",
			"properties": {
				"Code":     {"type": "integer"},
				"Msg": 		{"type": "string"}
			},
			"required": ["Code", "Msg"]
	}`

	deleteResult := e.DELETE("/classification/group/{id}", fake.ID).
		Expect().
		Status(http.StatusOK).JSON()

	deleteResult.Schema(schema)

	g := groupModel.Group{
		ID: fake.ID,
	}
	result := test.DB.First(&g)
	assert.Equal(t, result.Error.Error(), "record not found", "they should be equal")
}
