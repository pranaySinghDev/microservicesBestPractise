build-user:
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o user-service/.builds/user-service user-service/cmd/main.go

build-product:
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o product-service/.builds/product-service product-service/cmd/main.go

build-image-user:
	docker build -t pranaysinghdev/user-service:latest ./user-service 

build-image-product:
	docker build -t pranaysinghdev/product-service:latest ./product-service