package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const filename = "bookmarks.txt"

var bookmarks = make(map[string]string)

func main() {
	loadBookmarks()

	http.HandleFunc("/", showBookmarksHandler)
	http.HandleFunc("/add", addBookmarkHandler)
	http.HandleFunc("/delete", deleteBookmarkHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func showBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<meta charset='utf-8'>")

	fmt.Fprint(w, "<h2>Управление закладками</h2><a href='/'>Обновить список</a><br><br>")

	fmt.Fprint(w, `
		<form action="/" method="GET">
			<b>Поиск:</b> 
			<input type="text" name="search" placeholder="Что ищем?">
			<input type="submit" value="Найти">
			<a href="/">Сбросить поиск</a>
		</form><br>
	`)

	fmt.Fprint(w, `
		<form action="/add" method="GET">
			<b>Добавить:</b>
			<input type="text" name="name" placeholder="Название" required>
			<input type="text" name="url" placeholder="URL" required>
			<input type="submit" value="Сохранить">
		</form>
		<form action="/delete" method="GET">
			<b>Удалить:</b>
			<input type="text" name="name" placeholder="Название для удаления" required>
			<input type="submit" value="Удалить">
		</form><hr>
	`)

	if len(bookmarks) == 0 {
		fmt.Fprint(w, "<p>Список закладок пуст.</p>")
		return
	}

	searchQuery := r.URL.Query().Get("search")
	searchQuery = strings.ToLower(strings.TrimSpace(searchQuery))

	fmt.Fprint(w, "<h3>Ваши закладки:</h3><ul>")

	foundAny := false

	for name, url := range bookmarks {
		if searchQuery != "" {
			nameLower := strings.ToLower(name)
			urlLower := strings.ToLower(url)

			if !strings.Contains(nameLower, searchQuery) && !strings.Contains(urlLower, searchQuery) {
				continue
			}
		}

		foundAny = true
		fmt.Fprintf(w, "<li><b>%s</b>: <a href='%s' target='_blank'>%s</a></li>", name, url, url)
	}

	fmt.Fprint(w, "</ul>")

	if searchQuery != "" && !foundAny {
		fmt.Fprint(w, "<p>По вашему запросу ничего не найдено.</p>")
	}
}

func addBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	url := r.URL.Query().Get("url")

	if name != "" && url != "" {
		bookmarks[name] = url
		saveBookmarks()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		delete(bookmarks, name)
		saveBookmarks()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func saveBookmarks() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	for name, url := range bookmarks {
		fmt.Fprintf(file, "%s:%s\n", name, url)
	}
}

func loadBookmarks() {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			name := parts[0]
			url := parts[1]
			bookmarks[name] = url
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
