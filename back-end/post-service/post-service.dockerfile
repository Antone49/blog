# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY postApp /app

CMD [ "/app/postApp" ]