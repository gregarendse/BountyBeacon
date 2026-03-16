# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder
WORKDIR /src

ARG TARGETOS
ARG TARGETARCH

# Cache module downloads separately from source
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY main.go ./
COPY cli/ ./cli/
COPY lib/ ./lib/

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/bountybeacon .

FROM alpine:3.21 AS certs
RUN apk add --no-cache ca-certificates

FROM scratch
COPY --from=builder /out/bountybeacon /bountybeacon
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/bountybeacon"]
CMD ["claim"]

