FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app .

FROM alpine:3.14
RUN apk update && apk add ncurses
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE 8080
CMD ["/go/bin/app"]