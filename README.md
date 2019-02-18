# golang-postgresql

Install docker for postgresql

```docker run --name db -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:alpine```

Install postgresql admin
```docker run -p 8900:80 \
-e “PGADMIN_DEFAULT_EMAIL=email_cua_ban@gmail.com” \
-e “PGADMIN_DEFAULT_PASSWORD=password_cua_ban” \
-d dpage/pgadmin4
```
