# Cách giải quyết vấn đề cho bài toán tìm ngày active thực tế của số điện thoại
-------------------------------------------------------------------------------

## Lưu ý
Giải pháp dưới đây chỉ đúng khi data là hợp lệ. Data hợp lệ có các tính chất như sau:
* Với mỗi số điện thoại chỉ có 1 bản ghi duy nhất có deactivation là null
* Các khoảng thời gian activation - deactivation của một số điện thoại không bị trùng nhau.

## Yêu cầu đề bài
* Dữ liệu đầu vào không được sắp xếp, tuy nhiên mỗi số điện thoại tại 1 thời điểm chỉ có 1 người sử dụng duy nhất.
* Ngày active thực tế sẽ là ngày active đầu tiên của chuỗi lần sử dụng gần nhất.

## Giải pháp
* Sắp xếp dữ liệu đầu vào theo activation date trên toàn bộ tập data đọc được theo thứ tự tăng dần, không phần biệt của số điện thoại nào. Nếu 2 số có cùng activation date, số nào xếp trước sẽ đứng trên.
* Đưa thông tin vào map, có key là số điện thoại, value là cặp giá trị {activation date, deactivation date}.
  * Nếu chưa tồn tại số điện thoại, add cặp {activation date, deactivation date} hiện tại vào.
  * Nếu đã tồn tại số điện thoại:
    * Nếu activation date trùng với deactivation date trong map, update deactivation trong map bằng date mới vào.
    * Nếu activation date không trùng với deactivation date trong map, update cả activation date trong map bằng activation date mới, deactivation date trong map bằng deactivation date mới.
* Xuất list data với thông tin số điện thoại và activation date tương ứng ra, ta sẽ được kết quả.

