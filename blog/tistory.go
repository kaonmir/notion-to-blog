package blog

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/fatih/structs"
	"github.com/fedesog/webdriver"
)

type PostBody struct {
	Title    string
	Content  string
	Category string
	Tag      string
}

type Tistory struct {
	session  *webdriver.Session
	token    string
	blogName string
}

func NewTistory(session *webdriver.Session, blogName string) *Tistory {
	return &Tistory{
		session: session,
		// token:    token,
		blogName: blogName,
	}
}

// kind : write, modify
func (t *Tistory) Post(kind string, postBody PostBody, modifyId string) (string, error) {
	// blogName: Blog Name (필수)
	// title: 글 제목 (필수)
	// content: 글 내용
	// visibility: 발행상태 (0: 비공개 - 기본값, 1: 보호, 3: 발행)
	// category: 카테고리 아이디 (기본값: 0)
	// published: 발행시간 (TIMESTAMP 이며 미래의 시간을 넣을 경우 예약. 기본값: 현재시간)
	// tag: 태그 (',' 로 구분)

	body := structs.Map(postBody)
	body["access_token"] = t.token
	body["blogName"] = t.blogName
	body["output"] = "xml"
	body["visibility"] = 3
	body["postId"] = modifyId

	var url string
	if kind == "write" {
		url = "https://www.tistory.com/apis/post/write"
	} else {
		url = "https://www.tistory.com/apis/post/modify"
	}

	pbytes, _ := json.Marshal(body)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(pbytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// TODO : Return blog url
	return "", nil
}

// https://tistory.github.io/document-tistory-apis/apis/v1/post/attach.html
func (t *Tistory) Attach(name string, data []byte) error {
	// decode base64
	// sDec := base64.StdEncoding.DecodeString(url)

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	multi, err := writer.CreateFormFile("uploadedfile", name) // 업로드할 파일 이름
	if err != nil {
		return err
	}

	// Read image to Writer
	io.Copy(multi, bytes.NewReader(data))

	http.NewRequest("POST", "https://www.tistory.com/apis/post/attach", buf)

	return nil
}
