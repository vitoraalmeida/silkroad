package handler

import (
	_ "github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/usecase/customer"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type Customers struct {
	l  *log.Logger
	cs *customer.Service
}

func NewCustomers(l *log.Logger, cs *customer.Service) *Customers {
	return &Customers{
		l,
		cs,
	}
}

func (c *Customers) ListCustomers(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET customers")
	errorMessage := "Error reading accounts"
	data, err := c.cs.ListCustomers()
	if err != nil {
		c.l.Println(errorMessage)
		panic(err)
	}
	c.l.Println(data)
}

func (c *Customers) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST customers")
	errorMessage := "Error creating account"
	w.Header().Set("Content-type", "text/html")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	name := r.PostForm["name"][0]
	email := r.PostForm["email"][0]
	cpf := r.PostForm["cpf"][0]
	password := r.PostForm["password"][0]

	inputErr := c.Validate(name, email, cpf, password)
	if inputErr != "" {
		c.l.Println(errorMessage)
		c.l.Println(inputErr)
		return
	}

	secret := []byte(password)
	hashSecret, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
	if err != nil {
		c.l.Println(errorMessage)
		c.l.Println(err)
		return
	}

	password = string(hashSecret)
	id, err := c.cs.CreateCustomer(name, email, cpf, password)
	if err != nil {
		c.l.Println(err)
		return
	}
	c.l.Println("Customer created id: ", id)
	w.WriteHeader(http.StatusCreated)
}

func (c *Customers) Validate(name, email, cpf, password string) string {
	if name == "" {
		return "Account name must not be empty"
	}
	if email == "" {
		return "Account name must not be empty"
	}
	if cpf == "" || len(cpf) < 11 {
		return "CPF must be at least 11 digits"
	}
	_, err := strconv.Atoi(cpf)
	if err != nil {
		return "CPF must only contain numbers"
	}
	if len(password) < 6 {
		return "Password must be at least 6 digits"
	}
	return ""
}
