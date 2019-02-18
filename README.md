# golang-postgresql

Install docker for postgresql

```docker run --name db -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:alpine```

Install postgresql admin

```
docker run --name some-pgadmin4 \
           --link some-postgres:postgres \
           -p 5050:5050 \
           -d fenglc/pgadmin4
```
