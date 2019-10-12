FROM golang:1.13 AS builder
WORKDIR /go/src/github.com/erroneousboat/resume
COPY . .
RUN go get -d -v ./...
RUN make build
RUN ls ./bin/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/erroneousboat/resume/bin/resume /usr/bin/resume
COPY ./resume.tmpl .

ENV PORT 80
ENV PATH "/"
ENV USER ""
ENV PASS ""

EXPOSE $PORT

ENTRYPOINT /usr/bin/resume -user=$USER -pass=$PASS -port=$PORT -path=$PATH
