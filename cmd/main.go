package main

import (
	"os"

	"github.com/kaonmir/notion-to-blog/blog"
)

func main() {
	kakaoID := os.Getenv("KAKAO_ID")
	KakaoPassword := os.Getenv("KAKAO_PASSWORD")
	tistoryClientID := os.Getenv("TISTORY_CLIENT_ID")
	tistorySecretKey := os.Getenv("TISTORY_SECRET_KEY")
	tistoryBlogName := os.Getenv("TISTORY_BLOG_NAME")

	selenium := blog.NewSelenium()
	defer selenium.Close()

	selenium.TistoryLogin(kakaoID, KakaoPassword)
	code := selenium.GetAuthenticationCode(tistoryClientID, tistoryBlogName)
	accessToken := selenium.GetAccessToken(tistoryClientID, tistorySecretKey, tistoryBlogName, code) // TODO

	_ = accessToken

	// tistory := blog.NewTistory(selenium.Session, tistoryBlogName)
	// tistory.Login(tistoryClientID, tistorySecretKey)
}
