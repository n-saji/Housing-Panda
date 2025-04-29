FROM golang:1.24

WORKDIR /root

COPY ./ /root/

RUN go mod download

RUN go build -o housing_panda_app ./main.go

EXPOSE 5050

CMD ["./housing_panda_app"]