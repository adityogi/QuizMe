# QuizMe - Application for MCQ

## Go based REST API for serving Quiz question and maintaining Quiz
> Default port for this QuizMe image is 3000
 - /quiz - GET - returns all questions
 - /question/{id} - GET - returns the question for particular id
 - /question/{id} - POST - add a question with this id
 - /question/{id} - DELETE - remove a question with this id
 - /question/{id} - PUT - update the question and choice for this id

> Build and Run Docker image
```
cd quiz-api
docker build -t quiz-api:latest .
docker run quiz-api:latest -d -p 5050:3000
```
## REACT based web-ui for rendering the Quiz
> Default port for the react UI is 3000 to run locally,
  ``` 
  npm install prop-types
  npm install react-transition-group@1.x
  npm start
  ```
> Build and Run Docker image
```
cd quiz-api
docker build -t quiz-api:latest .
docker run quiz-api:latest -d -p 5050:3000`
```