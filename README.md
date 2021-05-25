# microservicesBestPractise

## Product service

`Endpoints`

BaseURL 

```
curl --location --request GET 'http://localhost:5001'
```

Create Product 

```
curl --location --request POST 'http://localhost:5001/products' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Item3",
    "category": "general",
    "price": 17.00
}'
```

Get Product
```
curl --location --request GET 'http://localhost:5001/products/447a74a9-769a-4145-8b9d-f55d3d3f9d73'
```


Get All Products

```
// Fetches first 10 records if limit and index not given
curl --location --request GET 'http://localhost:5001/products'
// limit = number of records , index =  page on which to fetch
curl --location --request GET 'http://localhost:5001/products?limit=10&index=0'
```

## User service

`Endpoints`

BaseURL 

```
curl --location --request GET 'http://localhost:5000'
```

Create User 

```
curl --location --request POST 'http://localhost:5000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Smith",
    "age": 27,
    "product": "44d5aee7-e05e-4e15-897a-0c59df5091cb"
}'
```

Get User
```
curl --location --request GET 'http://localhost:5000/users/ebe18547-adf0-45e9-8bd7-643740fef241'
```


Get All Users
```
// Fetches first 10 records if limit and index not given
curl --location --request GET 'http://localhost:5000/users'
// limit = number of records , index =  page on which to fetch
curl --location --request GET 'http://localhost:5000/users?limit=10&index=0'
```
