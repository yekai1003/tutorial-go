/*
	爬虫 分析500.com的数据
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/axgle/mahonia"
)

//go get -u github.com/axgle/mahonia

func fetchUrl(url string) string {
	cli := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//做请求
	response, err := cli.Do(req)
	if err != nil {
		log.Panic("failed to request ", err)
	}
	//处理响应结果
	if response.StatusCode != 200 {
		log.Panic("err statuscode ")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("failed to ReadAll ", err)
	}
	//当前属于gbk编码，go语言环境是utf-8编码
	decoder := mahonia.NewDecoder("GB18030")
	utf8_body := decoder.ConvertString(string(body))
	return utf8_body

}

type GameInfo struct {
	GameType  string
	GameTime  string
	RightClub string
	LeftClub  string
	Fid       string
	odds      [3]float32
}

var globalodds map[string]map[string][3]float32

func parseGame(body string) *GameInfo {
	gameInfo := &GameInfo{}
	typereg := regexp.MustCompile(`<td bgcolor=".+>(.+?)</a></td>`)
	gametypes := typereg.FindAllStringSubmatch(body, 1)
	//fmt.Println(gametypes[0][1]) //得到比赛类型
	gameInfo.GameType = gametypes[0][1]

	//获取比赛时间 <td align="center">08-16 00:45</td>
	timereg := regexp.MustCompile(`<td align="center">(.+?)</td>`)
	times := timereg.FindAllStringSubmatch(body, -1)
	//fmt.Println(times)
	gameInfo.GameTime = times[1][1]
	//<td align="right" class="p_lr01"><span class="gray"></span><a target="_blank" href="//liansai.500.com/team/1114/">莫斯巴达</a><span class="sp_rq">(-1)</span></td>
	rightreg := regexp.MustCompile(`<td align="right".+>(.+?)</a>.+</td>`)
	rights := rightreg.FindAllStringSubmatch(body, -1)
	gameInfo.RightClub = rights[0][1]
	//客队
	leftreg := regexp.MustCompile(`<td align="left".+>(.+?)</a>.+</td>`)
	lefts := leftreg.FindAllStringSubmatch(body, -1)
	gameInfo.LeftClub = lefts[0][1]
	//获取fid 比赛唯一标识
	fidreg := regexp.MustCompile(`fid="([0-9]+)"`)
	fids := fidreg.FindAllStringSubmatch(body, -1)
	//fmt.Println(figs)
	gameInfo.Fid = fids[0][1]
	gameInfo.odds = globalodds[gameInfo.Fid]["5"]
	fmt.Println(gameInfo)
	return gameInfo
}

//解析网页
/*
<tr id="a855640" order="4001" status="0" gy="欧罗巴,莫斯巴达,图恩" yy="欧罗巴,莫斯科斯巴達,图恩" lid="63" fid="855640" sid="5295" class="" infoid="132788" r="1">
    <td align="center" class=""><input type="checkbox" name="check_id[]" value="855640" />周四001</td>
    <td bgcolor="#6F00DD" class="ssbox_01"><a style="color:#fff" target="_blank" href="//liansai.500.com/zuqiu-5295/">欧罗巴</a></td>
    <td align="center">资格赛3</td>
    <td align="center">08-16 00:45</td>
    <td align="center">未</td>
    <td align="right" class="p_lr01"><span class="gray"></span><a target="_blank" href="//liansai.500.com/team/1114/">莫斯巴达</a><span class="sp_rq">(-1)</span></td>
    <td align="center"><div class="pk"><a href="./detail.php?fid=855640&r=1" target="_blank" class="clt1" ></a><a href="./detail.php?fid=855640&r=1" target="_blank" class="fgreen" data-ao="一球" data-pb="一球">一球</a><a href="./detail.php?fid=855640&r=1" target="_blank" class="clt3" ></a></div></td>
    <td align="left" class="p_lr01"><a target="_blank" href="//liansai.500.com/team/777/">图恩</a><span class="gray"></span></td>
    <td align="center" class="red"> - </td>
    <td align="center" class="bf_op">&nbsp;</td>
    <td align="center" class="red">&nbsp; <a href="./detail.php?fid=855640" class="live_animate" title="动画" target="_blank"></a></td>
    <td align="center" class="td_warn"><a target="_blank" href="//odds.500.com/fenxi/shuju-855640.shtml">析</a> <a target="_blank" href="//odds.500.com/fenxi/yazhi-855640.shtml">亚</a> <a target="_blank" href="//odds.500.com/fenxi/ouzhi-855640.shtml">欧</a> <a target="_blank" id="qing_855640" class="red hide"   href="//odds.500.com/fenxi/youliao-855640.shtml?channel=pc_score">情</a></td>
    <td align="center" class=""><a href="javascript:void(0)" class="icon_notop">置顶</a></td>
  </tr>

*/
func parseHtml(body string) {
	//正则表达式如何写
	//go语言如何处理
	gamereg := regexp.MustCompile("<tr id=\"(?s:(.+?))</tr>")
	//查找符合条件的代码段 FindAllString 只要匹配，就返回整个匹配结果
	games := gamereg.FindAllString(body, -1)
	//fmt.Println(games)
	for _, v := range games {
		//fmt.Println(k, v, "\n\n\n")
		parseGame(v)
		//go parseereryGame()
	}
}

/*
{"808835":{"0":[1.35,5.05,7.38],"280":[1.33,4.95,6.7]},"809489":{"0":[2.67,2.9,2.79],"3":[2.75,2.87,2.87]}}
*/
//解析足彩赔率
func parseOdds(body string) {
	oddsreg := regexp.MustCompile(`var liveOddsList=(.+?);`)
	oddsdata := oddsreg.FindAllStringSubmatch(body, 1)
	//fmt.Println(oddsdata[0][1])
	//字符串转json
	//定义map
	globalodds = make(map[string]map[string][3]float32)
	err := json.Unmarshal([]byte(oddsdata[0][1]), &globalodds)
	if err != nil {
		log.Panic("failed to unmarshal json ", err)
	}
	//globalodds = odds
	fmt.Println(globalodds["808835"]) //"5":[1.5,4.1,5.0]
}

func main() {
	body := fetchUrl("http://live.500.com")
	//fmt.Println(body)
	parseOdds(body)
	parseHtml(body)

}
