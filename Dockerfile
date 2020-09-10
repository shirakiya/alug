FROM golang:1.13.5

WORKDIR /opt/alug

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
  sh -s -- -b $(go env GOPATH)/bin v1.31.0

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build