<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dw/@timmygod/pi-web-local?label=downloads/wk&color=2ea043&cacheSeconds=86400)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb&cacheSeconds=86400)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · **日本語** · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

お使いの[pi](https://pi.dev)コーディングエージェントを、スマートフォン、タブレット、ノートパソコンから操作できます — ネットワーク内のどこからでも、あるいはTailscale経由でリモートからでも。

完全なPWAなので、任意のデバイスにインストールしてネイティブアプリのように使えます。自分専用のAIワークスペースと考えてください — ClaudeのCoworkのようなものですが、異なるモデルを使えます — モデルを切り替えてチャットしたり、スマホからコードを書いたり、あるいはあなたのマシン上で動く[パーソナルアシスタント](../en/personal-assistant.md)にすることもできます。

自分好みに：テーマやフォントを切り替え、自分の言語で使えます — pi-webは複数言語に対応しており、独自の言語を追加することもできます。さらに多くの機能が開発中ですが、肥大化はしません：不要な機能は設定でオフにできます。

</div>

## なぜこのローカルモデル版なのか？

元の pi-web は、共有機能と修正のためのアップストリームの基盤であり続けています。この版はその体験を維持し、ご自身のマシンや LAN 上の他の場所で実行されるモデルのための信頼性レイヤーを追加します。これらの環境では、生成はしばしば遅く、メモリは有限であり、長いコンテキストはそれ以外では健全なセッションを停滞させる可能性があります。

| 領域 | アップストリーム pi-web | この版 |
|------|-----------------|--------------|
| モデル/ランタイムポリシー | 標準的な pi-web の動作 | セッションごとの **Auto / Local / Cloud** モード、エンドポイント認識のローカル検出と永続的な手動オーバーライド付き |
| 長いコンテキストの処理 | 通常の pi の圧縮動作 | Local Mode は **65%** で積極的に圧縮し、別のプロバイダーリクエストの前に長いツール呼び出しループ内で再度チェックします |
| 圧縮の安全性 | 標準的な要約 | 有界のローリングチェックポイント、無効/上限に達した出力に対するより厳密な書き直し、そして無限の再圧縮の代わりに進捗なしの検出 |
| 中断された実行 | 通常のワーカーとエラー処理 | コンテキストオーバーフロー、思考のみの停止、および選択されたトランスポート中断に対する有界のリカバリー、永続的なループブレーカー付き |
| 手動レスキュー | 標準的なコンテキスト詳細 | **Force Compact** は会話を消去せずに明示的なリカバリーパスとして引き続き利用可能です |
| 互換性とリリース | 元のプロジェクトとリリースライン | ローカル専用の安全策は Local Mode の背後に留まります。Cloud Mode はアップストリームの動作を保持し、アップストリームの変更はここで独立してレビューおよびリリースされます |

これはアップストリームの書き直しや置き換えではありません。これは、脆弱な長時間実行セッションを受け入れることなく、ローカルモデルのプライバシーと制御を望む人々のために意図的に維持されている運用プロファイルです。ユーザー向けワークフローについては[ユーザーガイド](../en/README.md)を、実装と同期ポリシーについては[ローカルモデル版の開発](../../docs/dev/local-llm-development.md)を参照してください。

> [!WARNING]
> pi-webは現在**ベータ版**です。仕様の変更や互換性の破壊が発生します！

> [!TIP]
> 初めての方へ：機能の完全なツアー、インストール手順、ヒントについては **[ユーザーガイドを読む →](../en/README.md)** を参照してください。（[他の言語 →](../README.md)）

## スクリーンショット

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>デスクトップ</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile PWA" width="90%" /><br />
  <em>モバイルPWA</em>
</div>

## 全体の構成

```
 pi (ターミナル)                 ブラウザ (スマホ / タブレット / ノートPC)
      │                                │
      │  JSONLを書き込み               │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTPサーバー)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (セッションごと     (ライブ再読み込み)   (MagicDNS経由の
             のチャットワーカー)                      リモートHTTPS)
```

- **pi**は作業中に会話のJSONLを`~/.pi/agent/sessions/`に書き込みます。
- **pi-web**はGo製サーバーで、それらのファイルを読み取り、ブラウザでレンダリングし、SSE経由でライブ更新をストリーミングします。
- **pi --mode rpc**ワーカーがブラウザからのチャットを処理します — セッションごとに1つ、10分間アイドルで終了します。
- **fsnotify**がセッションディレクトリを監視し、新しい出力があれば数ミリ秒以内にブラウザが再読み込みされます。
- **Tailscale Serve**がローカルホストのサーバーをtailnet上のHTTPSエンドポイントとして公開します。

## インストール

```bash
pi install npm:@timmygod/pi-web-local@beta
```

以上です — 対応するバイナリをダウンロードし、自動起動を設定し、`/web`、`/pi-web`、`/remote`、`/refresh`コマンドを登録します。

インストール後、ブラウザで`http://127.0.0.1:31415`を開いてください。piからは`/web`を使って現在のセッションをブラウザで瞬時に開けます。Tailscaleがマシン上で動作している場合、pi-webは自動的にtailnet上のHTTPSエンドポイントを公開します — piから`/remote`を使うと、tailnet上の任意のデバイス向けのQRコードとURLを取得できます。

> **macOS のリモートアクセス:** Tailscale を対話的にインストールして開き、管理者プロンプトを承認してサインインしてください。その後 `/pi-web restart` を実行し、続けて `/remote` を実行します。

手動インストール、バイナリダウンロード、ソースからのビルドについては、[user-docs/install.md](../en/install.md)を参照してください。

## Pi連携

`pi install npm:@timmygod/pi-web-local@beta`を実行すると、以下が利用可能になります：

| コマンド | 機能 |
|---------|--------------|
| `/web` | 現在のセッションをブラウザで開く（SSH対応：ブラウザをスキップしてURLのみ表示） |
| `/pi-web` | ステータス、バージョンの表示、サーバーの起動/停止/再起動、または更新 |
| `/remote` | Tailscale経由のリモートアクセス用のQRコードとURLを表示 |
| `/refresh` | リモートブラウザから書き込まれた新しいメッセージをターミナルセッションに取り込む |

セッションの**自動タイトル付け**はpi-web自体に組み込まれており、`/settings`ページで設定できます。**デフォルトでオン**になっており、セッションに自動的に名前を付けます。選択できるオプション：

- **タイトルを付けるタイミング** — セッションごとに1回、または新しいメッセージごと（デフォルト）。
- **タイトルモデル** — デフォルトは無料で高速な**組み込み単語ヒューリスティック（AI不使用）**、またはモデル（例：小さくて高速なもの）を選んで、より賢いモデル生成タイトルを利用できます。

このパッケージはpi-webバイナリを`~/.pi/agent/bin/pi-web`にインストールし、ログイン時の自動起動も設定します。

## ログイン時の自動起動

`pi install npm:@timmygod/pi-web-local@beta`コマンドが自動的に設定します：

| OS | 仕組み |
|----|-----------|
| macOS | `~/Library/LaunchAgents/com.pi-web.plist`のlaunchd plist |
| Linux | `~/.config/systemd/user/pi-web.service`のsystemdユーザーサービス |

リモートアクセス用のトークンを設定するには、`~/.config/pi-web/env`を作成します：

```
PI_WEB_TOKEN=your-token-here
```

詳細（手動セットアップ、カスタムポート、非ループバックバインド）については、[user-docs/install.md](../en/install.md)を参照してください。

## 開発

```bash
make setup   # フロントエンド依存関係のインストールとGoモジュールのダウンロード
make check   # フロントエンドのテスト/ビルド + Goのテスト/vet
make build   # 必要に応じてsetupを実行し、フロントエンドをビルドし、その後./pi-webをビルド
```
