//Redirect strategy base on user's location.
package strategy

import "net/http"

type LocationStrategy struct{}

func NewLocationStrategy() (sty *LocationStrategy) {
	sty = &LocationStrategy{}
	return
}

func (sty *LocationStrategy) Init() {

}

func (sty *LocationStrategy) Redirect(w http.ResponseWriter, r *http.Request) {

}

//新浪的ip api服务
//http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip=116.6.79.87
//{
//"ret": 1,
//"start": -1,
//"end": -1,
//"country": "中国",
//"province": "广东",
//"city": "东莞",
//"district": "",
//"isp": "",
//"type": "",
//"desc": ""
//}

//淘宝ip api服务
//http://ip.taobao.com/service/getIpInfo.php?ip=114.114.114.114
//{
//"code": 0,
//"data": {
//"country": "中国",
//"country_id": "CN",
//"area": "华东",
//"area_id": "300000",
//"region": "江苏省",
//"region_id": "320000",
//"city": "南京市",
//"city_id": "320100",
//"county": "",
//"county_id": "-1",
//"isp": "",
//"isp_id": "-1",
//"ip": "114.114.114.114"
//}
//}
