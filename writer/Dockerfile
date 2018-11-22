# STEP 1 build executable binary
FROM golang as builder

COPY . $GOPATH/src/ciklum-test/writer
WORKDIR $GOPATH/src/ciklum-test/writer

#get dependancies
RUN go get -d
#build the binary
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /go/bin/writer

# STEP 2 build a small image
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/writer /go/bin/writer

# Postgres config
ENV DBHost db
ENV DBUser patrick
ENV DBPassword erasmusmundus
ENV DBName ciklum

CMD ["go/bin/writer", "-makemigrations","-port=5001"]