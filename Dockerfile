############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
RUN adduser -S scratchuser

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies
RUN go get -d -v && \
    CGO_ENABLED=0 go build -o /go/bin/hello && \
    chmod +x /go/bin/hello

############################
# STEP 2 build a small image
############################
FROM scratch

EXPOSE 8000
COPY --from=0 /etc/passwd /etc/passwd
USER scratchuser

COPY --from=builder /go/bin/hello /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]
