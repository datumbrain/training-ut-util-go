FROM golang:1.20

WORKDIR /training-ut-util-go

COPY . .

RUN go mod tidy

WORKDIR /training-ut-util-go/scripts

RUN chmod +x /training-ut-util-go/scripts/docker_entrypoint.sh

CMD ["/training-ut-util-go/scripts/docker_entrypoint.sh"]
