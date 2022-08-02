package word

import (
	"encoding/json"
	"fmt"
	"log"
	wordModel "myapp/model/classification/word"
	redisService "myapp/services/redis"
	"myapp/services/response"
	"myapp/services/rmq"
	"strconv"

	"github.com/go-redis/redis/v9"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type word struct {
	db  *gorm.DB
	rmq *amqp.Connection
	rdb *redis.Client
}

func New(db *gorm.DB, rmq *amqp.Connection, rdb *redis.Client) *word {
	return &word{db, rmq, rdb}
}

func (p *word) Create(ctx iris.Context) {
	var w wordModel.Word

	err := ctx.ReadJSON(&w)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Input Fail").DetailErr(err))
		return
	}

	result := p.db.Create(&w)
	if result.Error != nil {
		fmt.Println("Create fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(500, "Create Fail"))
		ctx.StopWithProblem(iris.StatusBadGateway, iris.NewProblem().
			Title("Word creation failure").DetailErr(err))
		return
	}

	if redisService.Scan(p.rdb, "wordList") {
		redisService.Del(p.rdb, "wordList")
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(response.Responser(200, "Created Success"))

	line := fmt.Sprintf("%s input: %#v", ctx.Method(), w)
	ctx.Application().Logger().Info(line)
}

func (p *word) List(ctx iris.Context) {
	page, err := strconv.Atoi(ctx.URLParam("page"))
	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(500, "Input Fail"))
		return
	}

	var w []wordModel.Word
	result := p.db.Limit(10).Offset(page - 1).Find(&w)
	if result.Error != nil {
		fmt.Println("List fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(500, "Input fail"))
		return
	}
	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response.Responser(200, "Success", w))
}

func (p *word) Read(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(500, "Input Fail"))
		return
	}

	var w wordModel.Word
	w.ID = id

	result := p.db.Find(&w)
	if result.Error != nil {
		fmt.Println("Read fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(500, "Input fail"))
		return
	}
	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response.Responser(200, "Success", w))
}

func (p *word) Update(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(500, "Input fail"))
		return
	}

	var w wordModel.Word

	if ctx.ReadJSON(&w) != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Input fail").DetailErr(err))
		return
	}

	w.ID = id
	fmt.Println(w)

	result := p.db.
		Model(&w).
		Updates(wordModel.Word{Weight: w.Weight, Word: w.Word})

	if result.Error != nil {
		fmt.Println("Update fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(500, "Update fail"))
		return
	}

	if result.RowsAffected < 1 {
		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(response.Responser(200, "Nothing Change"))
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response.Responser(200, "Update Success"))
}

func (p *word) Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response.Responser(500, "Input fail"))
	}

	var w wordModel.Word

	w.ID = id
	fmt.Println(w)

	result := p.db.
		Delete(&w)

	if result.Error != nil {
		fmt.Println("Delete fail")
		ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response.Responser(500, "Input fail"))
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response.Responser(200, "Delete Success"))
}

func (p *word) RmqAdd(ctx iris.Context) {
	var body struct {
		UserId   int64  `json:"userId"`
		Sentence string `json:"sentence"`
	}
	err := ctx.ReadJSON(&body)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Word Rmq failure").DetailErr(err))
		return
	}

	body1, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	rmq.Push(p.rmq, body1)
	log.Printf(" [x] Sent %s\n", body1)

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(response.Responser(200, "Add Success"))
}

func (p *word) Test(ctx iris.Context) {
	ctx.JSON(response.Responser(200, "Success"))
}
