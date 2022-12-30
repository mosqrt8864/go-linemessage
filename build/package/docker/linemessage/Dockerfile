FROM golang:1.18.0 AS go-builder
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor

WORKDIR /linemessage
COPY . .
RUN go build -o linemessage /linemessage/cmd/

FROM alpine:3.15
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add gcompat && \
    apk add libstdc++ && \
    apk add git
WORKDIR /linemessage
COPY --from=go-builder /linemessage/linemessage /linemessage/app
COPY --from=go-builder /linemessage/configs /linemessage/configs
EXPOSE 80
ENTRYPOINT ["/linemessage/app"]