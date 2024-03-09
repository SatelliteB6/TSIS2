# League of Legends Project

Welcome to the League of Legends Project, a simple Go project inspired by the world of League of Legends. 

## League REST API

```
GET /champions
POST /champions
GET /champions/:id
PUT /champions/:id
DELETE /champions/:id
```

## DB Structure

```
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

// Use DBML to define your database structure
// Docs: https://dbml.org/docs

Table champions {
  id serial [primary key]
  name varchar(255)
  class varchar(255)
  pice int
}
```
