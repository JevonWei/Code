package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// 从job.html文件中筛选
// func main() {
// 	//url := "https://www.zhipin.com/job_detail"

// 	cxt, _ := ioutil.ReadFile("job.html")
// 	reader := bytes.NewReader(cxt)

// 	document, _ := goquery.NewDocumentFromReader(reader)
// 	//fmt.Println(err, document)
// 	document.Find("div.job-primary").Each(func(index int, selection *goquery.Selection) {
// 		fmt.Println("-------------------------------------------------------------------------------")
// 		//fmt.Println(selection.Html())
// 		fmt.Println(selection.Find("div.info-company > div.company-text > h3.name > a").Text())

// 		tagA := selection.Find("div.info-primary > h3.name > a")
// 		fmt.Println(tagA.Find("div.job-title").Text())
// 		fmt.Println(tagA.Find("span").Text())
// 	})
// }

// 在线从浏览器中筛选数据
func main() {
	url := "https://www.zhipin.com/job_detail"
	cook := "lastCity=101020100; sid=sem_pz_bdpc_dasou_title; __c=1565432598; __g=sem_pz_bdpc_dasou_title; __l=l=%2Fwww.zhipin.com%2F%3Fsid%3Dsem_pz_bdpc_dasou_title&r=https%3A%2F%2Fsp0.baidu.com%2F9q9JcDHa2gU2pMbgoY3K%2Fadrc.php%3Ft%3D06KL00c00fDIFkY0IWPB0KZEgsDihY4I00000Kd7ZNC00000xDt46C.THdBULP1doZA8QMu1x60UWY1rH0YnW0dr7tkPNqEuydxuAT0T1dbuWnkPH-huW0snHTznWFW0ZRqwjKDwWckwjD4P1f3fYPKfWT1fWFAPYfswWN7nYmdwj60mHdL5iuVmv-b5HnznWfvnjf1Pj6hTZFEuA-b5HDv0ARqpZwYTZnlQzqLILT8Xh9GTA-8QhPEUitOTv-b5gP-UNqsX-qBuZKWgvw9TvqdgLwGIAk-0APzm1Y4P1bz%26tpl%3Dtpl_11534_19713_15764%26l%3D1511867677%26attach%3Dlocation%253D%2526linkName%253D%2525E6%2525A0%252587%2525E5%252587%252586%2525E5%2525A4%2525B4%2525E9%252583%2525A8-%2525E6%2525A0%252587%2525E9%2525A2%252598-%2525E4%2525B8%2525BB%2525E6%2525A0%252587%2525E9%2525A2%252598%2526linkText%253DBoss%2525E7%25259B%2525B4%2525E8%252581%252598%2525E2%252580%252594%2525E2%252580%252594%2525E6%252589%2525BE%2525E5%2525B7%2525A5%2525E4%2525BD%25259C%2525EF%2525BC%25258C%2525E6%252588%252591%2525E8%2525A6%252581%2525E8%2525B7%25259F%2525E8%252580%252581%2525E6%25259D%2525BF%2525E8%2525B0%252588%2525EF%2525BC%252581%2526xp%253Did(%252522m3224604348_canvas%252522)%25252FDIV%25255B1%25255D%25252FDIV%25255B1%25255D%25252FDIV%25255B1%25255D%25252FDIV%25255B1%25255D%25252FDIV%25255B1%25255D%25252FH2%25255B1%25255D%25252FA%25255B1%25255D%2526linkType%253D%2526checksum%253D8%26ie%3Dutf-8%26f%3D3%26tn%3D39042058_15_oem_dg%26wd%3Dboss%25E7%259B%25B4%25E8%2581%2598%25E5%25AE%2598%25E7%25BD%2591%26oq%3D%2525E5%25259B%2525BE%2525E7%252589%252587%26rqlang%3Dcn%26inputT%3D3139%26prefixsug%3Dboss%26rsp%3D1&g=%2Fwww.zhipin.com%2F%3Fsid%3Dsem_pz_bdpc_dasou_title; _uab_collina=156543259837117816175265; Hm_lvt_194df3105ad7148dcf2b98a91b5e727a=1565432598; __zp_stoken__=44fcDCEoS3Fj9n1kX4MABEDkDd%2FV%2Fab%2BycBfDDTUESDxOPaRczeLYsFJSfCTwQaPc7tZjsuQ6kjTtHNk5F3zCRPnmg%3D%3D; Hm_lpvt_194df3105ad7148dcf2b98a91b5e727a=1565434198; __a=40168454.1565432598..1565432598.6.1.6.6"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("cookie", cook)

	client := &http.Client{}

	response, _ := client.Do(request)
	document, _ := goquery.NewDocumentFromResponse(response)

	document.Find("div.job-primary").Each(func(index int, selection *goquery.Selection) {
		fmt.Println("-------------------------------------------------------------------------------")
		//fmt.Println(selection.Html())
		fmt.Println(selection.Find("div.info-company > div.company-text > h3.name > a").Text())

		tagA := selection.Find("div.info-primary > h3.name > a")
		fmt.Println(tagA.Find("div.job-title").Text())
		fmt.Println(tagA.Find("span").Text())
	})
}
