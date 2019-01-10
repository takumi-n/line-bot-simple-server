# line-bot-simple-server

LINE Messaging API を利用したシンプルなAPIサーバーです。
`/send-message` に以下の形式でJSONを `POST` するとBotと友達になっているユーザー全てにメッセージを送ります。

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
2. LINE Developersの管理画面から確認できる `Channel Secret` と `アクセストークン` の値をそれぞれ 環境変数 `CHANNEL_SECRET` `CHANNEL_TOKEN` に設定

LINE Developers の設定画面で必要な作業は以下の2点のみです。

1. `Webhook送信` を `利用する` に変更
2. `Webhook URL` を `<Heroku上にデプロイしたアプリのURL>/callback` に設定。（例: foobar.herokuapp.com/callback）
