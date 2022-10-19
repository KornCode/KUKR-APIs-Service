FROM golang:alpine as builder

LABEL maintainer="Korn Chotpitakkul <korn.chot@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/.env .

EXPOSE 3000

USER notroot:nonroot

CMD ["./server"]
