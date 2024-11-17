
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
