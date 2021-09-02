FROM golang:1.6-alpine
RUN mkdir /src
ADD . /src/
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o main .

FROM alpine
EXPOSE 8080
CMD ["/src"]

COPY --from=0 /src/main /src
