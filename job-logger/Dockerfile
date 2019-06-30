FROM golang:latest AS builder
RUN mkdir /job-logger
ADD main.go /job-logger/
WORKDIR /job-logger
COPY go.mod /job-logger/
COPY go.sum /job-logger/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o job-logger .

FROM scratch
COPY --from=builder /job-logger/job-logger .
ENTRYPOINT ["./job-logger"]