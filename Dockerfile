FROM golang:1.12 as BUILD_STAGE
RUN mkdir -p /src
WORKDIR /src
ENV CGO_ENABLED=0
ENV GOOS=linux
ADD . .
RUN go build -a -installsuffix cgo -o /build/main

FROM alpine:latest
RUN mkdir /app && apk add --update ca-certificates
COPY --from=BUILD_STAGE /build/main /app
WORKDIR /app
CMD ["./main"]