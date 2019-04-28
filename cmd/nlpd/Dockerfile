FROM golang:1.12.4 as build

WORKDIR /nlp
COPY . .
ENV CGO_ENABLED=0
RUN go build ./cmd/nlpd

FROM debian:stretch-slim
COPY --from=build /nlp/nlpd /usr/local/bin

ENV NLPD_PORT 8080
CMD nlpd
