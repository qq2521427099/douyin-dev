package conn

import (
	"fmt"
	"github.com/syyongx/go-wordsfilter"
)

var WordsFilter *wordsfilter.WordsFilter

var Root map[string]*wordsfilter.Node

func InitWordsFilter() {
	WordsFilter = wordsfilter.New()
	var err error
	Root, err = WordsFilter.GenerateWithFile("resources/sensitive_words/sensitive_words.txt")
	if err != nil {
		fmt.Println("初始化敏感词过滤器失败")

	} else {
		fmt.Println("初始化敏感词过滤器成功!!!")
	}
}
