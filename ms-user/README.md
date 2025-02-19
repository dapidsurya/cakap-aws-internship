
# User Microservice

This is microservice for maintaining user data.

## Setup Using Docker Compose
Run this command to build the image and run container using docker compose

```bash
docker compose up -d --build
```
If you see the result like this, then the apps should be run
```bash
 ✔ ms-user                 Built    0.0s 
 ✔ Network user-network    Started  0.1s
 ✔ Container ms-user-db    Started  0.2s
 ✔ Container ms-user       Started  0.3s
```

Make sure the apps run on docker correctly by running
```bash
docker compose ps
```

In the initial run, the ms-user wouldn't start because the database is not created. To create the database you can go to ms-user-db container by running this command
```bash
docker exec -it ms-user-db /bin/bash
```

If you successfully go inside the ms-user-db container, then connect to mysql server.
```bash
mysql -u root -p
```

Then create the database as well as the table
```sql
CREATE DATABASE aws_academy_user;

USE aws_academy_user;

CREATE TABLE `user` (
  `id_user` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `fullname` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE(email),
  UNIQUE(username)
);
```

Then make sure if ms-user run well by running this command
```bash
docker compose ps
```

## Register New User
Congratulations! Now your application is running. Now if you access the API user list, it will show empty datas.
```bash
curl --location 'http://localhost:8080/user/list'
```
To register new user, you need to call this API and please adjust the payload data accordingly.
```bash
curl --location 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "budi",
    "fullname": "Budi Anduk",
    "email": "budi@email.com"
}'
```