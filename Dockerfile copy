# docker build . -t ogn
FROM golang:1.15.5

RUN go get -u github.com/beego/bee

WORKDIR /OGN
COPY . .

ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor
# ENV GOFLAGS=-mod=readonly
# RUN go build
EXPOSE 8080
CMD ["bee", "run"]
# WORKDIR /OGN
# RUN ["./ogn.exe"]
