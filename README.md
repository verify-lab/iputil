# IP Address utils

![Tests](https://github.com/verify-lab/iputil/actions/workflows/tests.yml/badge.svg)

```sh
go get github.com/verify-lab/iputil
```

## IP util example

```go

ip := net.ParseIP("99.250.150.130")

fmt.Println(IsIPv6(ip))
fmt.Println(IsPrivateIP(ip))

```
