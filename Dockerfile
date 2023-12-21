FROM golang:1.21-alpine AS builder

ENV APP=go-grpc-server
ENV GO111MODULE=on

RUN apk update && apk add git make gdb
WORKDIR /go/src/github.com/lucassouzavieira/$APP/

# Handle dependencies
COPY go.mod go.sum ./
COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 go build -o /build/app -a ./cmd/app
RUN mv /build/app /app
RUN ls -la
RUN cp data/*.csv /
RUN ls -la /

# Final image
FROM alpine:latest AS compile

RUN apk --no-cache add ca-certificates
# Copy image binary and data for the created image
COPY --from=builder /app /app
COPY --from=builder /*.csv /data/ 

EXPOSE 9200

CMD ["./app"]