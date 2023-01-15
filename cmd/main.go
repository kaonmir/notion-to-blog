package main

import (
	"os"

	"github.com/kaonmir/notion-to-blog/blog"
)

func main() {
	// kakaoID := os.Getenv("KAKAO_ID")
	// KakaoPassword := os.Getenv("KAKAO_PASSWORD")
	// tistoryClientID := os.Getenv("TISTORY_CLIENT_ID")
	// tistorySecretKey := os.Getenv("TISTORY_SECRET_KEY")
	// tistoryBlogName := os.Getenv("TISTORY_BLOG_NAME")

	notionSecretKey := os.Getenv("NOTIONO_SECRET_KEY")
	notionDatabaseURL := os.Getenv("NOTION_DATABASE_URL")

	// Kakao Login and get Access Token
	// selenium := blog.NewSelenium()
	// defer selenium.Close()

	// selenium.TistoryLogin(kakaoID, KakaoPassword)
	// code := selenium.GetAuthenticationCode(tistoryClientID, tistoryBlogName)
	// accessToken := selenium.GetAccessToken(tistoryClientID, tistorySecretKey, tistoryBlogName, code) // TODO

	// Notion
	notion := blog.NewNotion(notionSecretKey)
	pages := notion.GetPostablePages(notionDatabaseURL)
	notion.ExportPage(pages[0])

	// Tistory
	// tistory := blog.NewTistory(accessToken, tistoryBlogName)

	_ = notion
}
