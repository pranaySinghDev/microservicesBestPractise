# SETUP

Envs Required
```
user_port=5000
product_port=5001
db_url="mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
product_url=http://localhost:5001
```

Build
```
on root of the repo run:
make build-user
make build-product
```
Build Images
```
make build-image-user
make build-image-product
```

Docker Images

```
User Image:pranaysinghdev/user-service:latest 
Product Image: pranaysinghdev/product-service:latest
```

Configure VScode for debugging
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "User Service",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/user-service/cmd/main.go",
            "envFile": "${workspaceFolder}/compose.env"
        },
        {
            "name": "Product Service",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/product-service/cmd/main.go",
            "envFile": "${workspaceFolder}/compose.env"
        }
    ]
}
```