# Cài đặt & Sử dụng

## Tính năng

### Điều khiển từ xa

- Tiếp tục bất kỳ phiên nào từ trình duyệt, kèm theo đính kèm văn bản hoặc hình ảnh
- Khởi tạo một phiên hoàn toàn mới cho bất kỳ đường dẫn dự án nào, ngay từ giao diện web
- Chuyển đổi model và bộ chọn mức độ suy luận trong trình duyệt, theo từng phiên
- Trạng thái worker theo từng phiên (đỡ / đang chạy / lỗi) với tự khôi phục khi crash
- Nhiều phiên chạy song song — khởi động công việc ở phiên này, xem phiên khác stream
- `PI_WEB_TOKEN` cho phơi bày LAN an toàn — bắt buộc mặc định cho bất kỳ việc bind không phải loopback nào

### Đọc phiên

- Duyệt các phiên xuyên suốt các dự án với bộ lọc, tìm kiếm, và điều hướng branch đầy đủ
- Cập nhật tăng phân trực tiếp khi pi vẫn đang chạy (qua fsnotify; độ trễ ~ms)
- Chế độ Follow để tail các phiên đang hoạt động
- Deep links đến từng tin nhắn
- Tải xuống một phiên dưới dạng JSONL
- Chia sẻ các bản chụp tĩnh dưới dạng GitHub Gist bí mật
- Các extension pi: `/web`, `/remote`, `/refresh`, `/pi-web token` và `/pi-web set-token` để mở phiên, QR từ xa, đồng bộ phiên, và quản lý token
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) để một phiên có thể quản lý lịch trình, bản nháp (scratchpad) của dự án và cài đặt bằng ngôn ngữ tự nhiên

## Chọn chế độ phiên

Phiên bản này sử dụng các provider và model đã được cấu hình trong pi; Local Mode
là một chính sách runtime, không phải trình cài đặt model riêng hay một màn hình API-key thứ hai.
Chọn một chế độ khi tạo phiên, hoặc thay đổi sau khi lượt chạy hiện tại ổn định:

| Chế độ | Dùng khi | Hành vi |
|------|-------------|----------|
| **Auto** | Bạn muốn pi-web tự quyết định | Giải quyết các endpoint local/LAN từ metadata của provider khi có thể; nếu không thì giữ đường dẫn bình thường |
| **Local** | Model đang chạy trên máy này hoặc LAN của bạn | Bật ranh giới nén (compaction) 65%, checkpoint có giới hạn, Force Compact, và khôi phục tự động có bảo vệ |
| **Cloud** | Model được chọn là dạng hosted và nên tuân theo hành vi upstream | Giữ chính sách nén và khôi phục chỉ-địa-phương nằm ngoài phiên |

Việc chọn thủ công Local hoặc Cloud sẽ thắng so với phát hiện tự động và được giữ lại qua
các lần tải lại và khởi động lại. Một phiên đang chạy sẽ từ chối thay đổi chế độ cho đến khi worker
của nó ổn định, nên chế độ hiển thị trong UI luôn khớp với chính sách thực sự đang được dùng.

## Yêu cầu

