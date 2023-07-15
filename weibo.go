package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
	"net/http"
	"strings"
	"time"
)

var token = flag.String("token", "", "APP_TOKEN AT_....")
var uid = flag.String("uid", "", "UID UID_...,UID_...")
var tag = flag.String("tag", "", "tag 如：热 爆 新 ")
var cookie = flag.String("cookie", "", "COOKIE")

const HOSTNAME = "https://s.weibo.com"

var APP_TOKEN = *token
var UID = *uid
var TAG = *tag
var COOKIE = *cookie		

func init() {
	flag.Parse()
	APP_TOKEN = *token
	UID = *uid
	TAG = *tag
	COOKIE = *cookie
}

func main() {
	SendMyMessage(TAG)
	//c := cron.New()
	////每天上午11点  推送 热 标签热搜
	//c.AddFunc("0 0 11 * * ?", func() {
	//	fmt.Println("执行定时任务")
	//	SendMyMessage("热")
	//	fmt.Println("定时任务执行结束") 
	//})
	////每1小时  推送 爆 标签热搜
	//c.AddFunc("@every 1h", func() {
	//	fmt.Println("执行定时任务")
	//	SendMyMessage("爆")
	//	fmt.Println("定时任务执行结束")
	//})
	//
	//c.Start()
	//select {}
}

func SendMyMessage(typeer string) {
	client := http.Client{
    		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://s.weibo.com/top/summary?cate=realtimehot", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.79")
	req.Header.Set("Cookie", COOKIE)
	res, err := client.Do(req)
	checkErr(err)
	defer func() { _ = res.Body.Close() }()
	//body, _ := ioutil.ReadAll(res.Body)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	var str = ""
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		//herfText := s.Find(".td-02 a").Text()
		//href,_ := s.Find(".td-02 a").Attr("href")
		redu := s.Find(".td-03 i").Text()

		if redu == typeer {
			//跳过置顶推荐
			if i==0 && typeer =="热"{
				return
			}			
			href, _ := s.Find(".td-02 a").Attr("href")
			herfText := s.Find(".td-02 a").Text()
			redu := s.Find(".td-03 i").Text()
			str += fmt.Sprintf(`
            <a class="weui-cell  weui-cell_access" href="%s%s">
                <div class="weui-cell__bd">
                    <p>%s</p>
                </div>
                <div class="weui-cell__ft">%s
                </div>
            </a>`, HOSTNAME, href, herfText, redu)
		}
	})
	if len(str) > 0 {
		//存在即发送
		fmt.Println("发送消息！")
		senMessage(fmt.Sprintf(`<head><link rel="stylesheet" href="https://res.wx.qq.com/open/libs/weui/2.3.0/weui.min.css"/></head>
			 <div class="page"><div class="weui-cells">%s</div></div>
			`, str))
	}

}

func senMessage(str string) {
	arr := strings.Split(UID, ",")
	for _, val := range arr {
		//TopicId 主题id   ContentType 2 html形式    1 普通文本
		msg := model.NewMessage(APP_TOKEN).
			SetContent(str).
			AddUId(val).SetSummary("微博热搜 「"+TAG+"」 推送").
			SetUrl("https://s.weibo.com/top/summary?cate=realtimehot").
		        SetContentType(2)
		msgArr, err := wxpusher.SendMessage(msg)
		fmt.Println(msgArr, err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
