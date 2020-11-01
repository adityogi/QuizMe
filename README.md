# QuizMe
Application for MCQ

# Golang REST API for serving Quiz question and maintaining Quiz
 ## Default port for this QuizMe iamge is 3000
 - /quiz - GET - returns all questions
 - /question/{id} - GET - returns the qeustion for particular id
 - /question/{id} - POST - add a question with this id
 - /question/{id} - DELETE - remove a question with this id
 - /question/{id} - PUT - update the question and choice for this id

# Build Docker image
 cd quiz-api
 docker build -t quizApi:latest . 
 docker run quizApi:latest -d -p 5050:3000

# REACT web-ui for rendering the Quiz
 ## Default port for the react UI is 3000
 
