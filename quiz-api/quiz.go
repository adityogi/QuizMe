package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Contact struct (Model)
type QuizItem struct {
	Id       int       `json:"id"`
	Question string    `json:"question"`
	Choice   [4]string `json:"choice"`
}

// Init contacts var as a slice Contact struct
var quiz []QuizItem

// Get all questions
func getQuiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}

// Get single question
func getQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		return
	}
	// Looping through contacts and find one with the id from the params
	for _, item := range quiz {
		if item.Id == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&QuizItem{})
}

// Add new contact
func addQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var quizItem QuizItem
	_ = json.NewDecoder(r.Body).Decode(&quizItem)
	quiz = append(quiz, quizItem)
	json.NewEncoder(w).Encode(quizItem)
}

// Delete contact
func deleteQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		return
	}
	for idx, item := range quiz {
		if item.Id == i {
			quiz = append(quiz[:idx], quiz[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(quiz)
}

// Update contact
func updateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		return
	}
	for idx, item := range quiz {
		if item.Id == i {
			quiz = append(quiz[:idx], quiz[idx+1:]...)
			var quizItem QuizItem
			_ = json.NewDecoder(r.Body).Decode(&quizItem)
			quizItem.Question = params["question"]
			quiz = append(quiz, quizItem)
			json.NewEncoder(w).Encode(quizItem)
			return
		}
	}
}

func main() {
	fmt.Println("quiz.go")
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	choice := [4]string{"Yes", "No"}
	quiz = append(quiz, QuizItem{Id: 1, Question: "Is Python an interpreted language?", Choice: choice})

	// Route handles & endpoints
	r.HandleFunc("/quiz", getQuiz).Methods("GET")
	r.HandleFunc("/question/{id}", getQuiz).Methods("GET")
	r.HandleFunc("/question", addQuestion).Methods("POST")
	r.HandleFunc("/question/{id}", updateQuestion).Methods("PUT")
	r.HandleFunc("/question/{id}", deleteQuestion).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}
