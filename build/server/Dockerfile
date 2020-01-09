FROM golang:1.13.5

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -color=true -log-prefix=false -graceful-kill=true -exclude-dir=.idea -exclude-dir=app -include=.env -build="go build -tags dev ./cmd/websrv" -command=./websrv
