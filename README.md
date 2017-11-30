# gnarr doc.

### thành phần

```go
package structs
định nghĩa các cấu trúc được sử dụng

package picker
chuyển đổi mảng byte ban đầu đọc từ file thành mảng byte JSON
	picker.go
	kiểm tra điều kiện, chọn chiến thuật và xử lý

	package picker/form*
	cách xử lý của từng định dạng
	
		regx.go
		các hằng regexp được sử dụng
		
		form*.go
		xử lý định dạng *
           	    func Check(legis []byte) int{
                       // trả về chỉ số khớp, dùng để lựa chọn chiến thuật
                    }
                    func Exec(legis []byte) *structs.Legistration{
                       // thực thi, trả về đối tượng Legistration hoàn thiện
                    }

package main
gnarr.go
nhận tham số là danh sách tên tệp plain text, tạo dạng JSON tương ứng trong thư mục output.
ghi chép quá trình trong tệp log.
```

