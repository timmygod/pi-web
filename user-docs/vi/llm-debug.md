> Hướng dẫn khắc phục sự cố này được bảo trì cho phiên bản local-model. Giữ các chi tiết triển khai cục bộ và ghi chú đồng bộ hóa với upstream nhất quán với [Phát triển phiên bản local-model](../../docs/dev/local-llm-development.md).

Người dùng đang cài đặt pi-web qua

- pi install npm:@timmygod/pi-web-local

Nào tự động chạy [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) và thiết lập pi.

Nếu người dùng gặp sự cố, có thể do việc thiết lập của install.sh gây ra. Bạn có thể xem xét và kiểm tra điều gì đang diễn ra, sau đó cho người dùng biết chính xác điều gì đang gây ra vấn đề. Và hỏi xem họ có muốn bạn sửa hay không. Luôn xác nhận với người dùng.

Để người dùng có thể truy cập pi trên điện thoại hoặc mạng khác, cách được khuyến nghị là sử dụng tailscale và truy cập nó từ mạng tailscale. Và người dùng sẽ phải bật HTTPs trong bảng điều khiển tailscale của họ - https://login.tailscale.com/admin/dns

Nếu họ không có tailscale cài đặt hoặc không muốn sử dụng tailscale, họ có thể chạy `pi-web status` và nhận được đường dẫn của binary, trạng thái của binary và endpoint cục bộ mà họ có thể truy cập ứng dụng. Lưu ý rằng họ sẽ không thể nhận được thông báo push vì nó đang ở trên http.

Trên macOS, nó sử dụng [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
Trên Linux, nó sử dụng [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

Trong trường hợp bạn cần gỡ lỗi thêm và xem điều gì đang diễn ra.
