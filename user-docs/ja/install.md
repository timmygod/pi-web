# インストールと使い方

## 機能

### リモート操作

- テキストや画像の添付つきで、ブラウザからあらゆるセッションを継続
- Web UI から直接、あらゆるプロジェクトパスに対してまったく新しいセッションを開始
- セッションごとに、ブラウザ内でモデルの切り替えと思考レベルのセレクタを選択
- セッションごとのワーカー状態（idle / running / error）、クラッシュ時の自動復旧付き
- 複数のセッションが並行して実行 — 1 つで作業を開始し、別のセッションのストリーミングを監視
- `PI_WEB_TOKEN` による安全な LAN 公開 — 明示的な非ループバックバインドにはデフォルトで必須

### セッションの閲覧

- フィルタ、検索、完全なブランチナビゲーションつきで、プロジェクトまたぎのセッションを閲覧
- pi がまだ実行中の間も、ライブの増分更新（fsnotify 経由; 遅延は ~ms）
- アクティブなセッションを追尾するフォロモード
- 個々のメッセージへの深リンク
- セッションを JSONL としてダウンロード
- シェア静スナップショットをシークレットの GitHub Gist として共有
- セッションのオープン、リモート QR、セッション同期、トークン管理のための `/web`、`/remote`、`/refresh`、`/pi-web token`、`/pi-web set-token` pi 拡張
- セッションがスケジュール、プロジェクトのスクラッチパッド、設定を自然言語で管理できるようにする `/skill:pi-web-schedule`、`/skill:pi-web-notes`、`/skill:pi-web-settings`（`pi-web-ctl`）

## セッションモードの選択

このエディションは、pi にすでに設定されているプロバイダとモデルを使用します。Local Mode はランタイムのポリシーであり、別のモデルインストーラや第 2 の API キー画面ではありません。セッション作成時にモードを選択するか、現在の実行が落ち着いてから変更できます：

| モード | 使用場面 | 挙動 |
|------|-------------|----------|
| **Auto** | pi-web に決定させたい場合 | 可能であればプロバイダメタデータから local/LAN のエンドポイントに解決; それ以外は通常パスを維持 |
| **Local** | モデルがこのマシンやあなたの LAN で実行されている場合 | 65% のコンパクション境界、バウンド付きのチェックポイント、Force Compact、ガード付きの自動復旧を有効化 |
| **Cloud** | 選択されたモデルがホストされており、アップストリームの挙動に従うべき場合 | local 専用のコンパクションと復旧ポリシーをセッションから除外 |

手動での Local または Cloud の選択は自動検出を優先し、リロードや再起動をまたいで持続します。実行中のセッションはワーカーが落ち着くまでモード変更を拒否するため、UI に表示されるモードは実際に使用されているポリシーと常に一致します。

## 要件

