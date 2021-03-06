FROM golang:1.17

WORKDIR /getirAssignment
COPY ./ /getirAssignment

RUN go get ./...

EXPOSE 8000

CMD [ "go", "run", "main.go" ]