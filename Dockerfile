# docker build . -t ogn
FROM golang:1.19

WORKDIR /OGN
COPY . .

ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor
# ENV GOFLAGS=-mod=readonly
RUN go mod download
RUN go mod verify
RUN go build -o ogn
EXPOSE 9000
WORKDIR /OGN
CMD ["./ogn"]
# docker build -t ogn .
#  docker tag ogn:latest oganator/dockerhub:push  -- 2
#  docker push oganator/dockerhub:push