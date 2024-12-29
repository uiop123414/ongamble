package models

var CreateAiArticlePayload struct {
	ArticleName string `json:"article_name"`
	Request     string `json:"request"`
	Type        string `json:"type"`
}