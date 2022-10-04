# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY imageApp /app

CMD [ "/app/imageApp" ]