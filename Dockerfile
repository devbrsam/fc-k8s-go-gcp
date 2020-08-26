FROM golang:alpine AS build

RUN mkdir /app
WORKDIR /app
COPY src/ .

#RUN go build -o main .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main .

FROM scratch
COPY --from=build app/main .

RUN mkdir /src
COPY --from=build app/template.html ./src

CMD ["/main"]