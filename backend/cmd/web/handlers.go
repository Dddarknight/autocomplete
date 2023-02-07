package main
 
import (
    "net/http"
	"encoding/json"
)
 
func autocomplete(writer http.ResponseWriter, request *http.Request) {
    word_to_find := request.URL.Query().Get("word")
    if word_to_find == "" {
        http.NotFound(writer, request)
        return
    }

    tree := MakeWordsTree()
	result_words := tree.Search(word_to_find, "")
	data := map[string]interface{} {
		"words": result_words,
	}
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}