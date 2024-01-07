FROM golang:1.21.5 AS build

WORKDIR /go/src/serverless-blog

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src .

RUN CGO_ENABLED=0 go build -o /go/bin/serverless-blog ./cmd/main.go

FROM gcr.io/distroless/base-debian12:latest-amd64

COPY --from=build /go/bin/serverless-blog /
COPY --from=build /go/src/serverless-blog/config/config.yaml /

EXPOSE 8080

CMD ["/serverless-blog", "--config_path=config.yaml"]