FROM  golang:1.24.2-alpine3.21 as builder

RUN apk --no-cache add bash fish ca-certificates curl build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN echo $(apk --print-arch)

RUN CGO_ENABLED=0 go build -v -o hornet ./cmd/

EXPOSE 8080

RUN chmod +x /app/hornet

CMD ["./hornet"]