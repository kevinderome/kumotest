FROM golang:1.14 AS build_image
ARG CGO_ENABLED=0 
ARG GOOS=linux 
WORKDIR /build
COPY src/ .
RUN go get -d -v .
RUN go build -a -installsuffix cgo -o app .
FROM alpine:latest
COPY --from=build_image /build/app /bin/app
CMD ["app"]