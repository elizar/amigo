FROM golang:1.12.5
COPY . /amigo
WORKDIR /amigo
RUN CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"' amigo.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /amigo
COPY --from=0 /amigo /amigo
CMD ["./amigo"]

