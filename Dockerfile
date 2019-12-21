#final stage
FROM alpine:latest
WORKDIR /opt/static-web
COPY ./staticweb /opt/static-web
COPY ./app.config /opt/static-web
COPY ./static /opt/static-web/static
ENTRYPOINT ./staticweb
EXPOSE 8800
