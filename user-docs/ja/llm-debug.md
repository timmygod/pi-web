I'll look at the current working directory to understand the project context, though the main task is translating the provided Markdown.

Let me just translate the provided content as requested.

> このトラブルシューティングガイドは、ローカルモデル版のためにメンテナンスされています。
> ローカルデプロイの詳細とアップストリーム同期に関する注記は、
> [ローカルモデル版の開発](../../docs/dev/local-llm-development.md)と整合させてください。

ユーザーは pi-web を以下を介してインストールしています。

- pi install npm:@timmygod/pi-web-local

これにより、[install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) が自動的に実行され、pi がセットアップされます。

ユーザーに問題がある場合は、install.sh のセットアップが原因である可能性があります。何が起きているかをレビューし確認し、問題の正確な原因をユーザーに伝えることができます。修正してほしいかどうかを尋ねてください。必ずユーザーと確認してください。

ユーザーがモバイルやその他のネットワークから pi にアクセスできるようにするため、推奨される方法は tailscale を使用し、tailscale ネットワーク経由でアクセスすることです。その場合、ユーザーは tailscale ダッシュボードで HTTPs を有効にする必要があります - https://login.tailscale.com/admin/dns

tailscale のインストールがない場合、または tailscale を使用したくない場合。`pi-web status` を実行して、バイナリのパス、バイナリのステータス、アプリケーションにアクセスできるローカルエンドポイントを取得できます。ただし、http であるため、プッシュ通知は受信できません。

macOS では [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist) を使用します。
Linux では [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service) を使用します。

さらなるデバッグが必要な場合、および何が起きているかを確認する場合に。
