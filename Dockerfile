FROM golang:latest
WORKDIR /fstest_repro
COPY fstest_repro.go .

CMD go run fstest_repro.go