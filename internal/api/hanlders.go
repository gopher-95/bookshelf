package api

import (
	"io"
	"net/http"
)

// читаем тело запроса и отправляем его в базу данных, предварительно сериализовав данные из Json в обычный формат
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	io.ReadAll(r.Body)
}
