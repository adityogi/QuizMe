FROM golang:alpine as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o quizMe

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT [ "/app/quizMe" ]
