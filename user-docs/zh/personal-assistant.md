# 将 pi-web 作为你的个人助理

此工作流受本地模型版本支持。关于该版本的开发、同步和发布策略，请参阅[本地模型版本开发](../../docs/dev/local-llm-development.md)。

pi-web 不仅仅用于编程——你可以将它变成一个**个人 AI 助理**，驻留在你的电脑上，就像拥有你自己的 OpenClaw 或 Hermes 一样。

## 工作原理

你在机器上创建一个专用文件夹——你的助理就住在那里。在文件夹内，放入一个 `APPEND_SYSTEM.md` 文件，用来定义你的助理是谁、它知道什么，以及它如何行事。pi-web 为你提供了一个精美的聊天界面，让你可以从任何设备与它对话。

## 分步指南

### 1. 创建你的助理文件夹

在你的电脑上选择一个文件夹。例如：

```
~/my-assistant/
```

### 2. 定义你的助理

在该文件夹内创建一个 `APPEND_SYSTEM.md` 文件。这里就是你告诉 pi 你的助理是谁的地方：

```markdown
# My Personal Assistant

You are Jarvis, my personal AI assistant. You help me with:

- Daily planning and reminders
- Research and summarization
- Drafting emails and messages
- Brainstorming ideas
- Keeping track of things I mention

## About me

- I'm a software engineer who works remotely
- I have a cat named Pixel
- I prefer short, direct answers
- My timezone is PST

## Rules

- Be concise — I value brevity
- If you don't know something, say so
- Proactively remind me of things I asked you to track
```

pi 会自动将此内容附加到每个对话的系统提示中，因此你的助理始终知道你是谁，以及如何帮助你。

### 3. 在该文件夹中启动一个会话

在 pi-web 中，创建一个新的会话并指向 `~/my-assistant/`（或者你给它起的任何名字）。就这样——你正在与你的个人助理对话。

### 4. 从任何地方使用它

在你的手机、平板电脑或笔记本电脑上将 pi-web 安装为 PWA。你的助理始终在那里——随时问它任何问题。

## 关于你的助理的创意

| 角色 | 在 APPEND_SYSTEM.md 中写什么 |
|---|---|
| 🧠 **生活教练** | 你的目标、正在培养的习惯、日记提示 |
| 🏠 **家庭管理者** | 购物清单格式、家庭成员的偏好、餐食计划 |
| 💼 **工作伙伴** | 你的角色、当前项目、会议记录格式、公司背景 |
| 📚 **学习伙伴** | 你在学什么、偏好的讲解风格、测验我模式 |
| ✍️ **写作助理** | 你的写作风格、语气偏好、你常用的格式 |

## 添加更多上下文

你可以在助理文件夹中放入任何有助于 pi 更实用的内容：

- `notes/` — 你的助理可以阅读的参考文件
- `context.md` — 关于你生活或工作的背景信息
- `projects.md` — 当前项目及其状态

pi 可以读取文件夹中的文件，因此你提供的上下文越多，它就变得越好。

## 让 pi-web 做事

在运行 `pi install npm:@timmygod/pi-web-local` 之后，会话可以与 pi-web 本身进行交互。
试试：

- "添加一个在新加坡时间凌晨 2 点总结我收件箱的计划"
- "列出我的 pi-web 计划"
- "暂停收件箱计划"
- "把这件事写进笔记里"
- "将 pi-web 切换到深色模式 / 关闭自动标题"

捆绑的 **/skill:pi-web-schedule** 技能会将其变成一个真正的 pi-web 计划（与你编辑的 `/schedules` 中的相同）。每次触发都会启动一个**新的**会话，因此指令必须能独立成立——"总结 ~/inbox 中的未读邮件" 可以；"继续我们之前在做的事" 则不行。

计划只在 pi-web 运行时才会执行。

---

> 💡 **提示：** 从简单开始。只写几行关于你是谁以及你希望助理如何行事的内容。随着你逐渐了解哪些方法有效，再随时间迭代改进。
