FROM busybox

COPY ./bin/resume resume
COPY ./resume.tmpl /

ENV PORT 80
ENV USER ""
ENV PASS ""

EXPOSE $PORT

ENTRYPOINT /resume -user=$USER -pass=$PASS -port=$PORT
