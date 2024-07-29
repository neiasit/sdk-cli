package templates

const DockerfileTemplate = `FROM golang:{{ .GolangVersion }}
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o root .

CMD ["./main"]
`
