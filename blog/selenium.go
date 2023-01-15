package blog

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/fedesog/webdriver"
)

type Selenium struct {
	driver  *webdriver.ChromeDriver
	Session *webdriver.Session
}

func NewSelenium() *Selenium {
	driver := webdriver.NewChromeDriver("../chromedriver")
	driver.Start()
	// defer driver.Stop()

	desired := webdriver.Capabilities{"Platform": "macos"}
	required := webdriver.Capabilities{}

	session, _ := driver.NewSession(desired, required)
	// defer session.Delete()

	return &Selenium{driver: driver, Session: session}
}

func (s *Selenium) Close() {
	s.Session.Delete()
	s.driver.Stop()
}

func (s *Selenium) TistoryLogin(KakaoID string, KakaoPassword string) error {
	s.Session.Url("https://accounts.kakao.com/login?continue=https%3A%2F%2Fkauth.kakao.com%2Foauth%2Fauthorize%3Fis_popup%3Dfalse%26ka%3Dsdk%252F1.43.0%2520os%252Fjavascript%2520sdk_type%252Fjavascript%2520lang%252Fen-US%2520device%252FMacIntel%2520origin%252Fhttps%25253A%25252F%25252Fwww.tistory.com%26auth_tran_id%3Dn11fn9p740o3e6ddd834b023f24221217e370daed18l9ms8up0%26response_type%3Dcode%26state%3DaHR0cHM6Ly93d3cudGlzdG9yeS5jb20v%26redirect_uri%3Dhttps%253A%252F%252Fwww.tistory.com%252Fauth%252Fkakao%252Fredirect%26through_account%3Dtrue%26client_id%3D3e6ddd834b023f24221217e370daed18&talk_login=hidden")
	time.Sleep(3 * time.Second)

	loginKey, _ := s.Session.FindElement(webdriver.ID, "input-loginKey")
	loginKey.SendKeys(KakaoID)
	time.Sleep(3 * time.Second)

	passwordKey, _ := s.Session.FindElement(webdriver.ID, "input-password")
	passwordKey.SendKeys(KakaoPassword)
	time.Sleep(3 * time.Second)

	button, _ := s.Session.FindElement(webdriver.XPath, `//*[@id="mainContent"]/div/div/form/div[4]/button[1]`)
	button.Click()
	time.Sleep(3 * time.Second)

	return nil
}

func (s *Selenium) GetAuthenticationCode(clientID string, blogName string) string {
	url := fmt.Sprint(
		"https://www.tistory.com/oauth/authorize?",
		"client_id=", clientID,
		"&redirect_uri=", fmt.Sprintf(`https://%s.tistory.com`, blogName),
		"&response_type=code",
		"&state=1234",
	)

	s.Session.Url(url)
	html, _ := s.Session.Source()
	r := regexp.MustCompile(`https://kaonmir.tistory.com\?code=(.*?)&state=(.*)`)
	// match := r.FindString(html)
	matches := r.FindStringSubmatch(html)
	authenticationCode := matches[1]
	return authenticationCode
}

// TODO: TODO TODO TODO WILLBEDONEIBELEIVEIWISH
func (s *Selenium) GetAccessToken(clientID string, secretKey string, blogName string, code string) string {
	url := fmt.Sprint(
		"https://www.tistory.com/oauth/access_token?",
		"client_id=", clientID,
		"&client_secret=", secretKey,
		"&redirect_uri=", fmt.Sprintf(`https://%s.tistory.com`, blogName),
		"&code=", code,
		"&grant_type=authorization_code",
	)

	// Get URL http
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// Read response b
	b, _ := ioutil.ReadAll(resp.Body)
	accessToekn := strings.Split(string(b), "=")[1]
	return accessToekn
}

// url := "https://www.tistory.com/oauth/access_token"
// body := map[string]string{
// 	"client_id":     clientID,
// 	"client_secret": secretKey,
// 	"redirect_uri":  t.redirectURI,
// 	"code":          "code",
// 	"grant_type":    "authorization_code",
// }
