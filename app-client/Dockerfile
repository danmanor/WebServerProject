FROM docker.io/library/golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

ENV GO111MODULE=on

COPY  . ./

RUN go mod download
# Download necessary Go modules

CMD ["go", "run", "main.go"]