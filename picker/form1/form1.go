package form1

import (
	"gnarr/structs"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*  Kiểm tra mức độ phù hợp
trả về chỉ số khớp nhỏ hơn bằng 9
trọng số của "ID" và "article" là 2, các trọng số này có thể căn chỉnh lại cho phù hợp.
các tiêu chí đánh giá:
	số hiệu
	tên
	căn cứ
	điều
	ngày hiệu lực
	ngày thông qua
	chữ ký
*/
func Check(legis []byte) int {
	match := 0
	matched, error := regexp.Match(idEnforcerRegx, legis)
	if matched {
		match += 2
	}

	matched, error = regexp.Match(nameRegx, legis)
	if matched {
		match += 1
	}
	matched, error = regexp.Match(baseonRegx, legis)
	if matched {
		match += 1
	}
	matched, error = regexp.Match(articleRegx, legis)
	if matched {
		match += 2
	}
	matched, error = regexp.Match(effectiveDateRegx, legis)
	if matched {
		match += 1
	}
	matched, error = regexp.Match(passDateRegx, legis)
	if matched {
		match += 1
	}
	matched, error = regexp.Match(signRegx, legis)
	if matched {
		match += 1
	}
	if error != nil {
	}

	return match
}

/*  triển khai
 */
func Exec(legis []byte) *structs.Legistration {

	//TODO: thiết kế file log

	//  cấu hình package "log"
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	}

	log.SetOutput(f)
	log.SetPrefix("form1: ")

	result := structs.Legistration{}

	//  tìm vị trí điều đầu tiên
	re0 := regexp.MustCompile(firstArticleRegx)
	matchfirstArticle := re0.FindSubmatchIndex(legis)

	//  cắt riêng phần đầu văn bản để tìm số hiệu, tên,...
	legisHeader := []byte{}
	if len(matchfirstArticle) != 0 {
		legisHeader = legis[0:matchfirstArticle[0]]
	} else {
		log.Println("WARNING: FIRST_ARTICLE_REGX no match")
	}

	//  tìm số hiệu và cơ quan ban hành
	re1 := regexp.MustCompile(idEnforcerRegx)
	matchIdEnforce := re1.FindSubmatch(legisHeader)
	if len(matchIdEnforce) != 0 {
		result.ID = string(matchIdEnforce[2])
		result.Enforcer = string(matchIdEnforce[3]) + string(matchIdEnforce[4])
	} else {
		log.Println("WARNING: ID_ENFORCER_REGX no match")
	}

	//  tìm tên văn bản
	re2 := regexp.MustCompile(nameRegx)
	matchName := re2.FindSubmatch(legisHeader)
	if len(matchName) != 0 {
		result.Name = string(matchName[1]) + " " + string(matchName[2])
	} else {
		log.Println("WARNING: ID_NAME_REGX no match")
	}

	//  xác định căn cứ
	re3 := regexp.MustCompile(baseonRegx)
	matchBaseon := re3.FindAllSubmatch(legisHeader, -1)
	if len(matchBaseon) != 0 {
		for i := range matchBaseon {
			result.Baseon += (string(matchBaseon[i][0]) + "\n")
		}
	} else {
		log.Println("WARNING: BASE_ON_REGX no match")
	}

	//  xác định ngày hiệu lực
	re4 := regexp.MustCompile(effectiveDateRegx)
	matchEffectiveDate := re4.FindAllSubmatch(legis, -1)
	i := len(matchEffectiveDate)
	if i != 0 {
		result.EffectiveDate = string(matchEffectiveDate[i-1][2]) + "-" +
			string(matchEffectiveDate[i-1][3]) + "-" + string(matchEffectiveDate[i-1][4])
	} else {
		log.Println("WARNING: EFFECTIVE_DATE_REGX no match")
	}

	//  xác định ngày thông qua
	re5 := regexp.MustCompile(passDateRegx)
	matchPassDate := re5.FindAllSubmatch(legis, -1)
	j := len(matchPassDate)
	if j != 0 {
		result.PassDate = string(matchPassDate[j-1][2]) + "-" +
			string(matchPassDate[j-1][3]) + "-" + string(matchPassDate[j-1][4])
	} else {
		log.Println("WARNING: PASS_DATE_REGX no match")
	}

	//  xác định chũ ký
	re6 := regexp.MustCompile(signRegx)
	matchSign := re6.FindAllSubmatch(legis, -1)
	k := len(matchSign)
	if k != 0 {
		result.Sign = string(matchSign[0][1])
		for t := 1; t < k; t++ {
			result.Sign += " / " + string(matchSign[t][1])
		}
	} else {
		log.Println("WARNING: SIGN_REGX no match")
	}

	/* Phức tạp: Xử lý các cấu trúc chương, mục, điều
	 * dựa vào vị trí các chương (nếu có) để chia nhỏ văn bản, tương tự với cấu trúc điều,
	 * sau đó tạo mảng article của đối tượng từ các đoạn văn bản nhỏ
	 */
	result.ChapterArray = []structs.Chapter{}
	result.ArticleArray = []structs.Article{}
	regx := regexp.MustCompile(chapterRegx)
	matchChapter := regx.FindAllSubmatch(legis, -1)
	matchChapterIndex := regx.FindAllSubmatchIndex(legis, -1)

	switch m := len(matchChapter); m {

	//  trường hợp biểu thức tìm chương không khớp (cả dấu hiệu kết thúc ở nhóm cuối cũng không khớp)
	case 0:
		log.Println("WARNING: END_REGX no match")

		//  tạo mảng article với phần tử [0] là chức vụ người ký
		regencyElement := structs.Article{
			Header:  "regency",
			Content: "unidentify",
		}
		result.ArticleArray = append(result.ArticleArray, regencyElement)

		//  tiếp tục tạo mảng article với tất cả các điều còn lại
		pushAllArticleArray(&result, legis)

	//  trường hợp chỉ bắt được dấu hiệu kết thúc
	case 1:
		//  tạo mảng article với phần tử [0] là chức vụ người ký
		regency := ""
		if len(matchChapter[0]) < 4 {
			log.Println("WARNING: CHAPTER_REGX run into a problem")
		} else {
			regency = string(matchChapter[0][3])
		}
		regencyElement := structs.Article{
			Header:  "regency",
			Content: regency,
		}
		result.ArticleArray = append(result.ArticleArray, regencyElement)

		//  tiếp tục tạo mảng article với các điều còn lại
		pushAllArticleArray(&result, legis)

	//  trường hợp có các chương
	default:
		//  tạo phần tử đầu tiên chứa chức vụ người ký
		regency := ""
		if len(matchChapter[m-1]) < 4 {
			log.Println("WARNING: CHAPTER_REGX run into a problem")
		} else {
			regency = string(matchChapter[m-1][3])
		}
		regencyElement := structs.Article{
			Header:  "regency",
			Content: regency,
		}
		result.ArticleArray = append(result.ArticleArray, regencyElement)

		//  cắt dữ liệu theo từng chương
		for i := 0; i < m-1; i++ {
			chapterContent := legis[matchChapterIndex[i][1] : matchChapterIndex[i+1][0]+1]
			chapterHeader := string(matchChapter[i][2])

			//  tạo cáu trúc chương với dữ liệu đã cắt
			creatChapter(&result, chapterContent, chapterHeader)
		}

		//  xử lý nhóm cuối (có thể là phần cuối của văn bản, có thể là phần cuối của một chương,tùy thuộc nhóm cuối có khớp không)
		finalChapterContent := legis[matchChapterIndex[m-1][1]:]
		test := structs.Legistration{}
		test1, test2 := pushArticleArray(&test, finalChapterContent)
		if test1 != -1 && test2 != -1 {
			finalChapterHeader := string(matchChapter[m-1][2])
			creatChapter(&result, finalChapterContent, finalChapterHeader)
		}
	}

	//  trả về đối tượng đã được khởi tạo
	return &result
}

