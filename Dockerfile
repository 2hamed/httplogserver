FROM golang as build

RUN mkdir /app

ADD main.go /app

WORKDIR /app

ENV CGO_ENABLED 0

RUN go build -o logserver main.go 

FROM alpine

COPY --from=build /app/logserver /bin/logserver

EXPOSE 8000

CMD ["logserver"]
