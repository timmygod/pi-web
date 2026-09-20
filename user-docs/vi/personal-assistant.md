# pi-web như trợ lý cá nhân của bạn

Workflow này được hỗ trợ bởi phiên bản local-model. Về chính sách phát triển, đồng bộ hóa và phát hành của phiên bản, xem
[Local-model edition development](../../docs/dev/local-llm-development.md).

pi-web không chỉ dành cho coding — bạn có thể biến nó thành **trợ lý AI cá nhân** chạy ngay trên máy của bạn, giống như có riêng OpenClaw hoặc Hermes của mình.

## Cách hoạt động

Bạn tạo một thư mục riêng trên máy của mình — đó là nơi trợ lý của bạn "sinh sống". Bên trong, bạn đặt file `APPEND_SYSTEM.md` để định nghĩa trợ lý của bạn là ai, biết gì và hành xử ra sao. pi-web cung cấp cho bạn một giao diện chat đẹp mắt để trò chuyện với nó từ bất kỳ thiết bị nào.

## Từng bước

### 1. Tạo thư mục trợ lý của bạn

Chọn một thư mục trên máy tính. Ví dụ:

```
~/my-assistant/
```

### 2. Định nghĩa trợ lý của bạn

Tạo file `APPEND_SYSTEM.md` bên trong thư mục đó. Đây là nơi bạn nói cho pi biết trợ lý của bạn là ai:

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

pi tự động thêm nội dung này vào system prompt của mỗi cuộc trò chuyện, vì vậy trợ lý của bạn luôn biết bạn là ai và cách giúp đỡ.

### 3. Khởi tạo phiên trong thư mục đó

Trong pi-web, tạo một phiên mới trỏ đến `~/my-assistant/` (hoặc bất kỳ tên nào bạn đặt). Xong — bạn đang trò chuyện với trợ lý cá nhân của mình.

### 4. Sử dụng ở bất cứ đâu

Cài đặt pi-web dưới dạng PWA trên điện thoại, máy tính bảng hoặc laptop của bạn. Trợ lý của bạn luôn ở đó — hỏi bất cứ điều gì, bất cứ khi nào.

## Ý tưởng cho trợ lý của bạn

| Vai trò | Nội dung đưa vào APPEND_SYSTEM.md |
|---|---|
| 🧠 **Trợ lý phát triển bản thân** | Mục tiêu, thói quen bạn đang rèn luyện, gợi ý viết nhật ký |
| 🏠 **Người quản lý gia đình** | Định dạng danh sách tạp hóa, sở thích của các thành viên, kế hoạch bữa ăn |
| 💼 **Đồng nghiệp ảo** | Vai trò của bạn, dự án hiện tại, định dạng ghi chú cuộc họp, bối cảnh công ty |
| 📚 **Bạn học cùng** | Những gì bạn đang học, phong cách giải thích ưa thích, chế độ quiz |
| ✍️ **Trợ lý viết lách** | Phong cách viết, sở thích giọng văn, các định dạng bạn thường dùng |

## Thêm nhiều ngữ cảnh hơn

Bạn có thể đặt bất cứ thứ gì trong thư mục trợ lý giúp ích cho việc pi trở nên hữu dụng hơn:

- `notes/` — các file tham khảo mà trợ lý có thể đọc
- `context.md` — thông tin nền về cuộc sống hoặc công việc của bạn
- `projects.md` — các dự án hiện tại và trạng thái của chúng

pi có thể đọc các file trong thư mục, nên bạn càng cung cấp nhiều ngữ cảnh, nó càng tốt hơn.

## Ra lệnh cho pi-web làm việc

Sau `pi install npm:@timmygod/pi-web-local`, các phiên có thể trò chuyện với chính pi-web.
Thử:

- "Thêm lịch trình lúc 2 giờ sáng giờ Singapore để tóm tắt hộp thư của tôi"
- "Liệt kê các lịch trình pi-web của tôi"
- "Tạm dừng lịch trình hộp thư"
- "Ghi điều này vào notes"
- "Chuyển pi-web sang dark mode / tắt auto-title"

Skill **/skill:pi-web-schedule** có sẵn sẽ biến điều đó thành một lịch trình pi-web thực sự (cùng loại mà bạn chỉnh sửa tại `/schedules`). Mỗi lần chạy sẽ khởi tạo một phiên **mới**, nên các chỉ thị phải tự hoàn chỉnh — "tóm tắt email chưa đọc trong ~/inbox" thì hoạt động; "tiếp tục việc chúng ta đang làm" thì không.

Các lịch trình chỉ chạy khi pi-web đang chạy.

---

> 💡 **Mẹo:** Hãy bắt đầu đơn giản. Chỉ vài dòng về bạn là ai và cách bạn muốn trợ lý hành xử. Cải tiến dần theo thời gian khi bạn tìm ra điều gì hiệu quả.
