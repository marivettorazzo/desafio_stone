FROM golang:onbuild
WORKDIR /src
COPY . .
RUN yarn install --production
CMD ["golang", "/src/main.go"]
EXPOSE 8080

