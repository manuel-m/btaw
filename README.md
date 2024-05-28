# BTAW

**Breaking Trend Analysis Wizard**



**run**
```
go mod tidy
go run ./cmd/api
```

**api**

```
# retreive history data
http://<host:port>/klines/<symbol>/<interval>

# retreive history data
http://<host:port>/health

```



**DB setup**

```
sudo -u postgres psql
  CREATE DATABASE btaw_dev;
  \l
  CREATE ROLE btaw_dev WITH LOGIN PASSWORD 'changeme';
  GRANT ALL PRIVILEGES ON DATABASE btaw_dev TO btaw_dev;
  \du

# [!] patch /var/lib/pgsql/data/pg_hba.conf

psql -U btaw_dev -d btaw_dev
```