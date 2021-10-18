FROM golang:1.6-alpine
RUN mkdir /desafio_stone/main/main.go

ADD . /desafio_stone/main/main.go
WORKDIR /desafio_stone/main/main.go
RUN CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o main .

FROM alpine
EXPOSE 8080
CMD ["/desafio_stone/main/main.go"]

COPY --from=0 /src/main /src