/*  hàm tạo các đối tượng chapter
 */
func creatChapter(target *structs.Legistration, contentPie []byte, header string) {
	regx := regexp.MustCompile(topicRegx)
	matchTopic := regx.FindAllSubmatch(contentPie, -1)

	switch n := len(matchTopic); n {

	//  trường hợp không có các mục
	case 0:
		first, last := pushArticleArray(target, contentPie)
		temp := structs.Chapter{
			Header:       header,
			FirstArticle: first,
			LastArticle:  last,
			TopicArray:   []structs.Topic{},
		}
		target.ChapterArray = append(target.ChapterArray, temp)

	//  trường hợp có các mục
	default:
		temp := structs.Chapter{
			Header:       header,
			FirstArticle: -1,
			LastArticle:  -1,
			TopicArray:   []structs.Topic{},
		}

		//  chia nhỏ đoạn văn bản thành các đoạn nhỏ hơn theo vị trí mục và tạo các mục này, đồng thời []article được tạo
		matchTopicIndex := regx.FindAllSubmatchIndex(contentPie, -1)
		for i := 0; i < n-1; i++ {
			topicContent := contentPie[matchTopicIndex[i][1] : matchTopicIndex[i+1][0]+1]
			topicHeader := string(matchTopic[i][3])
			creatTopic(target, &temp, topicContent, topicHeader)
		}
		finalTopicContent := contentPie[matchTopicIndex[n-1][1]:]
		finalTopicHeader := string(matchTopic[n-1][3])
		creatTopic(target, &temp, finalTopicContent, finalTopicHeader)
		temp.FirstArticle = temp.TopicArray[0].FirstArticle
		temp.LastArticle = temp.TopicArray[n-1].LastArticle

		//  thêm temp vào []chapter của đối tượng
		target.ChapterArray = append(target.ChapterArray, temp)
	}

}

