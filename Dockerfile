FROM golang:alpine as builder

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	ca-certificates

COPY . /go/src/github.com/erroneousboat/resume

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/erroneousboat/resume \
	&& make build \
	&& mv ./bin/resume /usr/bin/resume \
	&& apk del .build-deps \
	&& rm -rf /go

FROM alpine:latest

COPY --from=builder /usr/bin/resume /usr/bin/resume
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY ./resume.tmpl /

ENV PORT 80
ENV USER ""
ENV PASS ""

EXPOSE $PORT

ENTRYPOINT resume -user=$USER -pass=$PASS -port=$PORT
