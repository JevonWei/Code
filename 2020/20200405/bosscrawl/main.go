package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//url := "https://www.zhipin.com/job_detail/?query=go&city=101020100&industry=&position="

	// request, _ := http.NewRequest("GET", url, nil)
	// request.Header.Set("cookie", "lastCity=101020100; _uab_collina=156543259837117816175265; _bl_uid=6skqe1mhsjwmvyukbaL00zFgUpOp; __zp_seo_uuid__=dc491cd0-46e4-4b10-9f63-c6eaf618e332; Hm_lvt_194df3105ad7148dcf2b98a91b5e727a=1586172573; __c=1586172573; __g=-; __l=l=%2Fwww.zhipin.com%2Fshanghai%2F&r=https%3A%2F%2Fcn.bing.com%2F&friend_source=0&friend_source=0; Hm_lpvt_194df3105ad7148dcf2b98a91b5e727a=1586174818; __zp_stoken__=d3d3VbHJ9gUS8vDIJ0eZYEFgUYzATjeQIRiwT6tfwwcB41vaVex%2FTtF3DKeYQ%2Fqad5Jl4N5pe%2Fz2leHOwZIf7sn8fKYiofymUGicTP4%2FdRUc67hsQWl7cLnW5m15veXOLXB%2F; __a=89473349.1571214462.1575964305.1586172573.941.31.5.118")
	// client := &http.Client{}
	// response, _ := client.Do(request)

	// document, _ := goquery.NewDocumentFromResponse(response)

	cxt, _ := ioutil.ReadFile("job.html")
	reader := bytes.NewReader(cxt)
	document, _ := goquery.NewDocumentFromReader(reader)
	// fmt.Println(document, err)

	document.Find("div.job-primary").Each(func(index int, selection *goquery.Selection) {
		// fmt.Println(selection.Html())
		fmt.Println(selection.Find("div.info-primary > div.info-company > div.company-text > h3 > a").Text())

		tagA := selection.Find("div.info-primary > div.primary-wrapper > div.primary-box")
		fmt.Println(tagA.Find("div.job-title > span.job-name").Text())
		fmt.Println(tagA.Find("div.job-limit > span").Text())
		fmt.Println()
	})
}
