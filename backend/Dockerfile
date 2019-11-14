FROM golang:1.13

RUN mkdir /workspace
WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app

ENTRYPOINT ["/app"]

