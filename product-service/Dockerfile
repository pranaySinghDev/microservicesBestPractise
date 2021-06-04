FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD .builds/product-service /product-service 
ENTRYPOINT ["./product-service"]