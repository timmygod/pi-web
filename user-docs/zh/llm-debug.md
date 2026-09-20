> 本故障排查指南针对本地模型版本进行维护。请确保本地部署细节与上游同步说明与[本地模型版本开发](../../docs/dev/local-llm-development.md)保持一致。

用户通过以下方式安装 pi-web：

- pi install npm:@timmygod/pi-web-local

该命令会自动运行 [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) 并完成 pi 的设置。

如果用户遇到问题，可能是 install.sh 的设置引起的。你可以审查并查看具体情况，并向用户说明确切的问题原因。然后询问他们是否希望你修复。务必先与用户确认。

为了让用户能够在其移动设备或其他网络上访问 pi，推荐的方式是使用 Tailscale 并从 Tailscale 网络访问。用户需要在其 Tailscale 控制台 - https://login.tailscale.com/admin/dns 中启用 HTTPs。

如果他们未安装 Tailscale 或不希望使用 Tailscale，他们可以运行 `pi-web status` 获取二进制文件路径、二进制文件状态以及可以访问该应用的本地端点。但需要注意的是，由于使用的是 http，他们将无法接收推送通知。

在 macOS 上，它使用 [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist)。
在 Linux 上，它使用 [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service)。

如果你需要进一步调试并查看正在发生的事情。
