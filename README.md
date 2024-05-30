# BTAW

**Breaking Trend Analysis Wizard**

**help**
```
make
```

**api**

```
# retreive history data
http://<host:port>/klines/<symbol>/<tf>/<t0_ms>/<duration>

http://localhost:4000/klines/BTC-USDT/1h/1715555676044/1d
http://localhost:4000/klines/BTC-USDT/15m/1715555676044/1d
	

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