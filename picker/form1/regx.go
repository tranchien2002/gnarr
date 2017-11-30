package form1

/*  thi,s the regexship!
    na na, not relationship, it,s callled regexship!
    what,s cell[0]?
    the jail.
*/

/* this contain "ID" in 2nd jail cell,
   "Enforcer" in 3rd, 4th jail cell
*/
const idEnforcerRegx = `(Số|số|SỐ|SỐ|số)\s*(?::|\.{0,5})\s*(\d+\/\d+\/([A-Z][A-Z])(\d+))`

/*  this contain "Type" in first cell,
    "Name" in second
*/
const nameRegx = `\n(LUẬT|PHÁP LỆNH|BỘ LUẬT)\s*(.+)`

/*  this contain "BaseOn" in jail, alright, all of the jail */
const baseonRegx = `(Căn cứ|căn cứ)(.+)`

/*  this contain "Index of Chapter" in first cell,
    the 2nd,s for "Header of Chapter"
    haizz, do ya see last jail,
    wtf the last alive for?
*/
const chapterRegx = `\n(?:Chương|CHƯƠNG)\s*(M{1,4}(?:CM|CD|D?C{0,3})(?:XC|XL|L?X{0,3})(?:IX|IV|V?I{0,3})|M{0,4}(?:CM|C?D|D?C{1,3})(?:XC|XL|L?X{0,3})(?:IX|IV|V?I{0,3})|M{0,4}(?:CM|CD|D?C{0,3})(?:XC|X?L|L?X{1,3})(?:IX|IV|V?I{0,3})|M{0,4}(?:CM|CD|D?C{0,3})(?:XC|XL|L?X{0,3})(?:IX|I?V|V?I{1,3})|\d+)(?:\:|\.)?\s+(.+)|(?:\.\s*(C.+)\s*(?:\((?:Đ|đ)ã ký))`

/*  this contain "Index of Topic" in second cell,
    the 3rd ,s "Header of Topic"
*/
const topicRegx = `\n(Mục|MỤC)\s*(\d+)\s*(.+)`

/*  "Header of Article" in 4th cell, index in 3rd.
    last jail just for entertainnn.
    naa, kiding, last jail for the PEACE!
    shot 'em nowww
    actually it contain "regency" in 6th. The jail in Chapter regexship ,s theeee same.
*/
const articleRegx = `(\n.iê?.(u|ù)\s+(\d+)(?:\.|\s)\s+(.+))|(\.\s*(C.+)\s*(\((Đ|đ)ã ký))`

/*  match index of the first article,
    a point to devide 'em
*/
const firstArticleRegx = `\n.iê?.(u|ù)\s+1`

/*  "day" in 2nd, "month" in 3rd and "year" in 4th*/
const effectiveDateRegx = `(?:này.*thi\s*hành.*|hiệu\s*lực.*)(ngày\s*(\d{1,2})\s*tháng\s*(\d{1,2})\s*năm\s*(\d{4}))`

/*  nt
    do ya meann "nt"??
*/
const passDateRegx = `thông\s*qua\s*(ngày\s*(\d{1,2})\s*tháng\s*(\d{1,2})\s*năm\s*(\d{4}))`

/*  first cell, last jail*/
const signRegx = `ký\)(?:\s)+((?:(?:\p{Lu}\p{Ll}+)\s*)+(?:\p{Lu}\p{Ll}+))`
