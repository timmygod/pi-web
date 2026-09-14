# Tại sao pi-web?

Tôi khá là nghiện Claude Code. Tôi luôn sử dụng nó. Nếu không ngồi trước máy tính, tôi lại nghĩ về nó. Tôi cảm thấy như mình chưa đốt đủ token. Đó là những ngày đầu của Claude Code. Và tôi đã nghĩ, tại sao mình không thể tiếp tục từ điện thoại? Tôi đã cài đặt Termius và thực sự không thích nó.

Tôi bắt đầu tự tạo và dừng lại khi Claude giới thiệu ứng dụng di động Claude Code của họ.

Rồi tôi bị thoát vị đĩa đệm và thực sự không thể làm được gì nhiều. Thời gian trôi qua và tôi cảm thấy hồi phục một chút và muốn tiếp tục dự án Claude Code qua web/PWA của mình.

Rồi Claude Code bắt đầu cấm sử dụng bên ngoài harness của họ. Và tôi cảm thấy điều đó không đáng.

Rồi tôi tìm thấy pi.dev và khám phá một chút nhưng chưa thực sự đào sâu. Tôi đã đọc về nó, xem video về nó và quyết định thử hết mình và giờ tôi hoàn toàn say mê pi.

Vì nó là mã nguồn mở, tôi cảm thấy nó đáng để xây dựng. Tôi cũng có nhiều lựa chọn nhà cung cấp khác nhau. Tôi cũng cảm thấy việc phụ thuộc vào một nhà cung cấp/mô hình như Anthropic/Claude là không bền vững.

Vì vậy tôi đang xây dựng nó ở đây.

## Tại sao mô hình cục bộ cần một hồ sơ hoạt động khác

Trải nghiệm pi-web ban đầu là một nền tảng tuyệt vời, nhưng suy luận cục bộ có các chế độ lỗi khác với mô hình được lưu trữ thông thường. Một mô hình cục bộ có thể chậm lại đáng kể khi ngữ cảnh tăng lên, chia sẻ bộ nhớ hạn chế với phần còn lại của máy, dừng lại sau khi chỉ tạo ra suy luận, hoặc mất một quá trình chạy dài do lỗi truyền tải cục bộ tạm thời. Xử lý những trường hợp đó chính xác như lỗi đám mây khiến giao diện người dùng trông tương thích trong khi phiên thực tế vẫn dễ bị tổn thương.

Phiên bản này tiếp cận vấn đề theo từng lớp:

1. **Giữ nguyên upstream trước.** Hành vi giao diện người dùng và phiên chia sẻ tiếp tục đến từ pi-web; các thay đổi cục bộ được cô lập phía sau Local Mode hiệu quả.
2. **Ngăn ngừa trước khi phục hồi.** Ranh giới ngữ cảnh 65% dựa trên phần trăm được thực thi trước các lệnh gọi nhà cung cấp sau đó, bao gồm các lệnh gọi trong vòng lặp công cụ dài.
3. **Chỉ phục hồi khi có bằng chứng.** Việc tiếp tục tự động bị giới hạn ở các sự cố ngữ cảnh, truyền tải và chỉ-suy-nghĩ được nhận dạng, không phải lỗi xác thực, hạn mức hoặc lỗi nhà cung cấp tùy ý.
4. **Giới hạn mọi hành động tự chủ.** Các sự cố phục hồi được loại bỏ trùng lặp, tiến độ là bắt buộc trước một lần cứu hộ khác, và khởi động chỉ xem xét tối đa một phiên Local hoạt động gần đây.
5. **Giữ lối thoát thủ công.** Force Compact tóm tắt thay vì xóa lịch sử, để người dùng có thể cứu phiên mà không giả vờ rằng ngữ cảnh chưa bao giờ tồn tại.
6. **Bảo vệ tính tương thích đám mây.** Cloud Mode giữ nguyên ngữ nghĩa và điều khiển upstream; các tối ưu hóa mô hình cục bộ không âm thầm định nghĩa lại các phiên đám mây.

Đó là sự khác biệt thực sự trong fork này: nó coi suy luận cục bộ là một môi trường hoạt động riêng biệt, không chỉ đơn thuần là một tên mô hình khác trong danh sách thả xuống.
