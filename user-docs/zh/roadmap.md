# 路线图

本路线图属于 pi-web 的本地模型版本。上游功能会定期同步，而本地部署与可靠性相关工作则在该分支上进行验证和发布；参见[本地模型版本开发](../../docs/dev/local-llm-development.md)。

pi-web 面向两类用户：

- **面向开发者** —— 他们日常生活在终端中，但希望从移动端继续会话、交接至远程服务器，或随时随地关注长时间运行的任务。
- **面向非开发者** —— 他们只想要一个开箱即用、美观的 AI 应用。打开、输入、开聊。无需终端，无需 SSH，无需困惑。就像最易用的 AI 工具一样，但拥有模型选择权和开源的自由。

以下是即将推出的功能。

该版本在独立的发布线上跟踪上游 pi-web。上游功能会定期导入；本地模型的可靠性工作会在此处被优先处理并验证，且不会改变上游的发布历史。

---

## 当前（已发布）

[功能表格](README.md#what-you-can-do-with-pi-web)中列出的所有内容现已全部上线。

---

## 接下来

| # | 功能 | 作用 |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram 和 Discord 机器人** | 通过 Telegram 或 Discord 与 pi 聊天 —— 非常适合移动办公中的个人助理工作流。 |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **用量洞察** | Token 追踪、成本估算、会话分析 —— 了解你如何使用 pi。 |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **可配置的默认值** | 为所有会话中思考、工具及工具输出设置你偏好的可见性。 |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **引导 / 队列** | 在 pi 仍在运行时发送后续指令 —— 在过程中引导它。 |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` 命令** | 直接在 Web UI 中压缩长对话，无需终端。 |

---

## 计划中

| # | 功能 | 作用 |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **文件浏览器与 Git Diff** | 在 pi-web 中浏览项目文件树，并直接查看 git 更改。可选启用，不会打扰你的工作。 |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **调度器** | 定时自动运行提示词 —— 每日站会、晨报摘要、周期性任务。为安全起见受管理员管控。 |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **可自定义快捷键** | 重新映射每个键盘快捷键以匹配你的肌肉记忆。 |

---

## 愿景

长期目标：pi-web 应当成为 **pi 的界面** —— 面向所有人。

- **非开发者**像使用任何其他应用一样打开它。选择一个模型。输入。完成。从不需要命令行。
- **开发者**获得深度集成 —— 远程交接、多会话仪表盘、支持 git 的浏览、消息机器人。
- **所有人**都获得模型自由、开源透明，以及处处彰显细致体贴的 UI。

---

> 💡 有个想法？[提交一个 issue](https://github.com/timmygod/pi-web/issues/new) 或加入讨论。
