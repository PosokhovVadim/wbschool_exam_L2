package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var r bool

var visited = make(map[string]bool)

var allowDomains []string

func downloadFile(filepath string, url string) error {

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get: ", err)
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func makePath(baseURL string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	dir := path.Join("www", u.Host, path.Dir(u.Path))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return "", err
		}
	}

	if strings.HasSuffix(u.Path, "/") || path.Ext(u.Path) == "" {
		return path.Join(dir, "index.html"), nil
	}

	return path.Join(dir, path.Base(u.Path)), nil
}

func setAllowedDomain(baseURL string) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	allowDomains = append(allowDomains, u.Host)
	fmt.Println(allowDomains)
	return nil

}

func isAllowedDomain(link string) bool {
	for _, v := range allowDomains {
		if v == link {
			return true
		}
	}
	return false
}

func absoluteURL(baseURL string, link string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	base := u
	u, err = url.Parse(link)
	if err != nil {
		return ""
	}

	resolved := base.ResolveReference(u)
	if !isAllowedDomain(resolved.Host) {
		return ""
	}

	return resolved.String()
}

func scrapper(currURL string) error {
	if visited[currURL] {
		return nil
	}
	visited[currURL] = true

	resp, err := http.Get(currURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Get the URL: ", currURL)

	filePath, err := makePath(currURL)
	if err != nil {
		return err
	}

	err = downloadFile(filePath, currURL)
	if err != nil {
		return err
	}

	body := resp.Body
	doc, err := html.Parse(body)
	if err != nil {
		return err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "href" {
					absURL := absoluteURL(currURL, a.Val)

					if absURL == "" {
						return
					}

					err := scrapper(absURL)
					if err != nil {
						fmt.Println(err)
					}
				} else if a.Key == "src" {

					absURL := absoluteURL(currURL, a.Val)
					if absURL == "" {
						return
					}

					filePath, err := makePath(absURL)
					if err != nil {
						fmt.Println(err)
					}

					err = downloadFile(filePath, absURL)
					if err != nil {

						fmt.Println(err)
					}
				}

			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return nil
}

// TODO: add max depth
func proccess(baseURL string) error {
	path, err := makePath(baseURL)
	if err != nil {
		return err
	}

	setAllowedDomain(baseURL)

	if r {
		err := scrapper(baseURL)
		if err != nil {
			return err
		}
	} else {
		err := downloadFile(path, baseURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func run() error {
	flag.BoolVar(&r, "r", false, "recursive")

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		return fmt.Errorf("expected url, got %v", args)
	}

	if err := proccess(args[0]); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
