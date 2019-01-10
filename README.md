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
Heroku上で必要な作業は以下の2点のみです。

1. Postgres アドオンをインストール
2. 環境変数 `CHANNEL_SECRET` （Channel Secret） `CHANNEL_TOKEN` （アクセストークン）を設定

LINE Developers の設定画面で必要な作業は以下の2点のみです。

1. `Webhook送信` を `利用する` に変更
2. `Webhook URL` を `<Heroku上にデプロイしたアプリのURL>/callback` に設定。（例: foobar.herokuapp.com/callback）
