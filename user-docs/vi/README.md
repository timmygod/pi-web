# Chào mừng đến với pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · **Tiếng Việt** · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**Đang cân nhắc thử pi-web? Cứ thử đi — bạn sẽ yêu thích nó.**

pi-web là giao diện web và PWA đẹp mắt dành cho [pi](https://pi.dev) — agent mã nguồn mở về AI coding. Nó cho phép bạn duyệt, đọc và tiếp tục các phiên pi của mình từ bất kỳ trình duyệt nào, trên bất kỳ thiết bị nào, với những tính năng đầy tâm huyết ở từng bước.

## Điều gì khác biệt ở phiên bản này?

Kho mã này giữ lại giao diện và các tính năng dùng chung của pi-web phía upstream, nhưng thay đổi cách bảo vệ các phiên khi model được chọn chạy cục bộ hoặc trên LAN của bạn.

- **Chọn chính sách runtime cho từng phiên.** Tự động phát hiện các endpoint cục bộ/LAN khi metadata của provider rõ ràng; Local và Cloud là các ghi đè thủ công được lưu trữ.
- **Ngăn chặn các lỗi context sớm.** Local Mode thực hiện compaction khi đạt 65% mức sử dụng và kiểm tra lại giữa các cuộc gọi tool, trước yêu cầu model tiếp theo.
- **Giữ các tóm tắt nằm trong giới hạn.** Các checkpoint luân phiên tránh việc một tóm tắt cũ lớn lên vô hạn, thử lại một lần với ngân sách chặt hơn, và dừng một cách an toàn khi compaction không mang lại tiến bộ ý nghĩa.
- **Khôi phục một cách thận trọng.** Lỗi tràn context, các gián đoạn của transport được chọn, và các dừng sớm chỉ liên quan đến reasoning có thể tự động tiếp tục, nhưng deduplication sự cố và circuit breaker nhận biết tiến độ sẽ ngăn chặn các vòng lặp khôi phục.
- **Giữ cho người dùng quyền kiểm soát.** Force Compact luôn là đường lối cứu trợ thủ công rõ ràng, trong khi Cloud Mode giữ lại quy trình và điều khiển từ upstream.

Kết quả thực tiễn rất đơn giản: một tác vụ dài với model cục bộ nên được compaction trước khi sụp đổ, khôi phục một lần khi việc khôi phục là an toàn, và dừng gọn gàng thay vì lặp đi lặp lại khi nó không an toàn.

**pi-web được xây dựng dành cho hai loại người:**

- 🧑‍💻 **Dành cho lập trình viên** — những người sống trong terminal nhưng muốn tiếp tục phiên từ điện thoại, chuyển sang một máy chủ từ xa, hoặc theo dõi các tác vụ dài từ bất kỳ đâu.
- ✨ **Dành cho người không phải lập trình viên** — những người chỉ muốn một ứng dụng AI đẹp và hoạt động tốt. Mở ra, gõ, vào lúc vibe. Không có terminal, không có SSH, không có sự bối rối. Giống như các công cụ AI thân thiện nhất, nhưng với sự lựa chọn model và sự tự do mã nguồn mở.

---

## Tại sao là pi-web?

Bạn đã đang đắm chìm trong luồng tập trung với pi trong terminal của bạn. pi-web giữ cho động lượng đó tiếp tục khi bạn rời khỏi bàn làm việc:

- **Tiếp tục từ mọi nơi** — tiếp tục một phiên từ điện thoại, máy tính bảng, hoặc máy tính khác của bạn. Không cần SSH, không cần Termius — chỉ cần mở trình duyệt.
- **Bảng điều khiển đa phiên** — bắt đầu công việc trong một phiên trong khi theo dõi một phiên khác đang stream. Tìm kiếm xuyên suốt các dự án, lọc theo branch, tìm điều bạn cần nhanh chóng.
- **Nền tảng mã nguồn mở** — pi hoàn toàn mã nguồn mở và không phụ thuộc provider. Bạn không bị khóa vào một model hay vendor duy nhất. pi-web cũng là mã nguồn mở.
- **Truy cập từ xa an toàn** — xác thực token tích hợp sẵn để bạn có thể phơi bày nó trên LAN hoặc Tailscale của mình mà không lo lắng.
- **Chia sẻ công việc của bạn** — xuất các phiên dưới dạng các snapshot tĩnh hoặc GitHub Gist bí mật chỉ với một cú nhấp chuột.

> Tò mò về câu chuyện đằng sau? [Đọc tại sao chúng tôi xây dựng nó →](why.md)

---

## pi-web như không gian làm việc AI cá nhân của bạn 🏠

pi-web là một PWA (Progressive Web App), vì vậy bạn có thể **cài đặt nó như một ứng dụng native** trên máy để bàn, laptop, điện thoại hoặc máy tính bảng của bạn — không cần app store. Trên desktop nó mở trong cửa sổ riêng không có chrome của trình duyệt, nên nó trông và cảm giác như một ứng dụng desktop thực sự.

Hãy nghĩ về nó như **Claude Cowork của riêng bạn** — một không gian làm việc AI cá nhân sống trên máy của bạn — ngoại trừ việc nó mã nguồn mở và không phụ thuộc model:

- **Bạn sở hữu toàn bộ stack.** Chọn bất kỳ model nào, chuyển đổi bất cứ khi nào bạn thích. Chạy một model cục bộ và dữ liệu của bạn sẽ không bao giờ rời khỏi máy của bạn.
- **Người không chuyên kỹ thuật cũng có thể sử dụng.** Cài đặt pi-web trên máy của họ, chỉ cho họ cách sử dụng một lần, và họ sẽ ổn. Cha mẹ bạn, bạn đời của bạn, những người bạn không rành công nghệ của bạn — không có terminal, không có SSH, chỉ có một giao diện chat quen thuộc.
- **Một lần cài đặt, nhiều người dùng.** Cài đặt nó trên máy để bàn của bạn và chia sẻ màn hình, hoặc phơi bày nó trên mạng gia đình và để các thành viên gia đình mở nó trên thiết bị riêng của họ.

Muốn nhiều hơn mã hóa? Biến nó thành một [trợ lý cá nhân](personal-assistant.md) chuyên dụng biết bạn là ai và sống trên máy của bạn — giống như OpenClaw hay Hermes của riêng bạn.

> 💡 **Mẹo hay:** Cài đặt pi-web như một PWA từ Chrome/Edge (nhấp vào biểu tượng cài đặt trên thanh địa chỉ) hoặc Safari (Share → Add to Dock). Nó trở nên không thể phân biệt với một ứng dụng native.

---

## Những gì bạn có thể làm với pi-web

| | |
|---|---|
| 📱 **PWA** | Cài đặt pi-web như một Progressive Web App trên desktop, điện thoại hoặc máy tính bảng cho cảm giác native. |
| 🔄 **Tiếp tục phiên** | Tiếp nhận bất kỳ cuộc trò chuyện nào ngay ở nơi bạn đã dừng — văn bản, hình ảnh, chuyển đổi model, tất cả từ trình duyệt. |
| 🆕 **Bắt đầu phiên mới** | Tạo các phiên mới cho bất kỳ đường dẫn dự án nào, ngay từ giao diện web. |
| 📡 **Streaming trực tiếp** | Xem các phản hồi của pi được stream theo thời gian thực với độ trễ ~ms. Chế độ Follow giữ bạn bám theo nội dung mới nhất. |
| 🌲 **Góc nhìn cây** | Điều hướng cây message native của pi — xem cấu trúc cuộc trò chuyện đầy đủ, nhảy đến bất kỳ branch nào và fork từ bất kỳ điểm nào. |
| 🔀 **Fork phiên** | Fork một phiên từ bất kỳ message nào hoặc thậm chí từ một cuộc gọi tool cụ thể — khám phá các hướng khác nhau mà không mất vị trí của bạn. |
| 🔍 **Duyệt & tìm kiếm** | Lọc các phiên xuyên suốt các dự án, tìm kiếm theo tên, điều hướng các branch — toàn bộ lịch sử phiên của bạn trong một cái nhìn. |
| 🌿 **Tích hợp Git** | Xem branch hiện tại và mở một GitHub PR ngay từ trình xem phiên. |
| 📝 **Scratchpad** | Ghi lại các ghi chú, todos, hoặc ý tưởng nhanh bên cạnh các phiên của bạn mà không cần chuyển đổi ứng dụng. |
| 💬 **Ghi chú** | Tô sáng và bình luận bất kỳ phần nào của một phiên — tuyệt vời cho code review, phản hồi, hoặc đánh dấu các khoảnh khắc quan trọng. |
| 🎨 **Themes & tùy chỉnh** | Chuyển đổi giữa chế độ tối và sáng, chỉnh sửa UI theo ý bạn — khiến pi-web cảm giác giống như *của riêng bạn*. |
| 🌐 **Đa ngôn ngữ** | 14 ngôn ngữ tích hợp sẵn (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Thêm ngôn ngữ tùy chỉnh của bạn từ Settings. |
| 🐱 **Wellness & pomodoro** | Quá nhiều vibe coding không tốt cho sức khỏe. Bộ đếm thời gian pomodoro tích hợp với một người bạn đồng hành mèo và các nhắc nhở ngủ để giữ bạn cân bằng. |
| 📤 **Chia sẻ & xuất** | Tải xuống JSONL, xuất các snapshot tĩnh được render với diện mạo native `pi.dev` của pi, hoặc chia sẻ dưới dạng GitHub Gist riêng tư — tất cả đều được render phía client. |
| 🔔 **Âm thanh thông báo** | Chuông thông báo tùy chỉnh cho các sự kiện phiên — giữ liên lạc ngay cả khi pi-web ở một tab khác. |
| ⌨️ **Phím tắt** | Điều hướng phong cách Vim, các hành động nhanh — [tham khảo đầy đủ →](keyboard-shortcuts.md) |
| 🤖 **Trợ lý cá nhân** | Biến pi-web thành trợ lý AI của riêng bạn sống trên máy tính của bạn — như OpenClaw hay Hermes. [Cài đặt nó →](personal-assistant.md) |
| 🗓️ **Nói chuyện với lịch trình** | Từ một phiên pi, nói "thêm một lịch trình lúc 2 giờ sáng giờ Singapore để …" — `/skill:pi-web-schedule`. |
| 📝 **Nói chuyện với ghi chú & cài đặt** | "Viết điều này vào ghi chú" (`/skill:pi-web-notes`) hoặc "chuyển sang chế độ tối" (`/skill:pi-web-settings`). |

---

## Điều hướng nhanh

| Nếu bạn đang tìm kiếm… | Đọc |
|---|---|
| Cách cài đặt, cấu hình và sử dụng pi-web | [install.md](install.md) |
| Sử dụng pi-web như một trợ lý cá nhân | [personal-assistant.md](personal-assistant.md) |
| Tham khảo phím tắt | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Tại sao pi-web tồn tại | [why.md](why.md) |
| Điều gì sắp đến tiếp theo | [roadmap.md](roadmap.md) |
| Gặp sự cố khi cài đặt? Để LLM của bạn sửa — dán liên kết llm-debug.md cho họ | [llm-debug.md](llm-debug.md) |
| Bảo trì phiên bản model cục bộ này | [development notes](../../docs/dev/local-llm-development.md) |

---

## Ảnh chụp màn hình

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Nhà tài trợ

pi-web được xây dựng bằng tình yêu và rất nhiều đêm khuya. Tôi tự trả tiền cho các gói coding (Claude Code, OpenCode, v.v.) để giữ cho dự án này tiến lên phía trước. Nếu pi-web đã hữu ích cho bạn, sự hỗ trợ của bạn sẽ có ý nghĩa vô cùng lớn.

**Cách giúp đỡ:**

- 💰 **[Tài trợ trên GitHub](https://github.com/sponsors/setkyar)** — giúp trang trải các công cụ giúp điều này trở thành hiện thực
- ☕ **[Mua cho tôi một ly cà phê](https://buymeacoffee.com/setkyar)** — mỗi chút đều hữu ích
- ⭐ **Sao lưu kho** — chi phí bằng không và giúp nhiều người hơn khám phá pi-web
- 📢 **Chia sẻ với bạn bè & gia đình** — nếu bạn biết ai đó sẽ yêu thích pi-web, hãy gửi cho họ

Không thể tài trợ? Hoàn toàn không sao — một lượt sao và một lần chia sẻ cũng rất có giá trị. Cảm ơn bạn đã ở đây. 🙏

---

Chúc bạn coding vui vẻ! 🚀
