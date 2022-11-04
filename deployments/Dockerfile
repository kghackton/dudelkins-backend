FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /anomalius/cmd

COPY . /anomalius

RUN go build -o main .

FROM alpine:3.16

COPY --from=builder /anomalius/cmd/main /anomalius/cmd/main
COPY --from=builder /anomalius/assets/defectIdsDuration.json /anomalius/cmd/defectIdsDuration.json

ENTRYPOINT ["/anomalius/cmd/main"]
