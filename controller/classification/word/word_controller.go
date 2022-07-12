package word

import (
	"fmt"
	"log"
	wordModel "myapp/model/classification/word"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type word struct {
	db  *gorm.DB
	rmq *amqp.Connection
}

type response struct {
	Msg  string
	Code int
}

func New(db *gorm.DB, rmq *amqp.Connection) *word {
	return &word{db, rmq}
}

func (p *word) Create(ctx iris.Context) {
	var w wordModel.Word

	err := ctx.ReadJSON(&w)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Word creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
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

func (p *word) List(ctx iris.Context) {
	page, err := strconv.Atoi(ctx.Params().Get("page"))
	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response{Msg: "Input fail", Code: 500})
	}

	var w []wordModel.Word
	result := p.db.Limit(10).Offset(page - 1).Find(&w)
	if result.Error != nil {
		fmt.Println("List fail")
		// ctx.StatusCode(iris.StatusBadGateway)
		ctx.JSON(response{Msg: "List fail", Code: 500})
	}
	ctx.StatusCode(iris.StatusAccepted)
	fmt.Println(w)

	ctx.JSON(w)
}

func (p *word) Read(ctx iris.Context) {

}

func (p *word) Update(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response{Msg: "Input fail", Code: 500})
	}

	var w wordModel.Word

	if ctx.ReadJSON(&w) != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Word Update failure").DetailErr(err))
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

func (p *word) Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Input fail")
		ctx.JSON(response{Msg: "Input fail", Code: 500})
	}

	var w wordModel.Word

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

func (p *word) RmqAdd(ctx iris.Context) {
	ch, err := p.rmq.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	body := "Hello World!"
	//放進queue
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
	}
	log.Printf(" [x] Sent %s\n", body)
}
