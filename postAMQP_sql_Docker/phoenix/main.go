package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"

	"github.com/streadway/amqp"
)

//User --> Genera schema de la tabla User
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

func send(b []byte) {
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

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(b),
		})
	log.Printf(" [x] Sent %s", b)
	failOnError(err, "Failed to publish a message")
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

//HashPassword devuelve el Hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash valida el password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		email := r.FormValue("email")
		password, _ := HashPassword(r.FormValue("password"))
		phone := r.FormValue("phone")
		country := r.FormValue("country")
		city := r.FormValue("city")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Correo = %s\n", email)
		fmt.Fprintf(w, "Clave = %s\n", password)
		fmt.Fprintf(w, "Teléfono = %s\n", phone)
		fmt.Fprintf(w, "Pais = %s\n", country)
		fmt.Fprintf(w, "Ciudad = %s\n", city)
		fmt.Fprintf(w, "Dirección = %s\n", address)

		fmt.Println("Envia a AMQP")

		user := User{Name: name,
			Email:    email,
			Password: password,
			Phone:    phone,
			Country:  country,
			City:     city,
			Address:  address}

		b, _ := json.Marshal(&user)
		send(b)

		fmt.Fprintf(w, "New User Successfully Created")

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	initialMigration()
	http.HandleFunc("/user", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
