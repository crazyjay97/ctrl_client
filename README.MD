# Simple Serial-Port Tool Build By Fyne

### Command

##### Run

```shell
go run .\main.go
```

##### Build Without Windows Cmd's Window

```shell
go build  -ldflags -H=windowsgui .\main.go
```

```shell
go-bindata -o=".\assets\assets.go" -pkg=assets "./static/..."
```
