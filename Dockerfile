# Build image (Golang)
FROM golang:1.12-alpine3.10 AS build-stage
ENV GO111MODULE on
ENV CGO_ENABLED 0

RUN apk add --no-cache gcc git make

WORKDIR /src
ADD . .

RUN go mod download
RUN go build -o filestat

# Final Docker image
FROM alpine:3.10 AS final-stage
LABEL MAINTAINER "fiorm <fiorm.github@gmail.com."

RUN apk add --no-cache ca-certificates

# Create user filestat
RUN addgroup -S filestat && adduser -S filestat -G filestat
USER filestat

WORKDIR ${HOME}/app
COPY --from=build-stage /src/filestat .

EXPOSE 2801

ENTRYPOINT ["./filestat"]
