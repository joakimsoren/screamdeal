FROM golang:1.13

RUN mkdir /workspace
WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app

FROM scratch

COPY --from=0 /app /app

ENTRYPOINT ["/app"]
