
### Database Steps

`brew services start postgresql@14`

`docker-compose up -d`

check if there
`docker ps`


### restart
```
docker-compose down
docker-compose up -d
```

### check db and roles
`docker exec -it postgres_container psql -U postgres -d attendance`

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
