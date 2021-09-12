FROM golang:1.16.7-alpine

WORKDIR /go/src/github.com/Ras96/clean-architecture-sample
COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go install github.com/cosmtrek/air@v1.27.3

# usermodなどで手元のUIDが変わっている場合は.envに記述する
RUN chown -R ${UID:-1000}:${GID:-1000} ./

CMD ["air", "-c", ".air.toml"]
