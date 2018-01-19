# RUN INFLUXDB + GRAFANA
```
docker-compose up -d
```

# RUN GO TO POPULATE INFLUX DB
```
go run main.go seed.go
```

# ACCESS GRAFANA

>  `http://localhost:3001/`

```
USER: admin
PASS: admin
```
