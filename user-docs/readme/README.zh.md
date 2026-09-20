<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · **中文** · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

在你的手机、平板或笔记本电脑上驱动你的 [pi](https://pi.dev) 编码 agent —— 网络内任意位置，或经由 Tailscale 远程访问。

它是一个完整的 PWA，因此你可以安装它并在任何设备上像原生应用一样使用。把它想象成你自己的个人 AI 工作空间 —— 类似 Claude 的 Cowork，但使用不同的模型 —— 跨模型聊天、在手机上写代码，或把它变成一个运行在你机器上的[个人助理](../en/personal-assistant.md)。

让它成为你自己的：切换主题和字体，并以你自己的语言使用它 —— pi-web 内置多种语言，你还可以添加你自己的。更多功能正在路上，但它不会变得臃肿：任何你不需要都可以在设置中关闭。

</div>

## 为什么是本地模型版？

原版 pi-web 仍然是共享功能和修复的上游基础。此版本保留那种体验，然后为你的机器或局域网其他地方运行的模型添加一层可靠性 —— 那里的生成往往更慢，内存有限，而很长的上下文可能会卡住一个否则健康的会话。

| 领域 | 上游 pi-web | 此版本 |
|------|-----------------|--------------|
| 模型/运行时策略 | 标准 pi-web 行为 | 每会话的 **Auto / Local / Cloud** 模式，带有端点感知的本地检测和持久的手动覆盖 |
| 长上下文处理 | 正常的 pi 压缩行为 | Local Mode 在 **65%** 时主动压缩，并在长工具调用循环内、下一次模型请求之前再次检查 |
| 压缩安全性 | 标准摘要 | 有界的滚动检查点，针对无效/被截断的输出进行一次更紧凑的重写，以及无进展检测而非无休止的重新压缩 |
| 中断的运行 | 常规的 worker 和错误处理 | 对上下文溢出、仅思考即停止以及特定传输中断的有界恢复，带有持久的循环中断器 |
| 手动救援 | 标准上下文详情 | **Force Compact** 仍然可用，作为一条显式的恢复路径，而不会抹掉对话 |
| 兼容性和发布 | 原始项目和发布线 | 仅本地的保护始终留在 Local Mode 之后；Cloud Mode 保留上游行为，上游更改在此独立审阅并发布 |

这不是重写，也不是上游的替代品。它是一个刻意维护的运行配置文件，面向希望获得本地模型隐私和控制力、而不愿接受脆弱长时会话的人。请参阅[用户指南](../en/README.md)了解面向用户的工作流，参阅[本地模型版开发](../../docs/dev/local-llm-development.md)了解实现和同步策略。

> [!TIP]
> 初次接触？ **[阅读用户指南 →](../en/README.md)** 了解功能的完整导览、安装步骤和技巧。([其他语言 →](../README.md))

## 截图

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>桌面端</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>移动端</em>
</div>

## 各部分如何配合

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

- **pi** 在工作时把会话 JSONL 写入 `~/.pi/agent/sessions/`。
- **pi-web** 是一个 Go 服务器，读取这些文件，在浏览器中渲染它们，并通过 SSE 推送实时更新。
- **pi --mode rpc** workers 处理由浏览器发起的聊天 —— 每个会话一个，空闲 10 分钟后被回收。
- **fsnotify** 监视会话目录，使浏览器在新输出产生的毫秒之内重新加载。
- **Tailscale Serve** 把 localhost 服务器发布为你的 tailnet 上的一个 HTTPS 端点。

## 安装

```bash
pi install npm:@timmygod/pi-web-local
```

就这样 —— 它下载匹配的二进制文件，设置自动启动，并注册 `/web`、`/pi-web`、`/remote` 和 `/refresh` 命令。

安装完成后，在浏览器中打开 `http://127.0.0.1:31415`。从 pi 中使用 `/web` 可立即在浏览器中打开当前会话。如果你的机器上正在运行 Tailscale，pi-web 会自动在你的 tailnet 上发布一个 HTTPS 端点 —— 从 pi 中使用 `/remote` 可获得任何 tailnet 上设备所用的 QR 码和 URL。

> **macOS 远程访问：** 交互式地安装并打开 Tailscale，批准管理员提示，并登录。然后运行 `/pi-web restart`，接着运行 `/remote`。

如需手动安装、二进制下载或从源码构建，请参阅 [user-docs/install.md](../en/install.md)。

## Pi 集成

执行 `pi install npm:@timmygod/pi-web-local` 之后，你会得到：

| 命令 | 作用 |
|---------|--------------|
| `/web` | 在浏览器中打开当前会话（SSH 感知：跳过浏览器，仅显示 URL） |
| `/pi-web` | 显示状态、版本，启动/停止/重启服务器，或更新 |
| `/remote` | 显示用于经由 Tailscale 远程访问的 QR 码和 URL |
| `/refresh` | 把从远程浏览器写回的新消息拉回终端会话 |

会话**自动命名**内置于 pi-web 本身，并在 `/settings` 页面进行配置。它**默认开启**，会自动为会话命名。你可以选择：

- **何时命名** —— 每个会话一次，或每条新消息都命名（默认）。
- **命名模型** —— 默认使用免费、即时的**内置词启发式（不使用 AI）**，或选择一个模型（例如一个小型/快速模型）以获得更智能、由模型撰写的标题。

该软件包还会把 pi-web 二进制安装到 `~/.pi/agent/bin/pi-web`，并设置登录时的自动启动。

## 登录时自动启动

`pi install npm:@timmygod/pi-web-local` 命令会自动完成此项设置：

| 操作系统 | 机制 |
|----|-----------|
| macOS | 位于 `~/Library/LaunchAgents/com.pi-web.plist` 的 launchd plist |
| Linux | 位于 `~/.config/systemd/user/pi-web.service` 的 systemd 用户服务 |
| Windows | `HKCU` Run-key 条目，启动 `~/.config/pi-web/` 中的隐藏启动器 |

要为远程访问设置 token，请创建 `~/.config/pi-web/env`：

```
PI_WEB_TOKEN=your-token-here
```

如需更多细节（手动设置、自定义端口、非 loopback 绑定），请参阅 [user-docs/install.md](../en/install.md)。

## 开发

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

上游同步、本地模型测试以及并行发布工作流，请参阅[本地模型版开发](../../docs/dev/local-llm-development.md)。
