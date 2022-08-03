package group

import (
	"fmt"
	groupModel "myapp/model/classification/group"
	"myapp/services/response"
	"strconv"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type group struct {
	db *gorm.DB
}

func New(db *gorm.DB) *group {
	return &group{db}
}

func (p *group) Create(ctx iris.Context) {
	var g groupModel.Group

	err := ctx.ReadJSON(&g)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Group creation failure").DetailErr(err))
		return
	}

	result := p.db.Create(&g)
	if result.Error != nil {
		fmt.Println("Create fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Create Fail"}))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(response.Responser(response.Response{Code: 200, Msg: "Create Fail"}))
}

func (p *group) List(ctx iris.Context) {
	page, err := strconv.Atoi(ctx.URLParam("page"))
	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Input Fail"}))
		return
	}

	var g []groupModel.Group
	result := p.db.Limit(10).Offset(page - 1).Find(&g)
	if result.Error != nil {
		fmt.Println("List fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "List Fail"}))
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response.Responser(response.Response{Code: 200, Msg: "Success", Data: g}))
}

func (p *group) Read(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Input Fail"}))
		return
	}

	g := groupModel.Group{
		ID: id,
	}

	result := p.db.First(&g)
	if result.Error != nil {
		fmt.Println("Read fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Input Fail"}))
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(response.Responser(response.Response{Code: 200, Msg: "Success", Data: g}))
}

func (p *group) Update(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Input Fail"}))
		return
	}

	g := groupModel.Group{
		ID: id,
	}

	if ctx.ReadJSON(&g) != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Word Update failure").DetailErr(err))
		return
	}

	result := p.db.
		Model(&g).
		Updates(groupModel.Group{Name: g.Name})

	if result.Error != nil {
		fmt.Println("Update fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Update Fail"}))
		return
	}

	if result.RowsAffected < 1 {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(response.Responser(response.Response{Code: 200, Msg: "Update Nothing"}))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(response.Responser(response.Response{Code: 200, Msg: "Update Success"}))
}

func (p *group) Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Input Fail"}))
	}

	g := groupModel.Group{
		ID: id,
	}

	result := p.db.Delete(&g)

	if result.Error != nil {
		fmt.Println("Delete fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(response.Response{Code: 500, Msg: "Delete Fail"}))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(response.Responser(response.Response{Code: 200, Msg: "Delete Success"}))
}
