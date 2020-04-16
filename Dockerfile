FROM golang:1.13.5

WORKDIR /opt/alug

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /opt/alug/bin/alug ./cmd/main.go
