package db

import (
	"fmt"
	"graphql-boilerplate/configs"
	"strings"
)

// Применяется для исполнения запросов SELECT.
func QueryRowBytes(sqlText string, values ...interface{}) ([]byte, error) {
	var result []byte
	fmt.Printf(configs.DebugColor+configs.NormalColor, sqlText, values)
	err := Conn.QueryRowx(sqlText, values...).Scan(&result)
	return result, err
}

// Возвращает: строки для jsonb_build_object и перечисленных полей
// bf - buildFields; sf - selectedFields
func GetFields(selectedFields []string) (string, string) {
	var sf = make([]string, 0, len(selectedFields))
	// собираем поля для jsonb_build_object
	var buildFields = make([]string, 0, len(selectedFields))
	for _, v := range selectedFields {
		buildFields = append(buildFields, `'`+v+`', `+v)
		// есть поля-функции
		if val, ok := tableMap[v]; ok {
			sf = append(sf, val.Func)
			continue
		}
		sf = append(sf, v)
	}
	return strings.Join(buildFields, ", "), strings.Join(sf, ", ")
}

// TableMap структура для соответствия graphql полей и функций в БД
type TableMap struct {
	Func       string // вызываемая функция в PG
	ExtraField string // дополнительное поле
}

// Карта соответствия graphql полей и функций в БД
var tableMap = map[string]TableMap{
	"questions": TableMap{
		Func: "questions(id, order_questions)",
	},
	"greeting": TableMap{
		Func: "greeting(id)",
	},
	"results": TableMap{
		Func: "results(id)",
	},
	"kind": TableMap{
		Func: "kind(kind_id)",
	},
	"media": TableMap{
		Func: "media(media_id)",
	},
	"answers": TableMap{
		Func: "answers(id, order_answers)",
	},
}

// Filter for DB query
type Filter struct {
	Conditions []string
	Limit      int
	Offset     int
}
