---
title: "Installation"
---

# Using Docker

Not implemented yet

```sh
docker-compose up -d
```

# Setup Manually

## 0. Pre requisites

| Lang             | Version |
| :--------------- | :------ |
| Nodejs           | 16.0    |
| Golang           | 1.18    |
| Sqlite3 or Mysql | -       |
| Caddy            | 2.4.*   |

## 1. Clone this repository

```
git clone https://github.com/rakutentech/code-coverage-dashboard.git
```

## 2. Set up backend server


### Copy .env file from example then <br>

Update your database username/password in case of database driver.

```sh
cp server/.env.local cp server/.env
```

### Migrate tables

```sh
cd ./server
go run cmd/db_migrate/main.go
```

### Start API Server

```sh
cd ./server
go run main.go <port>
```

## 3. Start Static file server

You can choose any static file server of your choice.
The port is reverse proxied through the API server.

```sh
caddy file-server --root ./server/assets --listen localhost:3008 --browse
```

## 4. Start Client server

```sh
cp client/.env.local cp client/.env
pnpm install
pnpm build
pnpm start
```