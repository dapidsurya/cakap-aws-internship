
# Product Microservice

This is microservice for maintaining product data.

## Setup Using Docker Compose
Run this command to build the image and run container using docker compose

```bash
docker compose up -d --build
```
If you see the result like this, then the apps should be run
```bash
 ✔ ms-product                 Built    0.0s 
 ✔ Network product-network    Started  0.1s
 ✔ Container ms-product-db    Started  0.2s
 ✔ Container ms-product       Started  0.3s
```

Make sure the apps run on docker correctly by running
```bash
docker compose ps
```

In the initial run, the ms-product wouldn't start because the database is not created. To create the database you can go to ms-product-db container by running this command
```bash
docker exec -it ms-product-db /bin/bash
```

If you successfully go inside the ms-product-db container, then connect to mysql server.
```bash
mysql -u root -p
```

Then create the database as well as the table
```sql
CREATE DATABASE aws_academy_product;

USE aws_academy_product;

CREATE TABLE `product` (
  `id_product` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `stock` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id_product`)
);
```

Then make sure if ms-product run well by running this command
```bash
docker compose ps
```

## Create New Product
Congratulations! Now your application is running. Now if you access the API, it will show welcome message.
```bash
curl --location 'http://localhost:8090/'
```
To create new product based on userID, you need to call this API and please adjust the payload data accordingly.
```bash
curl --location 'localhost:8090/product' \
--header 'Content-Type: application/json' \
--data '{
    "userId": 1,
    "name": "USB type C Cable 1A",
    "stock": 11
}'
```