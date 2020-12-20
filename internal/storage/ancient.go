package storage

type ShiCi struct {
	ID      int    `gorm:"type:bigint(20);primary_key" json:"id"`
	Title   string `gorm:"type:varchar(64);index:title_idx;not null" json:"title,omitempty"`
	Author  string `gorm:"type:varchar(64);index:author_idx;not null" json:"author,omitempty"`
	Dynasty string `gorm:"type:varchar(32);index:dynasty_idx;not null" json:"dynasty,omitempty"`
	Content string `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
}

func (ShiCi) TableName() string {
	return "shici"
}

const AncientMappingForElasticsearch = `{
    "settings": {
        "analysis": {
            "tokenizer": {
                "ngram_tokenizer": {
                    "type": "nGram",
                    "min_gram": 1,
                    "max_gram": 10,
                    "token_chars": [
                        "letter",
                        "digit"
                    ]
                }
            },
            "analyzer": {
                "ngram_tokenizer_analyzer": {
                    "type": "custom",
                    "tokenizer": "ngram_tokenizer",
                    "filter": [
                        "lowercase"
                    ]
                }
            }
        },
        "max_ngram_diff": "10"
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "long"
			},
			"title": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"author": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"dynasty": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"content": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			}
		}
	}
}`
