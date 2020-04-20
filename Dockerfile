FROM golang:1.13.5

WORKDIR /opt/alug

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build
