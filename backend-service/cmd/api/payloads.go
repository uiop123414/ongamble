package main

var CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var LoginUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var CreateArticlePayload struct {
	ArticleName string        `json:"article_name"`
	Username    string        `json:"username"`
	Time        string        `json:"time"`
	Blocks      []interface{} `json:"blocks"`
	Publish     bool          `json:"publish"`
	Version     string        `json:"version"`
}

var CreateAiArticlePayload struct {
	ArticleName string `json:"article_name"`
	Request     string `json:"request"`
	Type        string `json:"type"`
}
