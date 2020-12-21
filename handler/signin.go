package handler

import (
	"github.com/vitoraalmeida/silkroad/usecase/customer"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type SignIn struct {
	l  *log.Logger
	cs *customer.Service
}

func NewSignIn(l *log.Logger, cs *customer.Service) *SignIn {
	return &SignIn{
		l,
		cs,
	}
}

func (s *SignIn) SignIn(w http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST signin")
	errorMessage := "Error signing in"
	w.Header().Set("Content-type", "text/html")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	email := r.PostForm["email"][0]
	password := r.PostForm["password"][0]
	inputErr := validate(email, password)
	if inputErr != "" {
		s.l.Println(errorMessage)
		s.l.Println(inputErr)
		return
	}

	customer, err := s.cs.GetCustomerByEmail(email)
	if err != nil {
		s.l.Println(errorMessage)
		s.l.Println(inputErr)
		return
	}
	if customer == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(customer.Password),
		[]byte(password),
	)
	if err != nil {
		s.l.Println(errorMessage)
		s.l.Println(inputErr)
		return
	}
	s.l.Println("Logged in")
	http.Redirect(w, r, "/", http.StatusSeeOther)

	//token, err := auth.CreateToken(user.ID)
	//if err != nil {
	//	http.Error(w, errorMessage, http.StatusInternalServerError)
	//	return
	//}

}

func validate(email, password string) string {
	if email == "" {
		return "Account name must not be empty"
	}
	if len(password) < 6 {
		return "Password must be at least 6 digits"
	}
	return ""
}
