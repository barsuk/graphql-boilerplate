package article

import (
	g "github.com/graphql-go/graphql"
	jsoniter "github.com/json-iterator/go"
	graphqlfields "github.com/maxchagin/graphql-go-fields"
	graphqlscalar "github.com/maxchagin/graphql-go-scalar"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type List struct {
	Items  []Item
	Length int
}

type Item struct {
	ID             int64  `db:"id" json:"id,omitempty"`
	Title          string `db:"title" json:"title,omitempty"`
	SubTitle       string `db:"sub_title" json:"sub_title,omitempty"`
	LinkTitle      string `db:"link_title" json:"link_title,omitempty"`
	Announce       string `db:"announce" json:"announce,omitempty"`
	Lead           string `db:"lead" json:"lead,omitempty"`
	Text           string `db:"text" json:"text,omitempty"`
	SeoTitle       string `db:"seo_title" json:"seo_title,omitempty"`
	SeoDescription string `db:"seo_description" json:"seo_description,omitempty"`
	IsActive       bool   `db:"is_active" json:"is_active,omitempty"`
	IsAdv          bool   `db:"is_adv" json:"is_adv,omitempty"`
	IsNews         bool   `db:"is_news" json:"is_news,omitempty"`
	IsBroken       bool   `db:"is_broken" json:"is_broken,omitempty"`
	IsPaid         bool   `db:"is_paid" json:"is_paid,omitempty"`
	IsSpiegel      bool   `db:"is_spiegel" json:"is_spiegel,omitempty"`
	CreatedAt      string `db:"created_at" json:"created_at,omitempty"`
	ActivatedAt    string `db:"activated_at" json:"activated_at,omitempty"`
	ModifiedAt     string `db:"modified_at" json:"modified_at,omitempty"`
}

var ListObject = g.NewObject(
	g.ObjectConfig{
		Name: "ArticleList",
		Fields: g.Fields{
			"items": &g.Field{
				Type:        g.NewList(ItemObject),
				Description: "Статьи",
			},
			"length": &g.Field{
				Type:        g.Int,
				Description: "длина списка",
			},
		},
	})

var ItemObject = g.NewObject(
	g.ObjectConfig{
		Name: "Article",
		Fields: g.Fields{
			"id": &g.Field{
				Type:        graphqlscalar.Int64,
				Description: "Идентификатор",
			},
			"title": &g.Field{
				Type:        g.String,
				Description: "Заголовок",
			},
			"sub_title": &g.Field{
				Type:        g.String,
				Description: "Подзаголовок",
			},
			"link_title": &g.Field{
				Type:        g.String,
				Description: "Заголовок-ссылка",
			},
			"announce": &g.Field{
				Type:        g.String,
				Description: "Микроанонс",
			},
			"lead": &g.Field{
				Type:        g.String,
				Description: "Лид",
			},
			"text": &g.Field{
				Type:        g.String,
				Description: "Текст статьи/новости",
			},
			"seo_title": &g.Field{
				Type:        g.String,
				Description: "SEO заголовок",
			},
			"seo_description": &g.Field{
				Type:        g.String,
				Description: "SEO описание",
			},
			"is_active": &g.Field{
				Type:        g.Boolean,
				Description: "опубликован?",
			},
			"is_adv": &g.Field{
				Type:        g.Boolean,
				Description: "рекламный материал",
			},
			"is_news": &g.Field{
				Type:        g.Boolean,
				Description: "Материал — новость",
			},
			"is_broken": &g.Field{
				Type:        g.Boolean,
				Description: "сломанный материал",
			},
			"is_paid": &g.Field{
				Type:        g.Boolean,
				Description: "оплачиваемый",
			},
			"is_spiegel": &g.Field{
				Type:        g.Boolean,
				Description: "в шпигеле",
			},
			"created_at": &g.Field{
				Type:        g.DateTime,
				Description: "Дата и время создания материала",
			},
			"activated_at": &g.Field{
				Type:        g.DateTime,
				Description: "Дата и время публикации материала",
			},
			"modified_at": &g.Field{
				Type:        g.DateTime,
				Description: "Дата и время последнего изменения материала",
			},
		},
		Description: "Статья или новость",
	})

// List список статей
func GetList() *g.Field {
	return &g.Field{
		Name: "ArticleList",
		Type: ListObject,
		Args: g.FieldConfigArgument{
			"limit": &g.ArgumentConfig{
				Type:         g.Int,
				DefaultValue: 10,
				Description:  "Max limit: 100",
			},
			"offset": &g.ArgumentConfig{
				Type:         g.Int,
				DefaultValue: 0,
				Description:  "Default: 0",
			},
		},
		Resolve: func(p g.ResolveParams) (interface{}, error) {
			// здесь я "почему-то" совершил ошибку — написал list_articles вместо list_article.
			// Т.е. я к чему — как бы этой непонятки избежать?
			la, err := list(p.Args, graphqlfields.GetSelectedFields([]string{"list_article", "items"}, p.Info.FieldASTs))
			//la, err := list()
			if err != nil {
				return nil, errors.Wrap(err, "Список статей")
			}
			return la, nil
		},
		Description: "",
	}
}
