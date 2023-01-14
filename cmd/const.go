package main

type Column struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Tag      string `json:"tag"`
	Status   string `json:"status"`
	Url      string `json:"url"`
}

const (
	UPLOAD_VALUE   = "발행 요청"
	MODIFY_VALUE   = "수정 요청"
	COMPLETE_VALUE = "발행 완료"
)

var (
	COLUMN = Column{
		Title:    "제목",
		Category: "카테고리",
		Tag:      "태그",
		Status:   "상태",
		Url:      "링크",
	}
)
