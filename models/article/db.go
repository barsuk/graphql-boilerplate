package article

import (
	"fmt"
	"github.com/pkg/errors"
	"graphql-boilerplate/db"
	"strings"
)

// list список статей
func list(args map[string]interface{}, selectedFields []string) (*List, error) {
	bytes, err := All(args, selectedFields)
	if err != nil {
		return nil, errors.Wrap(err, "Не могу достать список материалов")
	}
	var lq *List
	err = json.Unmarshal(bytes, &lq)
	if err != nil {
		return nil, errors.Wrap(err, "Не могу размаршалить список")
	}
	return lq, nil
}

// All получение списка статей
func All(args map[string]interface{}, selectedFields []string) ([]byte, error) {
	filter := getArticleFilter(args)
	// поля для запросов к БД
	bf, sf := db.GetFields(selectedFields)
	query := `SELECT list_articles($1, $2, $3, $4, $5)`
	return db.QueryRowBytes(query, strings.Join(filter.Conditions, " AND "), bf, sf, filter.Limit, filter.Offset)
}

// Получаем фильтр для поиска викторин
func getArticleFilter(args map[string]interface{}) db.Filter {
	var f db.Filter
	// значение по-умолчанию для условия (WHERE), без него если условие не установлено - падает с ошибкой
	f.Conditions = append(f.Conditions, "1=1")
	for i, v := range args {
		switch i {
		case "search":
			// В поиске необходим placeholder, для формирования безопасного запроса
			f.Conditions = append(f.Conditions, fmt.Sprintf(`to_tsvector('russian', name || ' ' || description) @@ plainto_tsquery('russian', '%s')`, v))
			continue
		case "date_start":
			f.Conditions = append(f.Conditions, fmt.Sprintf(`datetime > %d`, v))
			continue
		case "date_end":
			f.Conditions = append(f.Conditions, fmt.Sprintf(`datetime < %d`, v))
			continue
		case "limit":
			f.Limit = v.(int)
			// За один запрос можно получить не более 100 элементов (сделано для оптимизации)
			if f.Limit > 100 {
				f.Limit = 100
			}
			continue
		case "offset":
			f.Offset = v.(int)
			continue
		}
		f.Conditions = append(f.Conditions, fmt.Sprintf(`%s = %v`, i, v))
	}
	return f
}