/*  tạo các điều trong một đoạn văn bản, đưa vào []article của đối tượng,
trả về chỉ số của điều đầu tiên và cuối cùng trong đoạn văn bản.
*/
func pushArticleArray(target *structs.Legistration, contentPie []byte) (int64, int64) {
	regx := regexp.MustCompile(articleRegx)
	match := regx.FindAllSubmatch(contentPie, -1)
	matchIndex := regx.FindAllSubmatchIndex(contentPie, -1)
	n := len(match)

	//  không khớp
	if n == 0 {
		log.Println("WARNING: ARTICLE_REGX run into trouble, but don't worry until second time (_'')")
		return -1, -1
	}

	/* (_''_) so cumplikate
	     (__''_) (_'') (''_)
	   ('')(_''_) (''_)
	*/

	//  chia các diều là dựa vào vị trí, tạo các điều đưa vào []array
	for i := 0; i < n-1; i++ {
		content := strings.TrimSpace(string(contentPie[matchIndex[i][1] : matchIndex[i+1][0]+1]))
		temp := structs.Article{
			Header:  string(match[i][4]),
			Content: content,
		}
		target.ArticleArray = append(target.ArticleArray, temp)
	}

	//  tương tự như creatChapter, cần xử lý thêm nhóm cuối để tăng độ chính xác
	finalArticleContent := strings.TrimSpace(string(contentPie[matchIndex[n-1][1]+1:]))
	finalArticleHeader := string(match[n-1][4])
	finalArticle := structs.Article{
		Header:  finalArticleHeader,
		Content: finalArticleContent,
	}
	target.ArticleArray = append(target.ArticleArray, finalArticle)

	//  xác định chỉ số đầu và cuối
	a := string(match[0][3])
	first, err1 := strconv.ParseInt(a, 10, 64)
	if err1 != nil {
		log.Panic(err1)
	}
	b := string(match[n-1][3])
	last, err2 := strconv.ParseInt(b, 10, 64)
	if err2 != nil {
		log.Panic(err2)
	}
	return first, last
}

/*  đưa toàn bộ điều của toàn văn bản vào []article
 */
func pushAllArticleArray(target *structs.Legistration, contentPie []byte) {

	//  chia nhỏ thành các điều, tạo và đưa vào []article
	regx := regexp.MustCompile(articleRegx)
	match := regx.FindAllSubmatch(contentPie, -1)
	matchIndex := regx.FindAllSubmatchIndex(contentPie, -1)
	n := len(match)
	if n < 2 {
		log.Println("WARNING: ARTICLE_REGX run into a problem on PushAll")
	}
	for i := 0; i < n-1; i++ {
		content := strings.TrimSpace(string(contentPie[matchIndex[i][1] : matchIndex[i+1][0]+1]))
		temp := structs.Article{
			Header:  string(match[i][4]),
			Content: content,
		}
		target.ArticleArray = append(target.ArticleArray, temp)
	}
}

/*  tạo topic, tương tự như chapter
 */
func creatTopic(target *structs.Legistration, targetChapter *structs.Chapter, contentPie []byte, header string) {

	//  like dad, like son~
	first, last := pushArticleArray(target, contentPie)
	temp := structs.Topic{
		Header:       header,
		FirstArticle: first,
		LastArticle:  last,
	}
	targetChapter.TopicArray = append(targetChapter.TopicArray, temp)
}
