FROM alpine:3.12.0

COPY cms .
COPY build/*.dart.js dist/js/
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8107

ENTRYPOINT [ "./cms" ]