- [Go](https://go.dev) 1.25+ (chỉ để build từ source)
- `pi` trên `PATH` của bạn để chat/đổi model trong trình duyệt
- Tùy chọn: `gh` để chia sẻ
- Trên Windows: pi cần một shell bash cho công cụ shell của nó — [Git for Windows](https://git-scm.com/download/win) là đủ (xem tài liệu Windows của pi)

## Cài đặt

### Gói Pi (khuyến nghị)

```bash
pi install npm:@timmygod/pi-web-local
```

Một lệnh duy nhất này:
- Cài gói npm pi vào thư mục gói của pi
- Chạy script `postinstall` của gói (`install.sh`, hoặc `install.ps1` trên Windows)
- Tải binary pi-web tương ứng với phiên bản gói và nền tảng của bạn từ GitHub Releases
- Cài đặt nó vào `~/.pi/agent/bin/pi-web` (`pi-web.exe` trên Windows)
- Thiết lập tự động khởi chạy khi đăng nhập (launchd trên macOS, systemd trên Linux, trình khởi chạy Run-key trên Windows)
- Đăng ký các lệnh pi: `/web`, `/remote`, `/refresh`, `/pi-web token`, và `/pi-web set-token`

Chức năng tự động đặt tiêu đề phiên được tích hợp sẵn trong pi-web (không nằm trong extension) và được cấu hình trên trang `/settings`. Bật mặc định: pi-web tự động đặt tên phiên bằng một heuristic từ vựng tích hợp sẵn (không dùng AI), đặt lại tên ở mỗi tin nhắn mới. Bạn có thể chuyển sang chỉ đặt tiêu đề một lần mỗi phiên, và/hoặc chọn một model để viết các tiêu đề thông minh hơn thay vì heuristic.

Trên Linux, tự động khởi chạy được cấu hình như một dịch vụ systemd người dùng tại `~/.config/systemd/user/pi-web.service`. Trình cài đặt sẽ ghi đè lại `ExecStart` của nó thành đường dẫn binary đã cài thực sự. Nếu Tailscale khả dụng ở runtime, pi-web sẽ công bố (publish) máy chủ localhost với Tailscale Serve HTTPS. Nếu user systemd không khả dụng, hãy chạy thủ công với `~/.pi/agent/bin/pi-web -o`.

Để chỉ cài cho một dự án cụ thể (chia sẻ với đội ngũ của bạn qua `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

Sau đó khởi động lại pi (hoặc chạy `/reload`), và sử dụng `/web`, `/pi-web`, `/remote`, `/refresh`. Quản lý token truy cập của bạn với `/pi-web token` và `/pi-web set-token`.

Nếu npm bị dừng với `ENOTEMPTY` trong lúc đổi tên `@timmygod/pi-web-local`, hãy xóa các thư mục backup ẩn đã lỗi thời của npm và cài lại gói:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Cài nhanh (không cần công cụ build)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Điều này tải binary pi-web mới nhất, cài đặt nó vào `/usr/local/bin` (`~/.pi/agent/bin` trên Windows), và thiết lập tự động khởi chạy khi đăng nhập. Không cần Go, Node, hay pi.

### Tải binary

Các binary đã build sẵn được đính kèm với mỗi [GitHub Release](https://github.com/timmygod/pi-web/releases).

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

Sau đó chuyển nó vào PATH của bạn:

```bash
cp pi-web ~/.pi/agent/bin/
# hoặc toàn hệ thống:
sudo cp pi-web /usr/local/bin/
```

### Build từ source

Cái checkout này là phiên bản local-model của pi-web. Quy trình build bình thường tạo ra
ứng dụng web và backend cùng lúc; các biện pháp bảo vệ local-model được kích hoạt
ở runtime bởi Local Mode hiệu lực của phiên, chứ không phải bằng một binary riêng.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # builds the Vite bundle, then embeds it into the Go binary

# optional: put it on PATH
cp pi-web ~/.pi/agent/bin/
```

Frontend bundle được nhúng (embed) bởi `web/assets_embed.go`, nên `go build` cần
`web/dist` phải tồn tại trước. `make build` thực hiện cả hai bước theo thứ tự; nếu bạn build
thủ công, hãy chạy `npm --prefix web install && npm --prefix web run build` trước
`go build ./cmd/pi-web`.

Để xem quy trình fork được bảo trì, đồng bộ upstream, và danh sách kiểm tra xác minh Local Mode, xem [ghi chú phát triển local-model](../../docs/dev/local-llm-development.md).

### Phát triển song song với một instance đã cài

Giữ instance đã cài đang chạy trên cổng `31415`, sau đó khởi chạy cái
checkout source ở chế độ phát triển:

```bash
make dev
```

Mở `http://127.0.0.1:31416`. `make dev` đặt môi trường phát triển nội bộ `PI_WEB_DEV=1`,
nên cái checkout source sẽ dùng chung các phiên, cài đặt, và dữ liệu
SQLite với instance đã cài trong khi giữ một khóa runtime phát triển và
tệp state riêng. Các instance đã cài thông thường và được khởi chạy thủ công
không bị thay đổi và giữ nguyên hành vi single-instance ban đầu.

Để ngăn chặn công việc tự động trùng lặp, chế độ phát triển không chạy
schedule loop, chat-queue drainer, tự động đặt tiêu đề, hay thông báo đẩy. Các
yêu cầu trực tiếp được thực hiện qua UI phát triển vẫn hoạt động. Đừng điều khiển cùng
một phiên chat từ cả hai instance cùng một lúc; mỗi tiến trình có trình
quản lý RPC worker riêng.

`make dev` yêu cầu [Air](https://github.com/air-verse/air) để hot reload Go:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` là cơ chế chạy của môi trường phát triển, không phải chế độ
multi-instance production được hỗ trợ.

## Gỡ cài đặt

```bash
pi remove npm:@timmygod/pi-web-local
```

Điều này chạy script `preuninstall` của gói (`uninstall.sh`, hoặc `uninstall.ps1`
trên Windows), để dừng instance đang chạy và xóa:

- binary pi-web (`~/.pi/agent/bin/pi-web`, hoặc `/usr/local/bin/pi-web` đối với cài đặt standalone)
- tệp phiên bản (`~/.pi/agent/pi-web-version`)
- tệp state runtime (`~/.pi/agent/pi-web/pi-web-state.json`)
- cấu hình tự động khởi chạy (plist launchd trên macOS, dịch vụ systemd người dùng trên Linux, mục Run-key + các script khởi chạy trên Windows)

Dữ liệu của bạn được giữ lại để một lần cài lại sau đó sẽ tiếp tục từ chỗ bạn đã dừng:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, các tệp
phiên của bạn dưới `~/.pi/agent/sessions/`, và `~/.config/pi-web/env` (bao gồm
`PI_WEB_TOKEN`). Hãy tự xóa chúng nếu bạn muốn một trang giấy sạch.

## Sử dụng

```bash
# Start on the default port (31415)
pi-web

# Start and open a browser
pi-web -o

# Custom port
pi-web -p 8080

# Override bind host (loopback is unauthenticated by default)
pi-web --host 127.0.0.1

# Non-loopback bind requires a token — pi-web refuses to start otherwise
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Mặc định, pi-web bind đến `127.0.0.1`. Nếu Tailscale đang chạy với MagicDNS **và `PI_WEB_TOKEN` đã được đặt**, pi-web cũng sẽ chạy `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` và in ra URL tailnet HTTPS. Nếu không có token, pi-web chỉ dùng loopback và bỏ qua Tailscale Serve, nên các peer trong tailnet không thể tiếp cận agent mà không xác thực. Mọi việc bind không phải loopback rõ ràng cũng yêu cầu `PI_WEB_TOKEN` phải được đặt; truyền `--insecure` để ghi đè cho việc test nội bộ.

## Truy cập từ xa

Giữ pi-web đang lắng nghe cục bộ, sau đó dùng URL Tailscale HTTPS được in ra từ điện thoại hoặc laptop của bạn trên tailnet.

Trên macOS, cài đặt và mở Tailscale tương tác, chấp nhận prompt quản trị viên, và đăng nhập. Sau đó chạy `/pi-web restart`, tiếp theo là `/remote`.

Trên Linux, cho phép người dùng của bạn quản lý Tailscale trước khi cài/chạy pi-web, nếu không `tailscale serve` có thể yêu cầu sudo và tự động khởi chạy có thể thất bại:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Start pi-web with a token so it publishes the Tailscale HTTPS endpoint
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. From any other Tailscale-connected device, open the printed
#    "Tailscale HTTPS" URL and enter the token once.
```

> Mặc định, pi-web từ chối bind đến một địa chỉ không phải loopback trừ khi `PI_WEB_TOKEN` đã được đặt — nếu không, bất kỳ ai có thể tiếp cận địa chỉ đã bind đều có thể xem các phiên và gửi chỉ thị cho pi. Để ghi đè cơ chế bảo vệ này cho việc test mạng nội bộ, hãy truyền `--insecure`. **Đừng dùng `--insecure` trên Tailscale hay bất kỳ địa chỉ nào có thể tiếp cận từ bên ngoài máy của bạn.**
>
> Các client có thể truyền token qua header `Authorization: Bearer <token>`, header `X-Pi-Token`, hoặc một lần qua `?token=<token>` (lần này sẽ đặt cookie `pi_token` cho các yêu cầu sau). Token truyền qua `?token=` sẽ kết thúc trong lịch sử trình duyệt, server access logs, và các header `Referer` từ bất kỳ link nào trên trang — hãy dùng dạng header cho mọi thứ ngoài bookmark ban đầu.

## Chat trong trình duyệt

Mở trang phiên và dùng composer ở dưới cùng để tiếp tục đúng phiên đó.

- `Enter` để gửi, `Shift+Enter` để chèn dòng mới
- Kéo-thả hoặc dán hình ảnh trực tiếp vào composer
- Bộ chọn model và bộ chọn mức độ suy luận nằm trong header — các thay đổi áp dụng ngay lập tức lên worker pi nền tảng
- Mỗi phiên đang hoạt động có một worker `pi --mode rpc` chuyên dụng riêng, nên các phiên khác nhau không chặn nhau

## Chia sẻ phiên

Nhấp vào **Chia sẻ** trên trang phiên để tạo một GitHub Gist bí mật.

Yêu cầu:
- `gh` đã được cài
- `gh auth login` đã hoàn tất

Chia sẻ trả về:
- URL gist bí mật
- một URL xem trước tại `https://pi.dev/session/#<gistId>`

Các gist được chia sẻ là bản chụp (snapshot) và không tự cập nhật trực tiếp.

## Tự động khởi chạy khi đăng nhập

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Install the systemd user service
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Optional: set your PI_WEB_TOKEN for non-loopback binds
# (or use /pi-web set-token <token> from inside pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Enable and start
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Check status
systemctl --user status pi-web.service

# View logs
journalctl --user -u pi-web.service -f
```

> Để dịch vụ khởi động khi bật máy (trước khi đăng nhập), hãy dùng một dịch vụ hệ thống thay vì:
> sao chép `init/pi-web.service` vào `/etc/systemd/system/` và dùng `sudo systemctl`.

### Windows

Trình cài đặt tự động cấu hình điều này, mà không cần quyền admin: một
mục `pi-web` dưới `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
sẽ khởi chạy `~/.config/pi-web/pi-web-start.vbs` khi đăng nhập, sau đó sẽ khởi chạy binary
ở chế độ ẩn (không có cửa sổ console) sau khi nạp `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...).

Để tự quản lý:

```powershell
# Start / stop
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Remove auto-start
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Không có giám sát dịch vụ trên Windows: nếu pi-web crash, nó sẽ giữ trạng thái tắt
cho đến lần đăng nhập tiếp theo (launchd/systemd tự động khởi chạy lại nó trên các
nền tảng khác).
