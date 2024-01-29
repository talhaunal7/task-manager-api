FROM golang:1.20.3-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/tma

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/tma .
COPY .env /app

EXPOSE 8080

CMD ["./tma"]