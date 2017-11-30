package form1

/*  các hằng regexp sử dụng trong form1
    nhóm đầu tiên ([0]) trong định dạng trả về là toàn bộ xâu khớp
    các nhóm tiếp theo lần lượt ẳ [1] [2] ...
*/

/* trường "ID" ở nhóm 2,
   "Enforcer" ở nhóm 3 và 4 (Cơ quan ban hành và khóa?)
*/
const idEnforcerRegx = `(Số|số|SỐ|SỐ|số)\s*(?::|\.{0,5})\s*(\d+\/\d+\/([A-Z][A-Z])(\d+))`

/*  Loại văn bản ở nhóm 1
    "Name" ở nhóm 2
*/
const nameRegx = `\n(LUẬT|PHÁP LỆNH|BỘ LUẬT)\s*(.+)`

/*  toàn bộ xâu khớp là "căn cứ" */
const baseonRegx = `(Căn cứ|căn cứ)(.+)`

/*  Chỉ số chương ở nhóm 1,
    tiêu đề ở nhóm 2,

    nhóm cuối cuối cùng được khớp bởi biểu thức này kỳ vọng là phần có chữ ký, có chứa chức vụ người ký ở nhóm 3
*/
const chapterRegx = `\n(?:Chương|CHƯƠNG)\s*(M{1,4}(?:CM|CD|D?C{0,3})(?:XC|XL|L?X{0,3})(?:IX|IV|V?I{0,3})|M{0,4}(?:CM|C?D|D?C{1,3})(?:XC|XL|L?X{0,3})(?:IX|IV|V?I{0,3})|M{0,4}(?:CM|CD|D?C{0,3})(?:XC|X?L|L?X{1,3})(?:IX|IV|V?I{0,3})|M{0,4}(?:CM|CD|D?C{0,3})(?:XC|XL|L?X{0,3})(?:IX|I?V|V?I{1,3})|\d+)(?:\:|\.)?\s+(.+)|(?:\.\s*(C.+)\s*(?:\((?:Đ|đ)ã ký))`

/*  chỉ số mục ở nhóm 2
    tiêu đề mục ở nhóm 3
*/
const topicRegx = `\n(Mục|MỤC)\s*(\d+)\s*(.+)`

/*  tiêu đề của điều ở nhóm 4, chỉ số điều ở nhóm 3
    nhóm cuối được khớp bởi biểu thức này tương tự biểu thức chapterRegx, chứa chức vụ ở nhóm 6
*/
const articleRegx = `(\n.iê?.(u|ù)\s+(\d+)(?:\.|\s)\s+(.+))|(\.\s*(C.+)\s*(\((Đ|đ)ã ký))`

/*  khớp điều đầu tiên
 */
const firstArticleRegx = `\n.iê?.(u|ù)\s+1`

/*  ngày, tháng, năm trong các nhóm 2, 3, 4*/
const effectiveDateRegx = `(?:này.*thi\s*hành.*|hiệu\s*lực.*)(ngày\s*(\d{1,2})\s*tháng\s*(\d{1,2})\s*năm\s*(\d{4}))`
const passDateRegx = `thông\s*qua\s*(ngày\s*(\d{1,2})\s*tháng\s*(\d{1,2})\s*năm\s*(\d{4}))`

/*  người ký ở nhóm 1*/
const signRegx = `ký\)(?:\s)+((?:(?:\p{Lu}\p{Ll}+)\s*)+(?:\p{Lu}\p{Ll}+))`
