package schemas

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

const PATH = "/app/schemas/jsons"

func getJsonSchema(name string) gojsonschema.JSONLoader {
	return gojsonschema.NewReferenceLoader(fmt.Sprintf(`file://%s/%s.schema.json`, PATH, name))
}

var (
	CreateAiArticleLoader = getJsonSchema("CreateAiArticle")
)
