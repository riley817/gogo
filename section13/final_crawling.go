package main

import (
	 "bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// 최종 실습 예제
// 대상 사이트 : 루리웹 (ruliweb.com)

// 스크리팽 대상 URL
const (
	urlRoot = "https://www.ruliweb.com/"
)

// 첫 번째 방문 (메인페이지) 대상으로 원하는 URL을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	// must check for nil values
	if n.DataAtom == atom.A && n.Parent != nil {
		fmt.Println(scrape.Attr(n.Parent, "class"))
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

// 에러 체크 공통함수
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// Url 대상이 되는 페이지(서브페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContent(url string, fn string) {

	fmt.Println(fn)
	// 작업 종료 알림
	defer wg.Done()

	// Get 방식 요청
	resp, err := http.Get(url)
	errCheck(err)

	// 요청 body 닫기
	defer resp.Body.Close()

	// 응답 데이터(HTML)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	// Response 데이터(html)의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(n, "class") == "deco"
	}

	// 파일 스트림 생성 -> 파일명, 옵션, 권한
	file, err := os.OpenFile("src/section13/scrape/" + fn + ".txt",  os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	// 메소드 종료시 파일 닫기
	defer file.Close()

	// 쓰기 버퍼 선언
	w := bufio.NewWriter(file)

	// matchNode 메소드를 사용해서 원하는 노드 순회(Iterator) 하면서 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		// Url 및 해당 데이터 출력
		fmt.Println("Result : ", scrape.Text(g))
		// 파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}
	w.Flush()
}

// 동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

func main() {
	// 메인 페이지 Get 방식 요청
	resp, err := http.Get(urlRoot) // response(응답), request(요청)
	errCheck(err)

	// 요청 Body
	defer resp.Body.Close()

	// 응답 데이터 (html)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	// ParseMainNodes 메소드를 크롤링(스크래핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)
	fmt.Println(urlList)

	for _, link := range urlList {
		// 대상 URL 1차 출력
		// fmt.Println(" Target Url : ", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)
		// fmt.Println("FileName : ", fileName)

		// 작업 대기열에 추가
		wg.Add(1)

		// 고루틴 시작 -> 작업 대기열 개수와 같아야 함
		go scrapContent(scrape.Attr(link, "href"), fileName)
	}

	// 모든 작업이 종료 될 때 까지 대기
	wg.Wait()
}
