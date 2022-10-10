FROM alpine:latest

COPY out/httpWaitServer ./

CMD [ "/httpWaitServer" ]

