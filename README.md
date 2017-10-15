# hello-gae-go
Google App Engine 用のサンプルプログラム

## 起動方法(プレーン)
### go コマンドからの起動
```sh
$ go run main.go
```
### GAE go plugin からの起動 
```sh
$ goapp serve appengine/app.yaml
```

## 起動方法(echo)
### go コマンドからの起動
```sh
$ go run app-standalone.go
```
### GAE go plugin からの起動
```sh
$ goapp serve gae_echo/app-engine.yaml
```


