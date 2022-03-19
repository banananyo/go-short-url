# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY app/go.mod ./
COPY app/go.sum ./
RUN go mod download

COPY app ./
# COPY *.go ./
# RUN mkdir handler
# COPY handler/*.go ./handler
# RUN mkdir shortener
# COPY shortener/*.go ./shortener
# RUN mkdir store
# COPY store/*.go ./store

RUN go build -o /docker-short-url

RUN export GIN_MODE=release

EXPOSE 9808

CMD [ "/docker-short-url" ]