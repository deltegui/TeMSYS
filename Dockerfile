FROM golang:alpine AS golang-builder
WORKDIR /build
COPY ./back/ .
RUN go mod download
RUN go build ./cmd/temsys
WORKDIR /dist
RUN cp /build/temsys .
COPY ./back/dev.json .
COPY ./server.crt .
COPY ./server.key .
COPY ./back/ascii-art.ans .
RUN mkdir /build/public

FROM node:15.6.0-alpine3.10 as client-builder
WORKDIR /build
COPY ./client/ .
RUN npm i -g @angular/cli
RUN npm i
RUN npm run build

FROM alpine
EXPOSE 3000
COPY --from=golang-builder /dist/* /
COPY --from=client-builder /build/public/ /static/
RUN mkdir /data

ENTRYPOINT ["/temsys"]
