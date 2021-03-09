FROM golang:1.15 as builder

RUN mkdir /opt/app/
WORKDIR /opt/app/

ENV GOOS=js
ENV GOARCH=wasm

ADD golang/ssh-keygen.go /opt/app/
RUN go get golang.org/x/crypto/ssh
RUN go build -o ssh-keygen.wasm ssh-keygen.go

FROM nginx 

RUN sed -i -e "5a application/wasm wasm;" /etc/nginx/mime.types

COPY --from=builder /opt/app/ssh-keygen.wasm /usr/share/nginx/html/
ADD docs/* /usr/share/nginx/html/
