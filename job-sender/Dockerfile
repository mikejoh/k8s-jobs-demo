FROM golang:latest AS builder
RUN mkdir /job-sender
ADD main.go /job-sender/
WORKDIR /job-sender
COPY go.mod /job-sender/
COPY go.sum /job-sender/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o job-sender .

FROM scratch
COPY --from=builder /job-sender/job-sender .
ENTRYPOINT ["./job-sender"]