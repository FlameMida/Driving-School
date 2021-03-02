package utils

import (
	"bytes"
	"errors"
	"github.com/gurkankaymak/gosoup"
	"net/http"
	"regexp"
)

type TopList struct {
	Num    int    `json:"num"`
	Title  string `json:"title"`
	Id     string `json:"id"`
	Src    string `json:"src"`
	Pic    string `json:"pic"`
	Artist string `json:"artist"`
}

func GetTopList() (List []TopList, err error) {
	url := "http://music.163.com/discover/toplist"

	//处理返回结果
	text, err := requests(url)
	if err != nil {
		return nil, err
	}

	element, err := gosoup.ParseAsHTML(text)
	if err != nil {
		// log/handle error
		return List, err
	}
	musicListRaw := element.Find("ul", gosoup.Attributes{"class": "f-hide"})
	if musicListRaw == nil {
		return List, errors.New(" No Songs ")
	}
	musicList := musicListRaw.FindAllByTag("a")
	count := 1
	for _, e := range musicList {

		songName := e.LastChild.Data
		href, _ := e.GetAttribute("href")
		reg := regexp.MustCompile(`[0-9]+`)
		id := reg.FindAllString(href, -1)[0]
		num := count
		picUrl := "http://music.163.com/song?id=" + id

		res, _ := http.Get(picUrl)
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(res.Body)
		text = buf.String()

		picelement, _ := gosoup.ParseAsHTML(text)
		picRaw := picelement.Find("div", gosoup.Attributes{"class": "u-cover u-cover-6 f-fl"})
		picNode := picRaw.FindByTag("img")
		pic := picNode.Attr[2].Val
		artistRaw := picelement.Find("a", gosoup.Attributes{"class": "s-fc7"})
		artist := artistRaw.LastChild.Data
		src := "http://music.163.com/song/media/outer/url?id=" + id + ".mp3"
		List = append(List, TopList{num, songName, id, src, pic, artist})
		if count == 10 {
			break
		}
		count++
	}
	return List, nil

}

func requests(url string) (text string, err error) {
	client := &http.Client{}
	//提交请求
	requests, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	//增加header选项
	requests.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	res, _ := client.Do(requests)
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(res.Body)
	text = buf.String()
	return text, nil

}
