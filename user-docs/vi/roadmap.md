# Roadmap

Roadmap này thuộc về phiên bản local-model của pi-web. Các tính năng upstream được đồng bộ định kỳ, trong khi công việc triển khai và độ tin cậy cục bộ được xác nhận và phát hành trên dòng này; xem [Phát triển phiên bản local-model](../../docs/dev/local-llm-development.md).

pi-web được xây dựng cho hai đối tượng:

- **Đối với developer** — những người sống trong terminal nhưng muốn tiếp tục phiên làm việc từ điện thoại, chuyển giao sang remote server, hoặc theo dõi các tác vụ chạy dài từ bất cứ đâu.
- **Đối với người không phải developer** — những người chỉ muốn một ứng dụng AI đẹp mắt và hoạt động trơn tru. Mở lên, gõ, và vibe. Không cần terminal, không cần SSH, không cần bối rối. Giống như các công cụ AI thân thiện nhất, nhưng với quyền chọn model và tự do mã nguồn mở.

Dưới đây là những gì sắp đến.

Phiên bản này theo dõi pi-web upstream trên một dòng phát hành riêng. Các tính năng upstream được nhập định kỳ; công việc độ tin cậy local-model được ưu tiên và xác nhận tại đây mà không thay đổi lịch sử phát hành upstream.

---

## Hiện tại (đã phát hành)

Mọi thứ được liệt kê trong [bảng tính năng](README.md#what-you-can-do-with-pi-web) đều đang hoạt động ngay hôm nay.

---

## Tiếp theo

| # | Tính năng | Nó làm gì |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram & Discord bots** | Chat với pi qua Telegram hoặc Discord — hoàn hảo cho các workflow trợ lý cá nhân khi di chuyển. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | Theo dõi token, ước tính chi phí, phân tích phiên — biết bạn đang sử dụng pi như thế nào. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | Đặt mức hiển thị mong muốn cho thinking, tools và kết quả tool trên tất cả các phiên. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | Gửi hướng dẫn bổ sung trong khi pi vẫn đang chạy — điều hướng nó giữa chừng. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` command** | Nén các cuộc trò chuyện dài ngay từ web UI, không cần terminal. |

---

## Đang lên kế hoạch

| # | Tính năng | Nó làm gì |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | Duyệt cây file của dự án và xem thay đổi git trực tiếp trong pi-web. Chọn tham gia (opt-in) để không làm phiền bạn. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | Lên lịch chạy prompt tự động — standup hằng ngày, tóm tắt buổi sáng, tác vụ lặp lại. Có kiểm soát bởi admin vì lý do an toàn. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Customizable shortcuts** | Gán lại mọi phím tắt để phù hợp với trí nhớ vận động của bạn. |

---

## Tầm nhìn

Mục tiêu dài hạn: pi-web nên là **giao diện cho pi** — dành cho mọi người.

- **Người không phải dev** mở nó như bất kỳ ứng dụng nào khác. Chọn một model. Gõ. Xong. Không bao giờ cần command line.
- **Dev** được tích hợp sâu — chuyển giao remote, dashboard đa phiên, duyệt theo nhận thức git, bot nhắn tin.
- **Mọi người** được tự do model, minh bạch mã nguồn mở, và một UI cảm giác chu đáo ở mọi khía cạnh.

---

> 💡 Có một ý tưởng? [Mở một issue](https://github.com/timmygod/pi-web/issues/new) hoặc tham gia thảo luận.
