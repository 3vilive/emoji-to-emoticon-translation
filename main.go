package main

import (
	"fmt"

	"github.com/3vilive/emoji-to-emoticon-translation/translator"
)

func main() {
	cases := []string{
		"今天逛漫展看到了肥宅😅",
		"那你可真是好棒棒喔🙂",
		"都不知道说啥好🙄",
		"你说的都对😄",
		"你说的真有道理😊",
		"祝你身体健康😃",
		"哇，你真有两把刷子😁",
		"我都不知道要说啥了😂",
		"认真的吗🙃",
		"祝你好运😇",
	}

	for _, text := range cases {
		translated, err := translator.Translate(text)
		if err != nil {
			panic(err)
		}

		fmt.Printf("text: %s\ntranslated: %s\n\n", text, translated)
	}

}
