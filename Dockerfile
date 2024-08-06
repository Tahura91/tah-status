FROM golang:1.22.2

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN go build -o http-status-checker

ENTRYPOINT ["./http-status-checker"]
