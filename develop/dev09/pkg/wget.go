package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Wget struct {
	c *colly.Collector
}

func NewWget() *Wget {
	return &Wget{
		c: colly.NewCollector(),
	}

}

func (w *Wget) GetSite(url string) error {
	w.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/") {
			link = url + link
		}
		if strings.HasPrefix(link, "http") || strings.HasPrefix(link, "https") {
			downloadWeb(link)
		}
	})

	err := w.c.Visit(url)
	if err != nil {
		return err
	}

	return nil
}

func downloadWeb(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	filename := getFilenameFromURL(url)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func getFilenameFromURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
