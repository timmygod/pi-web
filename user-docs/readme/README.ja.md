<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · **日本語** · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

スマホ、タブレット、ラップトップから [pi](https://pi.dev) のコーディングエージェントを操作しよう — ネットワーク内ならどこからでも、Tailscale 経由ならリモートからも。

フル PWA なので、どのデバイスでもインストールしてネイティブアプリのように使える。自分だけのパーソナルな AI ワークスペースと思えばいい — Claude の Cowork に似ているが、使えるモデルが違う。モデル間を横断してチャットし、スマホからコードを書き、自分のマシンに常駐する [パーソナルアシスタント](../en/personal-assistant.md) にすることもできる。

自分好みにカスタマイズしよう: テーマやフォントを切り替え、自分の言語で使える — pi-web は複数言語に対応しており、独自の言語も追加できる。今後の機能追加も続くが、肥大化することはない: 必要のないものは設定でオフにできる。

</div>

## なぜこのローカルモデル版なのか?

オリジナルの pi-web は、共有機能と修正のアップストリーム基盤であり続ける。この版は、その体験を維持したうえで、自分のマシンや LAN 内の他の場所で動くモデル向けの信頼性レイヤーを追加する。こうした環境では、生成はしばしば遅く、メモリは有限で、長いコンテキストがそうでなければ健全なセッションを停滞させてしまうことがよくあるためだ。

| 領域 | アップストリーム pi-web | この版 |
|------|-----------------|--------------|
| モデル/ランタイム方針 | 標準の pi-web の動作 | セッションごとに **Auto / Local / Cloud** モード、エンドポイントに応じたローカル検出と永続的な手動オーバーライド付き |
| 長コンテキスト処理 | 通常の pi のコンパクション動作 | Local Mode は **65%** で能動的にコンパクションし、次のモデル要求の前に長いツール呼び出しループの中で再度確認する |
| コンパクションの安全性 | 標準の要約 | 限界付きローリングチェックポイント、無効/上限超過の出力に対するより厳格な書き換えを 1 回、そして無限の再コンパクションの代わりに進捗なし検出 |
| 中断した実行 | 通常のワーカーとエラー処理 | コンテキストオーバーフロー、thinking 専用の停止、選択されたトランスポート中断に対する限界付きの復旧、および永続的なループブレーカー付き |
| 手動レスキュー | 標準のコンテキスト詳細 | **Force Compact** は、会話を消去することなく明示的な復旧パスとして利用可能 |
| 互換性とリリース | オリジナルプロジェクトとリリースライン | ローカル専用セーフガードは Local Mode の裏側に留まり、Cloud Mode はアップストリームの動作を維持、アップストリームの変更はここで独立してレビュー・リリースされる |

これは、アップストリームの上書きや置換ではない。ローカルモデルによるプライバシーと制御を、脆い長時間実行セッションを受け入れずに求めている人々のために、意図的にメンテナンスされた運用プロファイルだ。ユーザー向けのワークフローは [ユーザーガイド](../en/README.md) を、実装と同期ポリシーは [ローカルモデル版の開発](../../docs/dev/local-llm-development.md) を参照すること。

> [!TIP]
> 初めての方は、機能の全般的な紹介、インストール手順、ヒントを含む **[ユーザーガイドを読む →](../en/README.md)** へ。 ([他の言語 →](../README.md))

## スクリーンショット

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>デスクトップ</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>モバイル</em>
</div>

## 構成の仕組み

```
 pi (terminal)                 Browser (phone / tablet / laptop)
      │                                │
      │  writes JSONL                  │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (per‑session       (live reload)      (remote HTTPS
             chat worker)                           via MagicDNS)
```

- **pi** は作業中に、会話の JSONL を `~/.pi/agent/sessions/` に書き込む。
- **pi-web** は Go で書かれたサーバーで、これらのファイルを読み取り、ブラウザでレンダリングし、SSE 経由でライブ更新をストリーミングする。
- **pi --mode rpc** ワーカーはブラウザから始まるチャットを処理する — セッションごとに 1 つ、10 分アイドルで回収される。
- **fsnotify** はセッションディレクトリを監視し、新しい出力から数ミリ秒以内にブラウザを再読み込みする。
- **Tailscale Serve** は、localhost サーバーを tailnet 上の HTTPS エンドポイントとして公開する。

## インストール

```bash
pi install npm:@timmygod/pi-web-local
```

これだけです — 対応するバイナリをダウンロードし、オートスタートを設定し、`/web`、`/pi-web`、`/remote`、`/refresh` コマンドを登録します。

インストール後、ブラウザで `http://127.0.0.1:31415` を開く。pi から `/web` を使うと、現在のセッションをブラウザですぐに開ける。マシンの上で Tailscale が実行されていれば、pi-web は tailnet 上に HTTPS エンドポイントを自動公開する — pi から `/remote` を使うと、tailnet 上のあらゆるデバイス用の QR コードと URL を取得できる。

> **macOS のリモートアクセス:** Tailscale を対話的にインストール・起動し、管理者プロンプトを承認してサインインする。そのあと `/pi-web restart` を実行し、続けて `/remote` を実行する。

手動インストール、バイナリのダウンロード、ソースからのビルドについては、[user-docs/install.md](../en/install.md) を参照。

## Pi 連携

`pi install npm:@timmygod/pi-web-local` を実行すると、以下が得られる:

| コマンド | 動作 |
|---------|--------------|
| `/web` | 現在のセッションをブラウザで開く (SSH 対応: ブラウザをスキップして URL だけ表示) |
| `/pi-web` | ステータス、バージョンの表示、サーバーの開始/停止/再起動、または更新 |
| `/remote` | Tailscale 経由のリモートアクセス用の QR コードと URL を表示 |
| `/refresh` | リモートブラウザから書かれた新しいメッセージをターミナルセッションに引き込む |

セッションの**自動タイトル付け**は pi-web 自体に組み込まれており、`/settings` ページで設定できる。**既定ではオン**で、セッションを自動的に名付ける。選択可能なのは:

- **タイトル付けのタイミング** — セッションごとに 1 回、または新しいメッセージのたびに (既定)。
- **タイトル付け用モデル** — 既定では無料で即時の **組み込み単語ヒューリスティック (AI 不使用)**、またはより賢くモデルが書いたタイトルのためにモデル (例: 小型/高速なもの) を選択。

このパッケージは、pi-web バイナリを `~/.pi/agent/bin/pi-web` にインストールし、ログイン時のオートスタートも設定する。

## ログイン時のオートスタート

`pi install npm:@timmygod/pi-web-local` コマンドは、これを自動的にセットアップする:

| OS | 仕組み |
|----|-----------|
| macOS | `~/Library/LaunchAgents/com.pi-web.plist` への launchd plist |
| Linux | `~/.config/systemd/user/pi-web.service` への systemd ユーザーサービス |
| Windows | `~/.config/pi-web/` の非表示スタターを起動する `HKCU` Run キーのエントリ |

リモートアクセス用のトークルを設定するには、`~/.config/pi-web/env` を作成する:

```
PI_WEB_TOKEN=your-token-here
```

詳細 (手動セットアップ、カスタムポート、非ループバックバインド) については、[user-docs/install.md](../en/install.md) を参照。

## 開発

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

アップストリーム同期、ローカルモデルのテスト、並行リリース
ワークフローについては、[ローカルモデル版の開発](../../docs/dev/local-llm-development.md) を参照。
