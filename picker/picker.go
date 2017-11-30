package picker

import (
	"encoding/json"
	"gnarr/picker/form1"
	"gnarr/structs"
)

func ToJSON(legis []byte) ([]byte, string) {
	format := checkFormat(legis)
	object := structs.Legistration{}

	switch format {
	case "form1":
		object = *form1.Exec(legis)
	default:
	}

	jsonOut, error := json.Marshal(object)
	if error != nil {
		panic(error)
	}

	name := "uid"
	if len(object.PassDate)-4 >= 0 {
		name = object.PassDate[len(object.PassDate)-4:]
	}
	name = object.Name + "_" + name

	return jsonOut, name
}

func checkFormat(legis []byte) string {

	matchMax := form1.Check(legis)
	format := "form1"
	if matchMax == 9 {
		return format
	}

	//TODO: Tìm định dạng phù hợp nhất

	return format
}
