# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY messageApp /app

CMD [ "/app/messageApp" ]