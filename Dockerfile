FROM golang:1.13.6
RUN mkdir /forum
ADD . /forum
WORKDIR /forum
# RUN go get github.com/mattn/go-sqlite3 v2.0.3+incompatible
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/satori/uuid
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/danonika/Forum
RUN go build main.go
# EXPOSE 8181
ENTRYPOINT ["/forum/main"]
