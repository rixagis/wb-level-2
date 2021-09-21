package wget

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// fixLinks заменяет в line все относительные ссылки, соответствующие re, на абсолютные, с добавлением хоста из base
func fixLinks(line string, base *url.URL, re *regexp.Regexp) string {
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		full, err := base.Parse(match[1])
		if err != nil {
			log.Println("error while parsing:", err)
		}
		if match[1] != full.String() {
			line = strings.Replace(line, match[1], full.String(), 1)
		}
	}
	return line
}

// Wget скачивает ресурс по ссылке urlPath и сохраняет его в файл filePath
func Wget(filePath string, urlPath string) {
	resp, err := http.Get(urlPath)
	if err != nil {
		log.Fatalf("Could not connect to %s", urlPath)
	}

	reader := bufio.NewReader(resp.Body)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Could not create file %s", filePath)
	}

	// заменяем относительные ссылки на абсолютные
	base, _ := url.Parse(urlPath)
	srcRE, err := regexp.Compile(`src=\"(.*?)\"`)
	if err != nil {
		log.Panicf("The regex is incorrect: %s", err)
	}
	hrefRE, err := regexp.Compile(`href=\"(.*?)\"`)
	if err != nil {
		log.Panicf("The regex is incorrect: %s", err)
	}

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Finished downloading file!")
				return
			}
			log.Fatal("Error while reading data:", err)
		}

		input = fixLinks(input, base, srcRE)
		input = fixLinks(input, base, hrefRE)
		

		file.WriteString(input)
	}
}

// MakeFileName делает имя файла из ссылки.
// Если ссылка корневая - функция вернет index.html.
// Если имя ресурса не имеет расширения - добалвяется .html
func MakeFileName(urlPath string) (string, error) {
	u, err := url.Parse(urlPath)
	if err != nil {
		return "", err
	}
	path := u.Path
	log.Println(path)
	if path == "" || path == "/" {
		return "index.html", nil
	}
	parts := strings.Split(urlPath, "/")
	filename := parts[len(parts) - 1]

	if !strings.Contains(filename, ".") {
		filename = filename + ".html"
	}
	return filename, err
}