FROM golang

RUN mkdir /trustedtext

ADD . /trustedtext

WORKDIR /trustedtext

RUN go mod download

RUN go build . 

CMD ["./trustedtext"]