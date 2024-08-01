```bash
# 修改 username password host port databaseName
goctl model mysql datasource --url="username:password@tcp(host:port)/databaseName" -c -d ./model --idea --style goZero -t "*"
```

```bash
# 例子
goctl model mysql datasource --url="root:123123@tcp(127.0.0.1:3306)/user_server" -c -d . --idea --style goZero -t "*"
```