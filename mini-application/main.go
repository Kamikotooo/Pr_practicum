package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

const fileName = "bookmarks.json"

type PageData struct {
	Bookmarks   map[string]string
	SearchQuery string
}

func loadBookmarks() map[string]string {
	bookmarks := make(map[string]string)
	if fileData, err := os.ReadFile(fileName); err == nil {
		json.Unmarshal(fileData, &bookmarks)
	}
	return bookmarks
}

func saveBookmarks(bookmarks map[string]string) {
	if fileData, err := json.MarshalIndent(bookmarks, "", "    "); err == nil {
		os.WriteFile(fileName, fileData, 0644)
	}
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Менеджер закладок</title>
</head>
<body>
    <h1>Управление закладками </h1>

    <!-- Поиск -->
    <form method="GET" action="/">
        <input type="text" name="search" value="{{.SearchQuery}}" placeholder="Поиск по закладкам...">
        <button type="submit">Найти</button>
        <a href="/">Сбросить</a>
    </form>

    <hr>

    <!-- Добавление -->
    <h3>Добавить закладку</h3>
    <form method="POST" action="/add">
        Название: <input type="text" name="name" required><br><br>
        URL: <input type="text" name="url" required><br><br>
        <button type="submit">Добавить</button>
    </form>

    <hr>

    <!-- Список -->
    <h3>Список закладок</h3>
    {{if .Bookmarks}}
        <ul>
        {{range $name, $url := .Bookmarks}}
            <li>
                <strong>{{$name}}</strong>: <a href="{{$url}}" target="_blank">{{$url}}</a>
                <form method="POST" action="/delete" style="display:inline;">
                    <input type="hidden" name="name" value="{{$name}}">
                    <button type="submit">Удалить</button>
                </form>
            </li>
        {{end}}
        </ul>
    {{else}}
        <p>Закладки не найдены.</p>
    {{end}}
</body>
</html>
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bookmarks := loadBookmarks()

		switch r.URL.Path {
		case "/":
			// 1. Просмотр и поиск закладок
			searchQuery := r.URL.Query().Get("search")
			filteredBookmarks := make(map[string]string)

			for name, url := range bookmarks {
				if searchQuery == "" ||
					strings.Contains(strings.ToLower(name), strings.ToLower(searchQuery)) ||
					strings.Contains(strings.ToLower(url), strings.ToLower(searchQuery)) {
					filteredBookmarks[name] = url
				}
			}

			tmpl, _ := template.New("index").Parse(htmlTemplate)
			data := PageData{
				Bookmarks:   filteredBookmarks,
				SearchQuery: searchQuery,
			}
			tmpl.Execute(w, data)

		case "/add":
			// 2. Добавление закладки
			if r.Method == http.MethodPost {
				name := r.FormValue("name")
				url := r.FormValue("url")
				if name != "" && url != "" {
					bookmarks[name] = url
					saveBookmarks(bookmarks)
				}
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)

		case "/delete":
			// 3. Удаление закладки
			if r.Method == http.MethodPost {
				name := r.FormValue("name")
				if name != "" {
					delete(bookmarks, name)
					saveBookmarks(bookmarks)
				}
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)

		default:

			http.NotFound(w, r)
		}
	})

	fmt.Println("Веб-сервер запущен! Откройте: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
