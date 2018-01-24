# WebsocketClient

websocketクライアントです。

## 実行ファイル一覧
`bin`ディレクトリにあります。

 ファイル名 | 対象OS | 備考
---------|----------|---------
 wsClient_x86.exe | windows 32bit | なし
 wsClient_x64.exe | windows 64bit | なし
 wsClient_OSX | OSX | なし

## jsonのスキーマ
このwebsocketクライントはjsonでやり取りを行う。

jsonのスキーマを以下の通り
```
{
    "Name": "string",
    "Message": "string"
}
```

## 実行方法
```
go run wsClient.go
```

もしくは

``` 
./wsServer_x64.exe (windows64bitの例)
```

localhost:8080にwebsocket接続を開始。

実際の接続URLは`ws://localhost:8080/ws`

## オプション

### addr: 接続先のアドレスを指定する

```
./wsServer_x64.exe -addr　billboard-wsserver.herokuapp.com
```

`ws://billboard-wsserver.herokuapp.com/ws` にwebsocket接続を開始

### name: クライアント名を指定

```
./wsServer_x64.exe -name user
```

jsonの`name`にuserを指定
