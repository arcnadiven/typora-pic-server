# typora-pic-server
A picture server for Typora running in local

run this server and copy the command to Typora Image Settings use custom command

```shell
go mod tidy && go run main.go

open another terminal

/bin/bash -ce 'for i in "$@";do curl -XPUT http://localhost:8008/v1/upload --upload-file "$i";done' --
```