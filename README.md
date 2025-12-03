# typora-pic-server
A picture server for Typora in custom command

copy the command to Typora Image Settings

```shell
/bin/bash -ce 'for i in "$@";do curl -XPUT http://172.16.106.106:8008/v1/upload --upload-file "$i";done' --
```