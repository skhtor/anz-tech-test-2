############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
RUN addgroup -S scratchuser && \
    adduser -S -G scratchuser scratchuser

WORKDIR /go/src/mypackage/myapp/
COPY . .

ENV BUILD_VERSION=${BUILD_VERSION}
ENV COMMIT_SHA=${COMMIT_SHA}

RUN go get -d -v && \
    CGO_ENABLED=0 go build -o /go/bin/hello && \
    chmod +x /go/bin/hello

############################
# STEP 2 build a small image
############################
FROM scratch

ARG BUILD_VERSION
ARG COMMIT_SHA

ENV BUILD_VERSION=${BUILD_VERSION}
ENV COMMIT_SHA=${COMMIT_SHA}

EXPOSE 8000

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

USER scratchuser:scratchuser

COPY --from=builder /go/bin/hello /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]
