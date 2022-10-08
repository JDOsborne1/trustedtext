FROM golang

RUN mkdir /trustedtext
ADD . /trustedtext

WORKDIR /trustedtext/cmd/webserver

RUN go build .

CMD ["./trustedtext-webserver"]