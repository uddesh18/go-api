FROM golang:1.16 AS build
WORKDIR /test
COPY . /test
RUN go build /test
EXPOSE 8000
ENTRYPOINT [ "./go-api" ]