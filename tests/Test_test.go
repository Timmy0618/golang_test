package tests

import (
	"net/http"
	"testing"

	"github.com/kataras/iris/v12"

	"myapp/router/classification/word"

	"github.com/gavv/httpexpect/v2"
)

func irisTester(route *iris.Application, t *testing.T) *httpexpect.Expect {
	handler := route

	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomething(t *testing.T) {
	test := New()
	e := irisTester(word.GetRoute(test.route, test.db, test.rmq, test.rdb), t)

	schema := `{
		"type": "array",
		"items": {
			"type": "object",
			"properties": {
				"name":        {"type": "string"},
				"description": {"type": "string"}
			},
			"required": ["name", "description"]
		}
	}`

	things := e.GET("/classification/word/1").
		Expect().
		Status(http.StatusOK).JSON()

	things.Schema(schema)

	// names := things.Path("$[*].name").Array()

	// names.Elements("foo", "bar")

	// for n, desc := range things.Path("$..description").Array().Iter() {
	// 	m := desc.String().Match("(.+) (.+)")

	// 	m.Index(1).Equal(names.Element(n).String().Raw())
	// 	m.Index(2).Equal("thing")
	// }

}
