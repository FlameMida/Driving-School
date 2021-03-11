package utils

import (
	"Driving-school/global"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/gurkankaymak/gosoup"
	"net/http"
	"regexp"
	"time"
)

type SongList struct {
	Num    int    `json:"num"`
	Name   string `json:"name"`
	Id     string `json:"id"`
	Url    string `json:"url"`
	Cover  string `json:"cover"`
	Artist string `json:"artist"`
}

func GetTopList() (List []SongList, err error) {
	ListId := global.GVA_CONFIG.Music.ListId
	url := fmt.Sprintf("http://music.163.com/discover/toplist?id=%s", ListId)
	KeyName := fmt.Sprintf("%s-List-%s-Music", carbon.Now().ToDateString(), ListId)
	if global.GVA_CONFIG.System.UseMultipoint {
		if global.GVA_REDIS.Exists(KeyName).Val() > 0 {
			songStr := global.GVA_REDIS.Get(KeyName).Val()
			_ = json.Unmarshal([]byte(songStr), &List)
			return List, nil
		}
	}
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
	length := global.GVA_CONFIG.Music.Length
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
		List = append(List, SongList{num, songName, id, src, pic, artist})
		if count == length {
			break
		}
		count++
	}
	if global.GVA_CONFIG.System.UseMultipoint {
		rdsVal, _ := json.Marshal(List)
		global.GVA_REDIS.Set(KeyName, rdsVal, time.Hour*24)
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
	text = ProcessBody(res)
	return text, nil

}

func ProcessBody(response *http.Response) string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(response.Body)
	return buf.String()
}
