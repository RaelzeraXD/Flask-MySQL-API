FROM golang:1.22-alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum main.go ./

RUN go mod tidy
RUN go build main.go

EXPOSE 8080

FROM scratch

COPY --from=build /usr/src/app/main ./main

CMD ["./main"]