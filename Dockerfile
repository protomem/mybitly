FROM golang:alpine AS builder
WORKDIR /github.com/protomem/mybitly
RUN apk add --no-cache make
COPY . .
RUN make build

FROM alpine
WORKDIR /root
COPY --from=builder /github.com/protomem/mybitly/build/ /root/
COPY --from=builder /github.com/protomem/mybitly/configs/ /root/configs/
CMD [ "/root/api" ]