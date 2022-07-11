package group

import (
	"fmt"
	groupModel "myapp/model/classification/group"
	"strconv"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type group struct {
	db *gorm.DB
}

type response struct {
	Msg  string
	Code int
}

func New(db *gorm.DB) *group {
	return &group{db}
}

func (p *group) Create(ctx iris.Context) {
	var w groupModel.UserClassificationGroup

	err := ctx.ReadJSON(&w)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Group creation failure").DetailErr(err))
		return
	}

	result := p.db.Create(&w)
	if result.Error != nil {
		fmt.Println("Create fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response{Msg: "Create Fail", Code: 500})
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(response{Msg: "Create Success", Code: 200})
}

func (p *group) List(ctx iris.Context) {
	page, err := strconv.Atoi(ctx.Params().Get("page"))
	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response{Msg: "Input fail", Code: 500})
	}

	var w []groupModel.UserClassificationGroup
	result := p.db.Limit(10).Offset(page - 1).Find(&w)
	if result.Error != nil {
		fmt.Println("List fail")
		// ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response{Msg: "List fail", Code: 500})
	}
	ctx.StatusCode(iris.StatusAccepted)

	ctx.JSON(w)
}

func (p *group) Read(ctx iris.Context) {

}

func (p *group) Update(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response{Msg: "Input fail", Code: 500})
	}

	var w groupModel.UserClassificationGroup

	if ctx.ReadJSON(&w) != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Word Update failure").DetailErr(err))
		return
	}

	w.ID = id

	result := p.db.
		Model(&w).
		Updates(groupModel.UserClassificationGroup{Name: w.Name})

	if result.Error != nil {
		fmt.Println("Update fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response{Msg: "Update fail", Code: 500})
		return
	}

	if result.RowsAffected < 1 {
		fmt.Println("Update fail")
		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(response{Msg: "Update Nothing", Code: 200})
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response{Msg: "Update Success", Code: 200})
}

func (p *group) Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response{Msg: "Input fail", Code: 500})
	}

	var w groupModel.UserClassificationGroup

	w.ID = id
	fmt.Println(w)

	result := p.db.
		Delete(&w)

	if result.Error != nil {
		fmt.Println("Delete fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response{Msg: "Delete fail", Code: 500})
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response{Msg: "Delete Success", Code: 200})
}
