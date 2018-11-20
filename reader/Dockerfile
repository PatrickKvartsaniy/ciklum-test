FROM golang as builder

RUN apt-get update && apt-get install git

RUN git clone https://github.com/PatrickKvartsaniy/ciklum-test

WORKDIR /go/src/reader/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/reader/reader /app/reader
EXPOSE 8888
ENTRYPOINT ./reader