FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD .builds/user-service /user-service 
ENTRYPOINT ["./user-service"]