
### Database Steps


`docker-compose up -d`

check if there
`docker ps`


### restart
```
docker-compose down
docker-compose up -d
```

### check db and roles
`docker exec -i postgres_container psql -U postgres -d attendance`
`-it`: Starts an interactive terminal session.
`psql`: The PostgreSQL CLI tool.
`-U postgres`: Logs in as the postgres user.
`-d attendance`: Connects to the attendance database.

check roles
`\du`

check tables
`\l`

having connection failures?
```
attendance=# \l
attendance=# \c attendance
You are now connected to database "attendance" as user "postgres".
attendance=# \du
attendance=# GRANT ALL PRIVILEGES ON DATABASE attendance TO postgres;
GRANT
```

---

├── cmd/
│   └── api/
│       └── main.go       # Entry point
├── internal/
│   ├── handler/
│   │   └── attendance.go # HTTP handlers
│   ├── model/
│   │   └── attendance.go # Data structures
│   ├── db/
│   │   └── postgres.go   # Database connection and queries
│   └── api/
│       └── response.go   # Response helpers
└── config/
    └── config.go         # Configuration handling

## Run basic queries

```
docker ps // check if docker running container

docker exec -it postgres_container psql -U postgres -d attendance


attendance#: 
SELECT * FROM attendance
ORDER BY created_at DESC
LIMIT 10;

```

```
SELECT created_at, address, entrance_status
FROM attendance
ORDER BY created_at DESC
LIMIT 10;
```
