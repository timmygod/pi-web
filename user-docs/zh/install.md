# 安装与使用

## 功能特性

### 远程控制

- 通过浏览器继续任意会话，支持文本或图片附件
- 直接通过 Web UI 针对任意项目路径启动全新会话
- 浏览器内的模型切换和 thinking-level 选择器，按会话独立配置
- 按会话显示 worker 状态（空闲 / 运行中 / 错误），崩溃时自动恢复
- 多个会话并行运行 —— 在一个会话中启动工作，同时观察另一个会话的流式输出
- `PI_WEB_TOKEN` 用于安全地暴露到局域网 —— 任何显式的非回环地址绑定默认都需要它

### 浏览会话

- 跨项目浏览会话，支持过滤器、搜索和完整的分支导航
- pi 仍在运行时通过 fsnotify 提供实时增量更新（延迟约毫秒级）
- 跟随模式，用于实时查看活跃会话
- 指向单条消息的深链接
- 将会话下载为 JSONL
- 以私有 GitHub Gist 的形式分享静态快照
- 通过 `/web`、`/remote`、`/refresh`、`/pi-web token` 和 `/pi-web set-token` 这些 pi 扩展来打开会话、远程 QR 登录、同步会话以及管理 token
- 通过 `/skill:pi-web-schedule`、`/skill:pi-web-notes`、`/skill:pi-web-settings`（`pi-web-ctl`），让会话能够用自然语言管理调度、项目草稿本和设置

## 选择会话模式

本版本使用 pi 中已配置好的 providers 和模型；Local Mode 是一项运行时策略，而不是一个独立的模型安装器或第二套 API-key 界面。
在创建会话时选择模式，或在当前运行稳定后再更改：

| 模式 | 适用场景 | 行为 |
|------|-------------|----------|
| **Auto** | 希望由 pi-web 自行决定 | 尽可能从 provider 元数据中解析本地/局域网端点；否则保持正常路径 |
| **Local** | 模型运行在本机或你的局域网中 | 启用 65% 压缩边界、有界检查点、Force Compact 以及带保护机制的自动恢复 |
| **Cloud** | 所选模型是托管的，应遵循上游行为 | 让本地专用的压缩和恢复策略不作用于该会话 |

手动选择 Local 或 Cloud 会优先于自动检测，并跨重载和重启持久保留。
运行中的会话会拒绝模式更改，直到其 worker 稳定下来，因此 UI 中显示的模式始终与当前实际生效的策略一致。

## 前置要求

