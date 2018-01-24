# WebsocketServer

websocketサーバーです。

heroku [https://billboard-wsserver.herokuapp.com/](https://billboard-wsserver.herokuapp.com/)

websocket接続用URLは billboard-wsserver.herokuapp.com/ws

## ローカル版実行ファイル一覧
`bin`ディレクトリにあります。

 ファイル名 | 対象OS | 備考
---------|----------|---------
 wsServer_x86.exe | windows 32bit | なし
 wsServer_x64.exe | windows 64bit | なし
 wsServer_OSX | OSX | なし

## jsonのスキーマ
このwebsocketサーバーはjsonでやり取りを行う。

jsonのスキーマを以下の通り
```
{
    "Name": "string",
    "Message": "string"
}
```

## 実行方法

`./wsServer_x64.exe` (windows64bitの場合)

もしくは

`go run main.go` (golangがインストール済みなら)

localhost:8080でwebsocketを待ち受けます。

## オプション

port: listenするポートを指定

`./wsServer_x64.exe -port 8000`

8000番ポートで待受