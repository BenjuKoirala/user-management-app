# User Management Application
This application is a simple front-end and back-end application to manage users. The management includes getting users information,
create users (name and email), delete and update users.

# Features
* Get users info
* Create user
* Update user
* Delete user

# Prerequisites
* Go 1.16+
* NodeJS
* PostgresSQL 

# Clone repository
```
git clone https://github.com/BenjuKoirala/user-management-app
```

# Run frontend

```
cd user-management-frontend
run `npm install --force`
run `ng serve`
open http://localhost:4200/
```

## Make necessary updates on the config files
1. Update config.json and db_config.json with necessary credentials before running the application

## Create database table 
```
create table users (id serial primary key, name varchar (50) not null unique, email varchar(255) not null unique);
```

# Run backend

```
cd user-management-backend
go build user-management-backend
```

### Then start using the frontend -
 * when you first  open http://localhost:4200/ it will hit get users, and then you can start creating, updating or deleting users
 * Just to test backend you could also make use of tools like Postman