- [Go](https://go.dev) 1.25+（仅从源码构建时需要）
- `PATH` 中可用的 `pi`，用于浏览器聊天/模型切换
- 可选：`gh`，用于分享
- Windows 上：pi 的 shell 工具需要 bash —— 安装 [Git for Windows](https://git-scm.com/download/win) 即可（见 pi 的 Windows 文档）

## 安装

### Pi 包（推荐）

```bash
pi install npm:@timmygod/pi-web-local
```

这一条命令会：
- 在 pi 的包目录下安装该 npm pi 包
- 运行包的 `postinstall` 脚本（Windows 上为 `install.sh`，或 `install.ps1`）
- 从 GitHub Releases 下载与你包版本和平台匹配的 pi-web 二进制文件
- 安装到 `~/.pi/agent/bin/pi-web`（Windows 上为 `pi-web.exe`）
- 设置登录时自动启动（macOS 上用 launchd，Linux 上用 systemd，Windows 上用 Run 键启动器）
- 注册 `/web`、`/remote`、`/refresh`、`/pi-web token` 和 `/pi-web set-token` 这些 pi 命令

会话自动命名已内置于 pi-web（而非扩展中），并在 `/settings` 页面进行配置。默认开启：pi-web 使用免费的内置词汇启发式（不依赖 AI）自动为会话命名，并在每条新消息时重新命名。你可以切换为每个会话只命名一次，和/或选择某个模型来替代启发式、写出更智能的标题。

在 Linux 上，自动启动被配置为一个用户 systemd 服务，位于 `~/.config/systemd/user/pi-web.service`。安装器会将其 `ExecStart` 重写为实际安装的二进制文件路径。如果运行时 Tailscale 可用，pi-web 会通过 Tailscale Serve HTTPS 发布 localhost 服务。如果用户 systemd 不可用，可以用 `~/.pi/agent/bin/pi-web -o` 手动运行。

如果只想为某个特定项目安装（通过 `.pi/settings.json` 与团队共享）：

```bash
pi install -l npm:@timmygod/pi-web-local
```

然后重启 pi（或运行 `/reload`），并使用 `/web`、`/pi-web`、`/remote`、`/refresh`。用 `/pi-web token` 和 `/pi-web set-token` 管理你的访问 token。

如果 npm 在重命名 `@timmygod/pi-web-local` 时以 `ENOTEMPTY` 中止，请删除 npm 遗留的隐藏备份目录并重新安装包：

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### 快速安装（无需构建工具）

macOS / Linux：

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows（PowerShell）：

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

这会下载最新的 pi-web 二进制文件，安装到 `/usr/local/bin`（Windows 上为 `~/.pi/agent/bin`），并设置登录时自动启动。无需 Go、Node 或 pi。

### 下载二进制文件

预构建的二进制文件附在每个 [GitHub Release](https://github.com/timmygod/pi-web/releases) 中。

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

然后把它放入你的 PATH：

```bash
cp pi-web ~/.pi/agent/bin/
# 或系统范围：
sudo cp pi-web /usr/local/bin/
```

### 从源码构建

这个检出是 pi-web 的本地模型版本。正常构建会同时生成 Web 应用和后端；本地模型保护机制由会话生效的 Local Mode 在运行时启用，而不是通过单独的二进制文件。

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # 构建 Vite bundle，然后将其嵌入到 Go 二进制文件中

# 可选：放入 PATH
cp pi-web ~/.pi/agent/bin/
```

前端 bundle 由 `web/assets_embed.go` 嵌入，因此 `go build` 需要 `web/dist` 已经存在。`make build` 按顺序执行这两步；如果你手动构建，请先在 `go build ./cmd/pi-web` 之前运行 `npm --prefix web install && npm --prefix web run build`。

关于维护中的 fork 工作流、上游同步以及 Local Mode 验证清单，请参阅 [本地模型开发笔记](../../docs/dev/local-llm-development.md)。

### 与已安装实例并行开发

让已安装实例继续在 `31415` 端口运行，然后以开发模式启动源码检出：

```bash
make dev
```

打开 `http://127.0.0.1:31416`。`make dev` 会设置内部的 `PI_WEB_DEV=1`
开发环境变量，因此源码检出会与已安装实例共享会话、设置和
SQLite 数据，同时保持独立的开发运行时锁和状态文件。正常安装和手动启动的实例
不受影响，并保留原有的单实例行为。

为避免重复的自主工作，开发模式不会运行
调度循环、聊天队列排空器、自动命名或推送通知。通过开发 UI 发出的直接
请求仍然有效。不要同时从两个实例驱动同一个
聊天会话；每个进程都有自己的 RPC worker
管理器。

`make dev` 需要 [Air](https://github.com/air-verse/air) 来实现 Go 热重载：

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` 是开发工具的管线机制，并非受支持的生产
多实例模式。

## 卸载

```bash
pi remove npm:@timmygod/pi-web-local
```

这会运行包的 `preuninstall` 脚本（Windows 上为 `uninstall.sh`，或 `uninstall.ps1`），它会停止运行中的实例并删除：

- pi-web 二进制文件（`~/.pi/agent/bin/pi-web`，独立安装则为 `/usr/local/bin/pi-web`）
- 版本文件（`~/.pi/agent/pi-web-version`）
- 运行时状态文件（`~/.pi/agent/pi-web/pi-web-state.json`）
- 自动启动配置（macOS 上的 launchd plist，Linux 上的 systemd 用户服务，Windows 上的 Run 键条目 + 启动器脚本）

你的数据会被保留，以便后续重新安装时接着上次继续：
`~/.pi/agent/pi-web.sqlite`、`~/.pi/agent/pi-web-memory.sqlite`、位于 `~/.pi/agent/sessions/` 下的会话文件，以及 `~/.config/pi-web/env`（包含
`PI_WEB_TOKEN`）。如果想要彻底清空，请手动删除这些文件。

## 使用

```bash
# 在默认端口（31415）上启动
pi-web

# 启动并打开浏览器
pi-web -o

# 自定义端口
pi-web -p 8080

# 覆盖绑定主机（回环默认无认证）
pi-web --host 127.0.0.1

# 非回环绑定需要 token —— 否则 pi-web 拒绝启动
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

默认情况下，pi-web 绑定到 `127.0.0.1`。如果 Tailscale 正在运行且启用了 MagicDNS，**并且设置了 `PI_WEB_TOKEN`**，pi-web 还会运行 `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` 并打印 HTTPS tailnet URL。如果没有 token，pi-web 仅保持回环绑定并跳过 Tailscale Serve，因此 tailnet 上的对端无法在未认证的情况下访问 agent。任何显式的非回环绑定也需要设置 `PI_WEB_TOKEN`；对于本地测试，可以传 `--insecure` 来覆盖。

## 远程访问

让 pi-web 继续在本机监听，然后在 tailnet 上的手机或笔记本上使用打印出的 Tailscale HTTPS URL。

在 macOS 上，交互方式安装并打开 Tailscale，批准管理员提示，并登录。然后运行 `/pi-web restart`，接着运行 `/remote`。

在 Linux 上，在安装/运行 pi-web 之前，允许你的用户管理 Tailscale，否则 `tailscale serve` 可能需要 sudo，自动启动可能会失败：

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. 带 token 启动 pi-web，以便它发布 Tailscale HTTPS 端点
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. 从任何其他连接 Tailscale 的设备，打开打印出的
#    "Tailscale HTTPS" URL，并只输入一次 token。
```

> 默认情况下，除非设置了 `PI_WEB_TOKEN`，pi-web 拒绝绑定到非回环地址 —— 否则任何能访问该绑定地址的人都可能查看会话并向 pi 发送指令。对于局域网测试，传 `--insecure` 可覆盖此保护。**不要在 Tailscale 或任何可从机器外部访问的地址上使用 `--insecure`。**
>
> 客户端可以通过 `Authorization: Bearer <token>` 头、`X-Pi-Token` 头，或一次性通过 `?token=<token>` 传递 token（后者会设置 `pi_token` cookie 用于后续请求）。通过 `?token=` 传递的 token 会进入浏览器历史、服务器访问日志，以及页面上任何链接产生的 `Referer` 头 —— 除了初始书签外，更推荐使用头的方式。

## 浏览器聊天

打开一个会话页面，使用底部的编辑器继续该精确会话。

- `Enter` 发送，`Shift+Enter` 插入换行
- 直接将图片拖放或粘贴到编辑器中
- 模型选择器和 thinking-level 选择器位于头部 —— 更改会立即作用于底层 pi worker
- 每个活跃会话拥有自己专属的 `pi --mode rpc` worker，因此不同会话不会互相阻塞

## 分享会话

在会话页面点击 **Share**，创建一个私有的 GitHub Gist。

前置要求：
- 已安装 `gh`
- 已完成 `gh auth login`

分享会返回：
- 私有 gist URL
- 位于 `https://pi.dev/session/#<gistId>` 的预览 URL

分享的 gist 是快照，不会实时更新。

## 登录时自动启动

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux（systemd）

```bash
# 安装 systemd 用户服务
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# 可选：为回环以外的绑定设置你的 PI_WEB_TOKEN
#（或在 pi 内部使用 /pi-web set-token <token>）
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# 启用并启动
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# 检查状态
systemctl --user status pi-web.service

# 查看日志
journalctl --user -u pi-web.service -f
```

> 为了让服务在开机时（登录之前）启动，请改用系统服务：
> 将 `init/pi-web.service` 复制到 `/etc/systemd/system/`，并使用 `sudo systemctl`。

### Windows

安装器会自动配置这一点，无需管理员权限：
在 `HKCU\Software\Microsoft\Windows\CurrentVersion\Run` 下的 `pi-web` 条目
会在登录时启动 `~/.config/pi-web/pi-web-start.vbs`，该脚本在加载 `~/.config/pi-web/env`
（`PI_WEB_TOKEN`、`PATH` 等）之后隐藏启动二进制文件（无控制台窗口）。

要手动管理它：

```powershell
# 启动 / 停止
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# 移除自动启动
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Windows 上没有服务监控：如果 pi-web 崩溃，它会一直停在那里直到下次登录（launchd/systemd 会在其他
平台上自动重启它）。
