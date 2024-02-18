# Prepare the build environment and build the image
FROM golang:1.22.0-alpine as builder
RUN mkdir -p /build
ADD * /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ripe-serivce cmd/main.go

# Prepare the image for the final run
FROM alpine:3.15.4
COPY --from=builder /build/ripe-serivce .

# Run the image with entrypoint
ENTRYPOINT [ "./ripe-serivce" ]