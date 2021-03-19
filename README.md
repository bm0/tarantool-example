```shell
docker build -t tarantool-example .
```

```shell
docker run --rm --name tarantool-example -d -p 3301:3301 tarantool-example
```

```shell
go run main.go
```
