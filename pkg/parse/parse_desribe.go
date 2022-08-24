package parse

//func XmlParse() {
//	paragraphs := []string{"技术领域", "技术领域段1", "技术领域段2", "技术背景", "技术背景段2"}
//	keys := []string{"技术领域", "背景技术", "技术背景"}
//
//	res := make(map[string]any)
//
//	var buf, currentKey string
//
//	if len(paragraphs) < 1 {
//		return
//	}
//
//	currentKey = paragraphs[0]
//
//	for _, p := range paragraphs {
//		p = strings.TrimSpace(p)
//		if IsKey(p, keys) {
//			// 遇到 key 且之前的 buf 不为空就把前面的段落存起来
//			if buf != "" {
//				res[currentKey] = buf
//			}
//			currentKey = p
//			buf = ""
//		} else {
//			buf += "\n" + p
//		}
//	}
//	res[currentKey] = buf
//
//	for k, v := range res {
//		fmt.Printf("%s:%s\n", k, v)
//	}
//
//}
//
//func IsKey(s string, keys []string) bool {
//	set := mapset.NewSet[string](keys...)
//	return set.Contains(s)
//}
