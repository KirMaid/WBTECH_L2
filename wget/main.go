package wget

import (
	"fmt"
	//"github.com/cavaliercoder/grab"
	"log"
	_ "net/http"
	"os"
	_ "path/filepath"
)

func main() {
	// Проверяем аргументы командной строки
	if len(os.Args) < 2 {
		log.Fatal("Укажите URL для скачивания")
	}
	url := os.Args[1]

	// Создаем директорию для сохранения файлов, если она не существует
	dir := "./site_downloads"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	// Скачиваем файл
	client := grab.NewClient()
	req, err := grab.NewRequest(dir, url)
	if err != nil {
		log.Fatal(err)
	}

	// Запускаем скачивание
	resp := client.Do(req)
	fmt.Printf("Скачивание %s...\n", req.URL())

	// Ждем завершения скачивания
	t := resp.Start()
	<-t.Done()

	// Обрабатываем результат
	if err := t.Error(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Файл сохранен как %s\n", resp.Filename)
	}
}
