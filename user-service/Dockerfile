FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY . ../.builds/user-service  
ENTRYPOINT ["./user-service"]