package picker

import (
	"encoding/json"
	"gnarr/picker/form1"
	"gnarr/structs"
)

/*  kiểm tra định dạng của văn bản thông qua chỉ số khớp được cung cấp từ các package form*, trả về định dạng có độ khớp cao nhất
 */
func checkFormat(legis []byte) string {

	matchMax := form1.Check(legis)
	format := "form1"
	if matchMax == 9 {
		return format
	}

	//TODO: Tìm định dạng phù hợp nhất

	return format
}

/*  xử lý dữ liệu thô từ file thành dữ liệu JSON
trả về dữ liệu JSON và string name
*/
func ToJSON(legis []byte) ([]byte, string) {

	//  kiểm tra định dạng
	format := checkFormat(legis)

	//đối tượng văn bản
	object := structs.Legistration{}

	//  chọn chiến thuật phù hợp với định dạng
	switch format {
	case "form1":
		object = *form1.Exec(legis)
	default:
	}

	//chuyển đối tượng thành dạng JSON
	jsonOut, error := json.Marshal(object)
	if error != nil {
		panic(error)
	}

	//tạo string name
	name := "uid"
	if len(object.PassDate)-4 >= 0 {
		name = object.PassDate[len(object.PassDate)-4:]
	}
	name = object.Name + "_" + name

	return jsonOut, name
}
