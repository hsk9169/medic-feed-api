FROM --platform=linux/amd64 centos:7 as builder

ENV GO_VERSION=1.20
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

USER root

RUN yum -y install make wget

RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar xvzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /app 

COPY Config Config/
COPY Controllers Controllers/
COPY Routes Routes/
COPY Services Services/
COPY Validations Validations/
COPY Utils Utils/
COPY go.mod go.sum ./
COPY cmd cmd/
COPY Makefile ./
COPY .env ./

RUN make build

FROM alpine:3.18.2
USER root
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
    echo "Asia/Seoul" > /etc/timezone \
    apk del tzdata

WORKDIR /app
COPY --from=builder /app/storage /app/storage

EXPOSE 8080
CMD ["./storage"]