- [Go](https://go.dev) 1.25+（ソースからのビルドが必要な場合のみ）
- ブラウザでのチャット/モデル切り替えのために `PATH` 上の `pi`
- 任意: 共有のための `gh`
- Windows の場合: pi はシェルツールに bash シェルが必要です — [Git for Windows](https://git-scm.com/download/win) で十分です（pi の Windows ドキュメントを参照）

## インストール

### Pi パッケージ（推奨）

```bash
pi install npm:@timmygod/pi-web-local
```

この単一のコマンドが:
- pi のパッケージディレクトリ配下に npm の pi パッケージをインストール
- パッケージの `postinstall` スクリプト（Windows では `install.sh`、または `install.ps1`）を実行
- GitHub Releases から、パッケージバージョンとプラットフォームに一致する pi-web バイナリをダウンロード
- `~/.pi/agent/bin/pi-web`（Windows では `pi-web.exe`）にインストール
- ログイン時の自動開始を設定（macOS では launchd、Linux では systemd、Windows では Run キーのランチャー）
- `/web`、`/remote`、`/refresh`、`/pi-web token`、`/pi-web set-token` の pi コマンドを登録

セッションの自動タイトル付けは pi-web に組み込まれています（拡張機能ではなく）、`/settings` ページで設定します。デフォルトではオンになっています: pi-web は無料の組み込みワードヘウリスティクス（AI を使わない）でセッションを自動的に命名し、新しいメッセージごとにタイトル付けし直します。セッションごとに一度だけタイトル付けにするように切り替えたり、あるいはヘウリスティクスの代わりにより賢いタイトルを書くモデルを選択したりできます。

Linux では、自動開始は `~/.config/systemd/user/pi-web.service` にあるユーザ systemd サービスとして設定されます。インストーラはその `ExecStart` を実際にインストールされたバイナリのパスに書き換えます。ランタイムで Tailscale が利用可能な場合、pi-web は Tailscale Serve HTTPS でローカルホストのサーバーを公開します。ユーザ systemd が利用できない場合、`~/.pi/agent/bin/pi-web -o` で手動で実行してください。

特定のプロジェクト用のみインストールする場合（`.pi/settings.json` 経由でチームと共有）:

```bash
pi install -l npm:@timmygod/pi-web-local
```

その後に pi を再起動（または `/reload` を実行）し、`/web`、`/pi-web`、`/remote`、`/refresh` を使用します。アクセストークンは `/pi-web token` と `/pi-web set-token` で管理します。

npm が `@timmygod/pi-web-local` のリネーム中に `ENOTEMPTY` で中止する場合、npm の古い非表示バックアップディレクトリを削除し、パッケージを再インストールしてください:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### クイックインストール（ビルドツール不要）

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

これは最新バージョンの pi-web バイナリをダウンロードし、`/usr/local/bin`（Windows では `~/.pi/agent/bin`）にインストールし、ログイン時の自動開始を設定します。Go、Node、pi は不要です。

### バイナリのダウンロード

ビルド済みのバイナリは各 [GitHub Release](https://github.com/timmygod/pi-web/releases) に添付されています。

```bash
# macOS (Apple Silicon)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-arm64
chmod +x pi-web

# macOS (Intel)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-amd64
chmod +x pi-web

# Linux (amd64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-amd64
chmod +x pi-web

# Linux (arm64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-arm64
chmod +x pi-web
```

```powershell
# Windows (x64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-amd64.exe

# Windows (ARM64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-arm64.exe
```

その後に PATH へ移動します:

```bash
cp pi-web ~/.pi/agent/bin/
# またはシステム全体向け:
sudo cp pi-web /usr/local/bin/
```

### ソースからのビルド

このチェックアウトは pi-web のローカルモデルエディションです。通常のビルドは Web アプリケーションとバックエンドを一緒に生成します。ローカルモデルのセーフガードは、別々のバイナリによってではなく、セッションの有効な Local Mode によってランタイムで有効化されます。

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # Vite バンドルをビルドしてから、Go バイナリに埋め込む

# 任意: PATH へ配置
cp pi-web ~/.pi/agent/bin/
```

フロントエンドのバンドルは `web/assets_embed.go` によって埋め込まれるため、`go build` にはまず `web/dist` の存在が必要です。`make build` は両方のステップを順序通りに実行します。手動でビルドする場合は、`go build ./cmd/pi-web` の前に `npm --prefix web install && npm --prefix web run build` を実行してください。

メンテナンスされているフォークのワークフロー、アップストリームへの同期、Local Mode の検証チェックリストについては、[ローカルモデル開発メモ](../../docs/dev/local-llm-development.md)を参照してください。

### インストール済みインスタンスと並行して開発

インストール済みインスタンスをポート `31415` で実行したままにし、ソースチェックアウトを開発モードで開始します:

```bash
make dev
```

`http://127.0.0.1:31416` を開きます。`make dev` は内部の `PI_WEB_DEV=1` 開発環境を設定するため、ソースチェックアウトはインストール済みインスタンスとセッション、設定、SQLite データを共有しつつ、別々の開発ランタイムロックと状態ファイルを保持します。通常のインストール済みインスタンスと手動で開始したインスタンスは変更されず、元の単一インスタンスの挙動を維持します。

重複した自律作業を防ぐため、開発モードではスケジュールループ、チャットキューのドレイン、自動タイトル付け、プッシュ通知は実行されません。開発 UI 経由の直接リクエストは引き続き機能します。同じチャットセッションを同時に両方のインスタンスから操作しないでください。各プロセスには独自の RPC ワーカーマネージャがあります。

`make dev` は Go のホットリロードに [Air](https://github.com/air-verse/air) が必要です:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` は開発ハーネスの配線であり、サポートされる本番の複数インスタンスモードではありません。

## 削除

```bash
pi remove npm:@timmygod/pi-web-local
```

これはパッケージの `preuninstall` スクリプト（Windows では `uninstall.sh`、または `uninstall.ps1`）を実行し、実行中のインスタンスを停止して、以下のものを削除します:

- pi-web バイナリ（`~/.pi/agent/bin/pi-web`、またはスタンドアロンインストールの場合は `/usr/local/bin/pi-web`）
- バージョンファイル（`~/.pi/agent/pi-web-version`）
- ランタイムの状態ファイル（`~/.pi/agent/pi-web/pi-web-state.json`）
- 自動開始の設定（macOS では launchd の plist、Linux では systemd のユーザサービス、Windows では Run キーのエントリ + ランチャースクリプト）

あなたのデータは保持されるため、後日の再インストールで中断したところから再開できます: `~/.pi/agent/pi-web.sqlite`、`~/.pi/agent/pi-web-memory.sqlite`、`~/.pi/agent/sessions/` 配下のセッションファイル、`~/.config/pi-web/env`（`PI_WEB_TOKEN` を含む）。完全なリセットが必要な場合は、これらを手動で削除してください。

## 使い方

```bash
# デフォルトポート（31415）で起動
pi-web

# 起動してブラウザを開く
pi-web -o

# カスタムポート
pi-web -p 8080

# バインドホストの上書き（ループバックはデフォルトで認証なし）
pi-web --host 127.0.0.1

# 非ループバックバインドにはトークンが必要 — 設定しない場合、pi-web は起動を拒否する
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

デフォルトでは、pi-web は `127.0.0.1` にバインドします。Tailscale が MagicDNS とともに実行中 **かつ `PI_WEB_TOKEN` が設定されている場合**、pi-web は `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` も実行し、HTTPS のテイルネット URL を出力します。トークンがない場合、pi-web はループバックのみに留まり Tailscale Serve をスキップするため、テイルネットのピアは認証なしでエージェントに到達できません。明示的な非ループバックバインドも `PI_WEB_TOKEN` の設定が必要です。ローカルテスト用に上書きするには `--insecure` を渡してください。

## リモートアクセス

pi-web をローカルでリッスンさせ、テイルネット上のスマートフォンやラップトップから出力された Tailscale の HTTPS URL を使用します。

macOS では、Tailscale を対話的にインストールして開き、管理者プロンプトを承認し、サインインします。その後に `/pi-web restart` を実行し、`/remote` を続行します。

Linux では、pi-web のインストール/実行の前に、ユーザが Tailscale を管理できることを許可してください。さもないと `tailscale serve` が sudo を必要とする可能性があり、自動開始が失敗する可能性があります:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. トークン付きで pi-web を起動し、Tailscale HTTPS エンドポイントを公開させる
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Tailscale に接続されている他のデバイスから、出力された
#    "Tailscale HTTPS" URL を開き、トークンを一度入力する。
```

> デフォルトでは、pi-web は `PI_WEB_TOKEN` が設定されていない限り非ループバックアドレスへのバインドを拒否します — そうでない場合、バインドされたアドレスに到達できる誰もがセッションを閲覧し、pi に指示を送る可能性があるためです。ローカルネットワークのテスト用にこのガードを上書きするには `--insecure` を渡してください。 **`--insecure` は Tailscale やあなたのマシン外部から到達可能なアドレスでは使用しないでください。**
>
> クライアントはトークンを `Authorization: Bearer <token>` ヘッダー、`X-Pi-Token` ヘッダー、または一度だけ `?token=<token>` 経由（これは後続のリクエスト用に `pi_token` クッキーを設定する）で渡せます。`?token=` 経由で渡されたトークンはブラウザの履歴、サーバーのアクセスログ、ページ上のリンクからの `Referer` ヘッダーに出力されてしまいます — 最初のブックマーク以降はヘッダー形式を推奨します。

## ブラウザチャット

セッションページを開き、下部のコンポーザでそのセッションをそのまま継続します。

- `Enter` で送信、`Shift+Enter` で改行を入力
- 画像をコンポーザへドラッグ&ドロップ、または直接貼り付け
- モデルピッカーと思考レベルのセレクタはヘッダーにあり — 変更はすぐに下部の pi ワーカーに適用
- 各アクティブなセッションは専用の `pi --mode rpc` ワーカーを持ち、そのため異なるセッションは互いにブロックし合わない

## セッションの共有

セッションページで **Share** をクリックすると、シークレットの GitHub Gist が作成されます。

要件:
- `gh` のインストール
- `gh auth login` の完了

共有の結果として返されるもの:
- シークレットの gist URL
- `https://pi.dev/session/#<gistId>` でのプレビュー URL

共有された gist はスナップショットであり、ライブ更新されません。

## ログイン時の自動開始

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# systemd ユーザサービスをインストール
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# 任意: 非ループバックバインド用の PI_WEB_TOKEN を設定
# （または pi の中で /pi-web set-token <token> を使用）
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# 有効化して開始
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# 状態の確認
systemctl --user status pi-web.service

# ログの表示
journalctl --user -u pi-web.service -f
```

> サービスを起動時に開始（ログイン前）するには、代わりにシステムサービスを使用してください: `init/pi-web.service` を `/etc/systemd/system/` にコピーし、`sudo systemctl` を使用します。

### Windows

インストーラは、管理者権限を必要とせずにこれを自動的に設定します: `HKCU\Software\Microsoft\Windows\CurrentVersion\Run` 配下の `pi-web` エントリがログイン時に `~/.config/pi-web/pi-web-start.vbs` を開始し、これが `~/.config/pi-web/env`（`PI_WEB_TOKEN`、`PATH` など）を読み込んだ後にバイナリを非表示（コンソールウィンドウなし）で開始します。

手動で管理するには:

```powershell
# 開始 / 停止
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# 自動開始の削除
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Windows にはサービスの監督はありません: pi-web がクラッシュしても、次のログインまで停止したままです（他のプラットフォームでは launchd/systemd が自動的に再起動します）。
