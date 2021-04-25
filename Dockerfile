FROM node:15.14.0 as client-builder
WORKDIR /build
COPY ./client/ .
RUN npm ci
RUN npm run build

FROM golang:alpine AS golang-builder
WORKDIR /build
COPY ./back/ .
RUN go mod download
RUN go build ./cmd/temsys
WORKDIR /dist
RUN cp /build/temsys .

FROM alpine
EXPOSE 3000
COPY --from=golang-builder /build/temsys /
COPY ./back/dev.example.json /dev.json
COPY ./back/ascii-art.ans /
COPY --from=client-builder /build/dist/ /static/
RUN mkdir /data
ENV CONFIG_TLSCRT=/data/server.crt
ENV CONFIG_TLSKEY=/data/server.key
ENV CONFIG_TLSENABLED=true

ENTRYPOINT ["/temsys"]
