package mangadex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mmcdole/gofeed"
)

type Mangadex struct {
	client *http.Client
}

func Initilize() *Mangadex {
	mangadex := Mangadex{
		client: http.DefaultClient,
	}
	return &mangadex
}

func (mangadex *Mangadex) Latest(token string) ([]string, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://mangadex.org/rss/" + token)
	if err != nil {
		return nil, err
	}

	var Links []string
	for _, item := range feed.Items {
		Links = append(Links, item.Link)
	}

	return Links, nil
}

func (mangadex *Mangadex) GetInfo(id string) (*Manga, error) {
	url := "https://mangadex.org/api/manga/" + id

	res, err := mangadex.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get %s: %s", url, res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)

	var resp response

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	var Chapters []Chapter

	for key, element := range resp.Chapter {
		element.ID, _ = strconv.ParseInt(key, 10, 32)
		Chapters = append(Chapters, element)
	}

	resp.Manga.Chapters = Chapters

	return &resp.Manga, err
}

func (mangadex *Mangadex) RetrieveImageLinks(id int64) (*Chapter, error) {
	url := "https://mangadex.org/api/?type=chapter&id=" + strconv.FormatInt(id, 10)

	res, err := mangadex.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get %s: %s", url, res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)

	var chapter Chapter
	err = json.Unmarshal(data, &chapter)
	if err != nil {
		return nil, err
	}

	var Links []string
	for _, element := range chapter.Links {
		link := chapter.Server + chapter.Hash + "/" + element
		Links = append(Links, link)
	}

	chapter.Links = Links

	return &chapter, err
}
