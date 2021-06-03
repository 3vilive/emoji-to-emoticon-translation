package translator

// 一个减少网络火药味的简易方案：
// 停用黄豆emoji：
// [太开心][可爱][微笑][鼓掌][可怜][悲伤][吐]😅🙄🙂
// 改用颜文字：
// σ`∀´)(`ヮ')∠( ᐛ 」∠)＿(╯°Д°）╯︵ /(.□ . \)_(:_」∠)_(´･ω･`)◡ ヽ(`Д´)ﾉ ┻━┻Σ（ﾟдﾟlll）( ´▽｀)(｡ì _ í｡)/( ゝ∀ ･)\(╬ﾟдﾟ)( ﾟ∀ﾟ)σ
// 举例：
// 今天逛漫展看到了肥宅😅
// 今天逛漫展看到了肥宅(`ヮ´ )σ`∀´) ﾟ∀ﾟ)σ
// 那你可真是好棒棒喔[太开心]
// 那你可真是好棒棒喔∠( ᐛ 」∠)＿

var mapping = map[rune]string{

	'👏': "(╯°Д°）╯︵",

	'🙄': "\\(╬ﾟдﾟ)",

	// face-smiling
	'😀': "(・∀・)",
	'😃': "(=^_^=)",
	'😄': "(´∀｀）",
	'😁': "(´w｀*)",
	'😆': "(ᗒᗊᗕ)",
	'😅': "(`ヮ´ )σ`∀´) ﾟ∀ﾟ)σ",
	'🤣': "｡゜(｀∀´)゜｡",
	'😂': "｡゜(｀∀´)゜｡",
	'🙂': "∠( ᐛ 」∠)＿",
	'🙃': "∩(´∀`∩)",
	'😉': "(ㆁᴗㆁ✿)",
	'😊': "（*´▽`*)",
	'😇': "(︶ω︶)",
}
