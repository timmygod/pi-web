<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dw/@timmygod/pi-web-local?label=downloads/wk&color=2ea043&cacheSeconds=86400)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb&cacheSeconds=86400)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · **中文** · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

从您的手机、平板或笔记本电脑上驱动 [pi](https://pi.dev) 编程助手 — 无论在您的网络内任何地方，还是通过 Tailscale 远程访问。

它是一个完整的 PWA，因此您可以安装它，像原生应用一样在任何设备上使用。把它想象成您个人的 AI 工作空间 — 类似于 Claude 的 Cowork，但可以使用不同的模型 — 跨模型对话，从手机编写代码，或将其变成一个运行在您机器上的[个人助理](../en/personal-assistant.md)。

把它变成您的专属工具：切换主题和字体，使用您自己的语言 — pi-web 自带多种语言，您还可以添加自己的语言。更多功能即将推出，但它不会变得臃肿：任何您不需要的功能都可以在设置中关闭。

</div>

## 为什么选择本地模型版本？

原始的 pi-web 仍然是共享功能和修复的上游基础。此版本保留了该体验，并为在您自己的机器或 LAN 上其他位置运行的模型添加了一层可靠性——在这些环境中，生成通常较慢，内存有限，长上下文可能会使原本健康的会话停滞。

| 领域 | 上游 pi-web | 此版本 |
|------|-----------------|--------------|
| 模型/运行时策略 | 标准 pi-web 行为 | 每个会话的 **Auto / Local / Cloud** 模式，具有端点感知的本地检测和持久的手动覆盖 |
| 长上下文处理 | 正常的 pi 压缩行为 | Local Mode 在 **65%** 时主动压缩，并在长工具调用循环中于下一次模型请求前再次检查 |
| 压缩安全性 | 标准摘要 | 有界的滚动检查点，针对无效/受限输出的一次更紧凑的重写，以及无进展检测，而不是无休止的重新压缩 |
| 中断的运行 | 正常的工作进程和错误处理 | 针对上下文溢出、仅思考停止和选定的传输中断的有界恢复，带有持久的循环断路器 |
| 手动救援 | 标准上下文详情 | **Force Compact** 仍然作为显式恢复路径可用，而不会擦除对话 |
| 兼容性和发布 | 原始项目和发布线 | 仅限本地的安全措施保留在 Local Mode 之后；Cloud Mode 保留上游行为，上游更改在此独立审查和发布 |

这不是对上游的重写或替代。这是一个刻意维护的操作配置文件，适用于那些希望拥有本地模型隐私和控制权，但又不愿接受脆弱的长时间运行会话的用户。请参阅[用户指南](../en/README.md)了解面向用户的工作流程，以及[本地模型版本开发](../../docs/dev/local-llm-development.md)了解实现和同步策略。

> [!WARNING]
> pi-web 目前处于 **beta** 阶段。一切都可能发生变化和故障！

> [!TIP]
> 新用户？**[阅读用户指南 →](../en/README.md)** 获取完整的功能介绍、安装步骤和使用技巧。（[其他语言 →](../README.md)）

## 截图

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="桌面端" width="90%" /><br />
  <em>桌面端</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="移动端 PWA" width="90%" /><br />
  <em>移动端 PWA</em>
</div>

## 架构概览

```
 pi (终端)                      浏览器 (手机 / 平板 / 笔记本电脑)
      │                                │
      │  写入 JSONL                    │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP 服务)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (每个会话          (实时重载)         (远程 HTTPS
             的聊天 worker)                        通过 MagicDNS)
```

- **pi** 在工作时将对话 JSONL 写入 `~/.pi/agent/sessions/`。
- **pi-web** 是一个 Go 服务，它读取这些文件，在浏览器中渲染它们，并通过 SSE 推送实时更新。
- **pi --mode rpc** worker 处理来自浏览器的聊天 — 每个会话一个，空闲 10 分钟后回收。
- **fsnotify** 监控会话目录，以便浏览器在新输出产生的几毫秒内重载。
- **Tailscale Serve** 将本地服务发布为您 tailnet 上的 HTTPS 端点。

## 安装

```bash
pi install npm:@timmygod/pi-web-local@beta
```

就这些 — 它会下载匹配的二进制文件，设置自动启动，并注册 `/web`、`/pi-web`、`/remote` 和 `/refresh` 命令。

安装完成后，在浏览器中打开 `http://127.0.0.1:31415`。在 pi 中，使用 `/web` 可立即在浏览器中打开当前会话。如果您的机器上运行着 Tailscale，pi-web 会自动在您的 tailnet 上发布一个 HTTPS 端点 — 在 pi 中使用 `/remote` 即可获取 QR 码和 URL，以便 tailnet 上的任何设备访问。

> **macOS 远程访问：**以交互方式安装并打开 Tailscale，批准管理员提示并登录。然后运行 `/pi-web restart`，再运行 `/remote`。

有关手动安装、二进制下载或从源码构建，请参阅 [user-docs/install.md](../en/install.md)。

## Pi 集成

运行 `pi install npm:@timmygod/pi-web-local@beta` 后，您将获得：

| 命令 | 功能描述 |
|---------|--------------|
| `/web` | 在浏览器中打开当前会话（支持 SSH 感知：跳过浏览器仅显示 URL） |
| `/pi-web` | 显示状态、版本、启动/停止/重启服务或更新 |
| `/remote` | 显示用于通过 Tailscale 远程访问的 QR 码和 URL |
| `/refresh` | 将远程浏览器写入的新消息拉回终端会话中 |

会话**自动标题**内置于 pi-web 中，可在 `/settings` 页面配置。它**默认开启**，会自动为会话命名。您可以选择：

- **何时生成标题** — 每个会话仅一次，或在每条新消息时生成（默认）。
- **标题模型** — 默认使用免费、即时的**内置词语启发式算法（无 AI）**，或选择一个模型（例如小型/快速的模型）来获得更智能的、由模型编写的标题。

该包还会将 pi-web 二进制文件安装到 `~/.pi/agent/bin/pi-web`，并设置登录时自动启动。

## 登录时自动启动

`pi install npm:@timmygod/pi-web-local@beta` 命令会自动完成此设置：

| 操作系统 | 机制 |
|----|-----------|
| macOS | launchd plist，位于 `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd 用户服务，位于 `~/.config/systemd/user/pi-web.service` |

要设置远程访问令牌，请创建 `~/.config/pi-web/env`：

```
PI_WEB_TOKEN=your-token-here
```

更多细节（手动设置、自定义端口、非环回绑定），请参阅 [user-docs/install.md](../en/install.md)。

## 开发

```bash
make setup   # 安装前端依赖并下载 Go 模块
make check   # 前端测试/构建 + Go 测试/审查
make build   # 如有需要则先 setup，构建前端，然后构建 ./pi-web
```
