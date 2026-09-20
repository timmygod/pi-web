# Vì sao chọn pi-web?

Tôi hơi bị mê mẩn Claude Code. Tôi luôn dùng nó. Nếu không ngồi trước máy tính, tôi cũng nghĩ về nó. Tôi có cảm giác như mình chưa đốt đủ token. Đó là những ngày đầu của Claude Code. Và tôi tự hỏi, tại sao tôi không thể tiếp tục từ điện thoại? Tôi đã cài đặt Termius nhưng không thực sự thích nó.

Tôi bắt đầu tự tạo cho mình, rồi dừng lại khi Claude ra mắt ứng dụng di động Claude Code của họ.

Sau đó tôi bị thoát vị đĩa đệm và không thực sự làm được gì nhiều. Thời gian trôi qua, tôi cảm thấy hồi phục đôi chút và muốn tiếp tục dự án Claude Code qua web/PWA của mình.

Rồi Claude Code bắt đầu cấm việc sử dụng bên ngoài harness của chính họ. Và tôi cảm thấy không còn đáng làm nữa.

Sau đó tôi tìm thấy pi.dev và khám phá đôi chút, nhưng chưa thực sự đi sâu. Tôi đã đọc về nó, xem các video về nó và quyết định thử hoàn toàn, và giờ tôi đã hoàn toàn say mê pi.

Vì đây là mã nguồn mở, tôi cảm thấy đáng để đầu tư xây dựng. Tôi cũng có thêm các lựa chọn nhà cung cấp khác. Tôi cũng cảm thấy việc phụ thuộc vào một nhà cung cấp/mô hình như Anthropic/Claude không bền vững.

Vậy nên tôi đang xây dựng nó ở đây.

Phiên bản checkout này được bảo trì như một bản local-model của pi-web. Nó làm theo dự án upstream cho các cải tiến chung, đồng thời giữ việc triển khai cục bộ, tính ổn định ngữ cảnh và kiểm tra mô hình cục bộ trên một tuyến phát hành riêng.

## Vì sao một mô hình cục bộ cần một hồ sơ vận hành khác

Trải nghiệm pi-web gốc là một nền tảng tuyệt vời, nhưng suy luận cục bộ có các chế độ lỗi khác so với một mô hình được lưu trữ theo kiểu thông thường. Một mô hình cục bộ có thể chậm lại đột ngột khi ngữ cảnh mở rộng, chia sẻ bộ nhớ giới hạn với phần còn lại của máy, dừng lại sau khi chỉ tạo ra nội dung suy luận, hoặc mất một phiên chạy dài do một lỗi truyền tải cục bộ tạm thời. Xử lý những trường hợp đó hoàn toàn giống với lỗi cloud khiến giao diện trông như tương thích trong khi phiên thực sự vẫn mong manh.

Phiên bản này tiếp cận vấn đề theo từng lớp:

1. **Giữ nguyên upstream trước.** Giao diện và hành vi phiên chung vẫn đến từ pi-web; các thay đổi cục bộ được cô lập phía sau Local Mode hiệu lực.
2. **Phòng ngừa trước khi phục hồi.** Một ranh giới ngữ cảnh 65% dựa trên phần trăm được áp dụng trước các lần gọi mô hình tiếp theo, bao gồm cả các lần gọi bên trong các vòng lặp tool dài.
3. **Chỉ phục hồi khi có bằng chứng.** Việc tiếp tục tự động được giới hạn ở các sự cố ngữ cảnh, truyền tải và chỉ-suy-đ Coalition đã được nhận dạng—không bao gồm lỗi xác thực, hạn ngạch hoặc lỗi nhà cung cấp tùy ý.
4. **Đặt giới hạn cho mọi hành động tự chủ.** Các sự cố phục hồi được khử trùng lặp, phải có tiến trình trước khi cứu hộ tiếp theo, và khi khởi động chỉ xem xét tối đa một phiên Local gần đây đang hoạt động.
5. **Giữ một lối thoát thủ công.** Force Compact tóm tắt thay vì xóa lịch sử, để người dùng có thể cứu hộ một phiên mà không phải giả vờ rằng ngữ cảnh chưa bao giờ tồn tại.
6. **Bảo vệ tính tương thích cloud.** Cloud Mode giữ nguyên ngữ nghĩa và các điều khiển của upstream; các tối ưu hóa cho mô hình cục bộ không âm thầm định nghĩa lại các phiên cloud.

Đó chính là sự khác biệt thực sự trong nhánh fork này: nó coi suy luận cục bộ như một môi trường vận hành riêng biệt, chứ không đơn thuần là một tên mô hình khác trong danh sách thả xuống.
