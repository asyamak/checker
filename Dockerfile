FROM golang:1.20-alpine as build
LABEL maintainer = "asyamak"
WORKDIR /app
COPY . .
RUN go build -o api cmd/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=build /app .
EXPOSE 9090

CMD ["/app/api"]