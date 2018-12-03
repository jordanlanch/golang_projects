package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/streadway/amqp"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Country  string
	City     string
	Address  string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func receive() {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbit1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		bytes := []byte(d.Body)
		log.Printf("Received a message: %s", bytes)
		createUser(bytes)
		break
	}

}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "../db/test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "../db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func createUser(bytes []byte) {
	var userDecode User
	json.Unmarshal(bytes, &userDecode)

	fmt.Println("New User Endpoint Hit")
	res2B, _ := json.Marshal(&userDecode)
	fmt.Print("lo que debe ir a la db " + string(res2B))

	db, err := gorm.Open("sqlite3", "../db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	if db.NewRecord(userDecode) == true {
		fmt.Println("returns `true` as primary key is blank")
	}
	db.Create(&userDecode)
	if db.NewRecord(userDecode) == true {
		fmt.Println("return `false` after `user` created")
	}

}

func main() {
	initialMigration()
	receive()
	http.HandleFunc("/users", allUsers)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
