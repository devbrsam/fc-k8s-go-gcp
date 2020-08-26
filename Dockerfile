FROM golang:alpine AS build

RUN mkdir /app
WORKDIR /app
COPY src/ .

#RUN go build -o main .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

FROM scratch
COPY --from=build app/main .
CMD ["/main"]