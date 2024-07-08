FROM golang:1.20

WORKDIR /training-ut-util-go

COPY . .

RUN go mod tidy

WORKDIR /training-ut-util-go/scripts

