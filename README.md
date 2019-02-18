# golang-postgresql

Install docker for postgresql

```docker run --name db -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:alpine```

Install postgresql admin

```
docker run --name some-pgadmin4 \
           --link db \
           -p 5050:5050 \
           -d fenglc/pgadmin4
```
Create server 
```
Host name/address: docker ip
Port: 5432
Username: postgres
Password: 123
```
