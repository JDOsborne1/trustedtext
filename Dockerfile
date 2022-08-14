FROM golang

RUN mkdir /trustedtext
ADD go.mod /trustedtext
ADD go.sum /trustedtext


WORKDIR /trustedtext

RUN go mod download
ADD . /trustedtext
RUN go build . 

WORKDIR /trustedtext/trustedtext-webserver

RUN go build .

CMD ["./trustedtext-webserver"]