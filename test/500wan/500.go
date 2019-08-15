package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	//"strings"

	"github.com/axgle/mahonia"
)

type GameInfo struct {
	number    string
	gametype  string
	round     string
	timestamp string
	isrun     string
	rightgame string
	leftgame  string
	pankou    string
	result    string
	gameid    string
	fenxiurl  string
}

func fetch(url string) []byte {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return nil
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return nil
	}
	return body
}

func parseOneGame(body string) *GameInfo {
	regtd := regexp.MustCompile("<td .*>(.*?)</td>")
	data := regtd.FindAllStringSubmatch(body, -1)
	//fmt.Println(data)
	game := &GameInfo{}

	// for k, v := range data {
	// 	fmt.Println(k, v)
	// 	game.number = v[1]
	// 	game.gametype
	// }
	game.number = data[0][1]
	//<td bgcolor="#00A78D" class="ssbox_01"><a style="color:#fff" target="_blank" href="//liansai.500.com/zuqiu-5050/">亚冠杯</a></td>
	regbstype := regexp.MustCompile("<td .*><a .+>(.+?)</a></td>")
	bstypes := regbstype.FindAllStringSubmatch(data[1][0], -1)
	//fmt.Println(bstypes)
	game.gametype = bstypes[0][1]
	game.round = data[2][1]
	game.timestamp = data[3][1]
	game.isrun = data[4][1]
	game.result = data[8][1]
	game.fenxiurl = data[11][1]

	rgamereg := regexp.MustCompile("<td align=\"right\".+><a .+>(.+?)</a>")
	rgamedata := rgamereg.FindAllStringSubmatch(data[5][0], -1)
	game.rightgame = rgamedata[0][1]
	lgamereg := regexp.MustCompile("<td align=\"left\" .+><a .+>(.+?)</a>")
	lgamedata := lgamereg.FindAllStringSubmatch(data[7][0], -1)
	game.leftgame = lgamedata[0][1]

	pkreg := regexp.MustCompile("<a href=.+data-ao=.+>(.+?)</a><a href=")
	pkdata := pkreg.FindAllStringSubmatch(data[6][0], 1)
	//fmt.Println(pkdata[0])
	game.pankou = pkdata[0][1]

	fidreg := regexp.MustCompile(`fid="([\d]+)"`)
	fiddata := fidreg.FindAllStringSubmatch(body, -1)
	//fmt.Println(fiddata)
	game.gameid = fiddata[0][1]
	game.fenxiurl = fmt.Sprintf("http://odds.500.com/fenxi/shuju-%s.shtml", game.gameid)
	fmt.Printf("%+v\n", game)
	return game
}

/*
<tr id="a808831" order="7010" status="4" gy="荷甲,瓦尔韦克,阿尔克马" yy="荷甲,RKC華域克,艾克馬亞" lid="91" fid="808831" sid="5289" class="bg02" infoid="132677" r="1">
    <td align="center" class=""><input type="checkbox" name="check_id[]" value="808831" />周日010</td>
    <td bgcolor="#ff6699" class="ssbox_01"><a style="color:#fff" target="_blank" href="//liansai.500.com/zuqiu-5289/">荷甲</a></td>
    <td align="center">第2轮</td>
    <td align="center">08-11 20:30</td>
    <td align="center"><span class="red">完</span></td>
    <td align="right" class="p_lr01"><span class="gray">[14]</span><span class="yellowcard">2</span><span class="redcard">1</span><a target="_blank" href="//liansai.500.com/team/385/">瓦尔韦克</a><span class="sp_sr">(+1)</span></td>
    <td align="center"><div class="pk"><a href="./detail.php?fid=808831&r=1" target="_blank" class="clt1" >0</a><a href="./detail.php?fid=808831&r=1" target="_blank" class="fhuise" data-ao="受球半" data-pb="受球半">受球半</a><a href="./detail.php?fid=808831&r=1" target="_blank" class="clt3" >2</a></div></td>
    <td align="left" class="p_lr01"><a target="_blank" href="//liansai.500.com/team/224/">阿尔克马</a><span class="yellowcard">1</span><span class="gray">[05]</span></td>
    <td align="center" class="red">0 - 2</td>
    <td align="center" class="bf_op">&nbsp;</td>
    <td align="center" class="red">负 </td>
    <td align="center" class="td_warn"><a target="_blank" href="//odds.500.com/fenxi/shuju-808831.shtml">析</a> <a target="_blank" href="//odds.500.com/fenxi/yazhi-808831.shtml">亚</a> <a target="_blank" href="//odds.500.com/fenxi/ouzhi-808831.shtml">欧</a> <a target="_blank" id="qing_808831" class="red hide"   href="//odds.500.com/fenxi/youliao-808831.shtml?channel=pc_score">情</a></td>
    <td align="center" class=""><a href="javascript:void(0)" class="icon_notop">置顶</a></td>
  </tr>
*/

func parseHtml(body string) {
	//strings.Replace(body, "\n", "", -1)

	bsinforeg := regexp.MustCompile("<tr id=\"(?s:(.+?))</tr>")
	//fmt.Println(body)
	infos := bsinforeg.FindAllString(body, -1)

	for _, v := range infos {
		//fmt.Println(k, v, "\n\n\n")
		parseOneGame(v)
	}

	//fmt.Println(infos[0])
	//parseOneGame(infos[0])
}

func parseOdds(body string) {
	oddsreg := regexp.MustCompile(`var liveOddsList=(.+?);`)
	oddsdata := oddsreg.FindAllStringSubmatch(body, 1)
	fmt.Println(oddsdata[0][1])
	oddsmap := make(map[string]map[string][3]float32)

	err := json.Unmarshal([]byte(oddsdata[0][1]), &oddsmap)
	if err != nil {
		log.Panic("failed to Unmarshal ", err)
	}
	fmt.Println(oddsmap["804172"])
}

func main() {
	body := fetch("http://live.500.com")
	//fmt.Println(string(body))

	decoder := mahonia.NewDecoder("GB18030")
	utf8_body := decoder.ConvertString(string(body))
	//fmt.Println(utf8_body)
	//fmt.Println(decoder.ConvertString(string(body)))
	parseHtml(utf8_body)
	//parseOdds(utf8_body)
	return

}
