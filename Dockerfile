FROM golang

RUN mkdir /gocode

ENV GOPATH="/gocode"

ENV PATH="$PATH:$GOPATH/bin"

RUN apt-get update

RUN apt-get install sqlite3 libsqlite3-dev

RUN go get github.com/mattn/go-sqlite3

RUN go get github.com/jmoiron/sqlx

RUN go get github.com/gorilla/mux

RUN go get github.com/gorilla/sessions

RUN apt-get install python

RUN apt-get install -y python-pip

RUN pip install -U pip setuptools

RUN easy_install supervisor

ENV APPLOGPATH="/var/log/"

WORKDIR $GOPATH
