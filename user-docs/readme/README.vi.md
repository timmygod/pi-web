<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · **Tiếng Việt** · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Điều khiển [pi](https://pi.dev) coding agent của bạn từ điện thoại, tablet hoặc laptop — ở bất cứ đâu trong mạng của bạn, hoặc từ xa qua Tailscale.

Đây là một PWA hoàn chỉnh, nên bạn có thể cài đặt và sử dụng nó như một ứng dụng gốc trên bất kỳ thiết bị nào. Hãy nghĩ về nó như không gian làm việc AI cá nhân của riêng bạn — giống như Cowork của Claude, nhưng với các model khác — trò chuyện chéo giữa các model, viết code từ điện thoại, hoặc biến nó thành một [trợ lý cá nhân](../en/personal-assistant.md) chạy trên máy của bạn.

Làm cho nó mang dấu ấn của riêng bạn: đổi theme và font, và sử dụng nó bằng ngôn ngữ của bạn — pi-web đi kèm với nhiều ngôn ngữ và bạn có thể thêm ngôn ngữ của mình. Nhiều tính năng hơn đang trên đường đến, nhưng nó sẽ không trở nên bloated: bất cứ thứ gì bạn không cần đều có thể tắt trong cài đặt.

</div>

## Why this local-model edition?

Phiên bản pi-web gốc vẫn là nền tảng upstream cho các tính năng và bản sửa lỗi được chia sẻ. Phiên bản này giữ lại trải nghiệm đó, sau đó thêm một lớp độ tin cậy cho các model chạy trên máy của bạn hoặc ở nơi khác trong LAN của bạn — nơi mà việc sinh nội dung thường chậm hơn, bộ nhớ có hạn, và một context dài có thể làm treo một phiên hoạt động bình thường khác.

| Area | Upstream pi-web | This edition |
|------|-----------------|--------------|
| Model/runtime policy | Hành vi chuẩn của pi-web | Chế độ **Auto / Local / Cloud** theo từng phiên, với phát hiện local nhận diện endpoint và một chế độ ghi đè thủ công được lưu bền vững |
| Long-context handling | Hành vi nén chuẩn của pi | Chế độ Local chủ động nén ở **65%** và kiểm tra lại bên trong các vòng lặp gọi tool dài trước yêu cầu model tiếp theo |
| Compaction safety | Tóm tắt chuẩn | Các checkpoint rolling có giới hạn, một lần viết lại chặt chẽ hơn cho đầu ra không hợp lệ/vượt giới hạn, và phát hiện không tiến triển thay vì tái nén vô hạn |
| Interrupted runs | Xử lý worker và lỗi chuẩn | Khôi phục có giới hạn cho tràn context, dừng chỉ-thinking, và các ngắt kết nối transport được chọn, với các bộ phá vòng lặp lưu bền vững |
| Manual rescue | Chi tiết context chuẩn | **Force Compact** vẫn khả dụng như một đường dẫn khôi phục rõ ràng mà không xóa cuộc hội thoại |
| Compatibility and releases | Dự án và dòng phát hành gốc | Các biện pháp bảo vệ chỉ-local nằm sau chế độ Local; chế độ Cloud giữ nguyên hành vi upstream, và các thay đổi upstream được xem xét và phát hành tại đây độc lập |

Đây không phải là viết lại hay thay thế cho upstream. Đây là một hồ sơ vận hành được duy trì một cách có chủ đích cho những người muốn quyền riêng tư và quyền kiểm soát model-local mà không chấp nhận các phiên chạy dài kém ổn định. Xem
[hướng dẫn người dùng](../en/README.md) cho luồng công việc hướng tới người dùng và
[phát triển phiên bản local-model](../../docs/dev/local-llm-development.md) cho phần triển khai và chính sách đồng bộ hóa.

> [!TIP]
> Mới đến đây? **[Đọc hướng dẫn người dùng →](../en/README.md)** để có một tour đầy đủ về tính năng, các bước cài đặt và mẹo. ([Các ngôn ngữ khác →](../README.md))

## Screenshots

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## How It Fits Together

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

- **pi** ghi JSONL của cuộc hội thoại vào `~/.pi/agent/sessions/` khi nó làm việc.
- **pi-web** là một server Go đọc các tệp đó, hiển thị chúng trong trình duyệt, và stream các cập nhật trực tiếp qua SSE.
- Các worker **pi --mode rpc** xử lý chat khởi tạo từ trình duyệt — một worker cho mỗi phiên, được thu dọn sau 10 phút inactive.
- **fsnotify** theo dõi thư mục sessions để trình duyệt reload lại trong vòng vài mili giây sau đầu ra mới.
- **Tailscale Serve** xuất bản server localhost như một endpoint HTTPS trên tailnet của bạn.

## Install

```bash
pi install npm:@timmygod/pi-web-local
```

Chỉ có vậy — nó tải về binary tương ứng, thiết lập auto-start, và đăng ký các lệnh `/web`, `/pi-web`, `/remote`, và `/refresh`.

Sau khi cài đặt, mở `http://127.0.0.1:31415` trong trình duyệt của bạn. Từ pi, dùng `/web` để mở phiên hiện tại trong trình duyệt ngay lập tức. Nếu Tailscale đang chạy trên máy của bạn, pi-web tự động xuất bản một endpoint HTTPS trên tailnet của bạn — dùng `/remote` từ pi để lấy mã QR và URL cho bất kỳ thiết bị nào trên tailnet của bạn.

> **Truy cập từ xa trên macOS:** Cài đặt và mở Tailscale tương tác, chấp nhận lời nhắc quản trị viên, và đăng nhập. Sau đó chạy `/pi-web restart`, rồi chạy `/remote`.

Đối với cài đặt thủ công, tải binary, hoặc build từ source, xem [user-docs/install.md](../en/install.md).

## Pi Integration

Sau `pi install npm:@timmygod/pi-web-local`, bạn sẽ có:

| Command | What it does |
|---------|--------------|
| `/web` | Mở phiên hiện tại trong trình duyệt (nhận biết SSH: bỏ qua trình duyệt và chỉ hiển thị URL) |
| `/pi-web` | Hiển thị trạng thái, phiên bản, khởi động/dừng/khởi động lại server, hoặc cập nhật |
| `/remote` | Hiển thị mã QR và URL để truy cập từ xa qua Tailscale |
| `/refresh` | Kéo các tin nhắn mới được ghi từ các trình duyệt từ xa trở lại phiên terminal |

**Đặt tiêu đề tự động** cho phiên được tích hợp sẵn trong chính pi-web và được cấu hình trên trang `/settings`. Nó **bật theo mặc định** và tự động đặt tên cho các phiên. Bạn có thể chọn:

- **Khi nào đặt tiêu đề** — một lần mỗi phiên, hoặc với mỗi tin nhắn mới (mặc định).
- **Model đặt tiêu đề** — theo mặc định là **heuristic từ tích hợp, miễn phí và tức thời (không dùng AI)**, hoặc chọn một model (ví dụ một model nhỏ/nhanh) để có tiêu đề do model viết thông minh hơn.

Gói này cũng cài đặt binary pi-web vào `~/.pi/agent/bin/pi-web` và thiết lập auto-start khi đăng nhập.

## Auto-Start on Login

Lệnh `pi install npm:@timmygod/pi-web-local` thiết lập điều này tự động:

| OS | Mechanism |
|----|-----------|
| macOS | launchd plist tại `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service tại `~/.config/systemd/user/pi-web.service` |
| Windows | mục Run-key `HKCU` khởi động một trình khởi động ẩn trong `~/.config/pi-web/` |

Để đặt một token cho truy cập từ xa, tạo `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

Để biết thêm chi tiết (thiết lập thủ công, custom ports, non-loopback binds), xem [user-docs/install.md](../en/install.md).

## Development

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Đối với đồng bộ hóa upstream, kiểm thử local-model, và luồng phát hành song song, xem [Local-model edition development](../../docs/dev/local-llm-development.md).
