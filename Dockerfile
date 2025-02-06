FROM golang:1.22-alpine AS builder

RUN apk add build-base

WORKDIR /quiz

COPY ./go.* ./
RUN go mod download

COPY ./ ./

RUN go test ./...

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal -g ./internal/rest/api.go -o ./internal/rest -ot "json" 

RUN go build -v -o api ./cmd/api

FROM alpine:latest AS runtime

COPY --from=builder /quiz/api /quiz/api
COPY --from=builder /quiz/cmd/api/config.yml /quiz/config.yml

CMD ["/quiz/api"]