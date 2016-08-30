# Golan 1.7
FROM golang:1.7

ENV APPDIR $GOPATH/src/github.com/elizar/amigo
CMD mkdir -p $APPDIR
COPY . $APPDIR

EXPOSE 8080

WORKDIR ${APPDIR}
CMD go build amigo.go
ENTRYPOINT ./amigo
