FROM golang
RUN go get -u github.com/yanzay/tbot
RUN go get github.com/kyokomi/emoji
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/sirupsen/logrus