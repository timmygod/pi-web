# Phím tắt bàn phím

Các phím tắt này áp dụng cho phiên bản local-model của pi-web. Hành vi thời gian chạy cụ thể theo phiên bản được tài liệu hóa trong [Phát triển phiên bản local-model](../../docs/dev/local-llm-development.md).

## Trang index (`/`)

### Cuộn trang (phong cách vim)

Các phím tắt phong cách vim tương tự hoạt động trên mọi trang khi focus **không** nằm trong phần tử input, textarea, hay contenteditable.

| Phím tắt | Hành động |
|----------|--------|
| `j` | Cuộn xuống 300px |
| `k` | Cuộn lên 300px |
| `g g` | Cuộn đến đầu trang |
| `G` (Shift+G) | Cuộn đến cuối trang |
| `Escape` | Mất focus khỏi input đang hoạt động để điều hướng j/k hoạt động |

### Lệnh trang index

| Phím tắt | Ngữ cảnh | Hành động |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Cấp trang | Mở palette tìm kiếm/sessions |
| `⌘⇧L` / `Ctrl+Shift+L` | Cấp trang | Chuyển đổi giao diện hệ thống (sáng/tối) |
| `Escape` | Cấp trang | Đóng palette, menu, hoặc modal |
| `Enter` | Input đường dẫn session mới | Tạo session mới |

> `⌘K` / `Ctrl+K` cũng là phím tắt "focus thanh địa chỉ" của Chrome. Trình duyệt có thể chặn nó trừ khi focus nằm bên trong một input văn bản.

## Trang chi tiết session (`/session?id=...`)

### Cuộn trang (phong cách vim)

Các phím tắt này hoạt động trên cả trang index và trang session khi focus **không** nằm trong input, textarea, hay phần tử contenteditable.

| Phím tắt | Hành động |
|----------|--------|
| `j` | Cuộn xuống 300px |
| `k` | Cuộn lên 300px |
| `g g` | Cuộn đến đầu trang |
| `G` (Shift+G) | Cuộn đến cuối trang |
| `I` (Shift+I) | Focus vào textarea trình soạn thảo chat |
| `Escape` | Mất focus khỏi input đang hoạt động để điều hướng j/k hoạt động |

### Sidebar & điều hướng

| Phím tắt | Ngữ cảnh | Hành động |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Cấp trang | Chuyển đổi hiển thị sidebar |
| `⌘K` / `Ctrl+K` | Cấp trang | Mở palette danh sách session |
| `⌘T` / `Ctrl+T` | Cấp trang | Session mới |
| `⌘⇧L` / `Ctrl+Shift+L` | Cấp trang | Chuyển đổi giao diện hệ thống (sáng/tối) |
| `⌘⇧N` / `Ctrl+Shift+N` | Cấp trang | Chuyển đổi sidebar scratchpad / notes |

> `⌘K` và `⌘T` cũng là phím tắt trình duyệt (focus thanh địa chỉ / tab mới). Trình duyệt có thể chặn chúng trừ khi focus nằm bên trong một input văn bản.

### Trình soạn thảo chat

| Phím tắt | Ngữ cảnh | Hành động |
|----------|---------|--------|
| `Enter` | Textarea chat | Gửi tin nhắn |
| `Shift+Enter` | Textarea chat | Chèn xuống dòng |
| `Shift+Tab` | Textarea chat | Chuyển đến mức thinking tiếp theo (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Textarea chat | Mở popup trình chọn model (gõ để lọc, Enter để chọn, focus quay lại textarea) |

### Chuyển đổi hiển thị mục nhập

| Phím tắt | Ngữ cảnh | Hành động |
|----------|---------|--------|
| `t` | Khi focus **không** nằm trong input/textarea | Chuyển đổi hiển thị thinking |
| `o` | Khi focus **không** nằm trong input/textarea | Chuyển đổi hiển thị tools |
| `p` | Khi focus **không** nằm trong input/textarea | Chuyển đổi output của tool |

### Palette, menu & sheet

| Phím tắt | Ngữ cảnh | Hành động |
|----------|---------|--------|
| `Escape` | Cấp trang | Đóng palette, menu, hoặc sheet đang mở |
| `⌘K` / `Ctrl+K` | Cấp trang | Mở palette danh sách session |
| `ArrowUp` / `ArrowDown` | Palette danh sách session | Điều hướng kết quả session |
| `Enter` | Palette danh sách session | Mở session được chọn (hoặc session đầu tiên) |
| `ArrowUp` / `ArrowDown` | Popup trình chọn model | Điều hướng danh sách model |
| `Enter` | Popup trình chọn model | Chọn model đang được đánh dấu |
| `ArrowUp` / `ArrowDown` | Modal fork | Điều hướng tin nhắn |
| `Enter` | Modal fork | Fork từ tin nhắn đang được đánh dấu |
| `Tab` | Sheet toàn màn hình | Chuyển focus bên trong sheet |
| `Escape` | Sheet toàn màn hình | Đóng sheet |
