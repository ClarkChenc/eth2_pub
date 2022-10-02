package htmlparser

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	// fmt.Println("options time used: ", time.Now().Unix()-start.Unix())

	c, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	chromeCtx, cancel := chromedp.NewContext(c)
	defer cancel()

	// fmt.Println("context: ", time.Now().Unix()-start.Unix())

	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)
	// fmt.Println("run: ", time.Now().Unix()-start.Unix())

	timeoutCtx, timeOutCancel := context.WithTimeout(chromeCtx, 10*time.Second)
	defer timeOutCancel()

	// fmt.Println("time used: ", time.Now().Unix()-start.Unix())
	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}

func GetSpecialData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var str string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = selection.Text()
	})
	return str, nil
}

func GetHTMLElement(url string, selector string) string {
	param := `document.querySelector("body")`
	html, err := GetHttpHtmlContent(url, selector, param)
	if err != nil {
		fmt.Println(err)
	}

	res, _ := GetSpecialData(html, selector)
	return res
}

func ReadLine(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				return result, nil
			}
			return nil, err
		}
		result = append(result, line)
	}
	return result, nil
}

func GetETH2AccountBalance() {
	accountFilePath := "eth2.txt"

	accountList, err := ReadLine(accountFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("account size:", len(accountList))

	for _, account := range accountList {
		url := "https://beaconscan.com/main/validator/" + account
		selector := "#ContentPlaceHolder1_divOverviewCard > div:nth-child(3) > div.col-md-9.js-focus-state.font-size-1"
		balanceStr := GetHTMLElement(url, selector)
		balanceStr = strings.Trim(balanceStr, "\n")

		fmt.Printf("%s\t%s\n", account, balanceStr)
	}

	fmt.Println("end")
}
