# line-bot-simple-server

LINE Messaging API を利用したシンプルなAPIサーバーです。
`/send-message` に以下の形式でJSONを `POST` すると友達になっているユーザー全てにメッセージを送ります。

```json
{
    "message": "送りたいメッセージ"
}
```

簡易的なLINE通知用APIサーバーとして使えます。

## セットアップ

Heroku上で動作させることを想定しています。
PostgresがHeroku上にインストールされている環境に対してデプロイするだけで動作します。