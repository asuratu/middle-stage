package es

type User struct {
	Name         string `json:"name"`
	City         string `json:"city"`
	Introduction string `json:"introduction"`
}

const (
	UserMappings = `
	{
		"aliases": {
			"user":{}
		},
		"settings": {
			"analysis": {
				"normalizer": {
					"lowercase": {
						"type": "custom",
						"char_filter": [],
						"filter": ["lowercase"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "text",
					"analyzer": "ik_max_word"
				},
				"city": {
					"type": "keyword",
					"normalizer": "lowercase"	
				},
				"introduction": {
					"type": "text",
					"analyzer": "ik_max_word"
				}
			}
		}
	}
	`
)
