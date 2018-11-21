# STEP 1 build executable binary
FROM golang as builder

COPY . $GOPATH/src/ciklum-test/reader
WORKDIR $GOPATH/src/ciklum-test/reader

#get dependancies
RUN go get -d
#build the binary
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /go/bin/reader
# STEP 2 build a small image
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/reader /go/bin/reader
# Writer service adress
ENV gHOST writer
ENV gPORT 5001

CMD ["go/bin/reader", "-port=8080"]