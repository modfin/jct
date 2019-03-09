package jct

func Toggle(in []byte, from, to Case) ([]byte) {

	readString := func(start int) (end int, buf []byte) {

		var last byte
		for i := start + 1; i < len(in); i++ {
			end = i
			var cur = in[i]
			if cur == '"' && last != '\\' {
				break
			}
			buf = append(buf, cur)
			last = cur
		}
		return end, buf
	}

	isKey := func(start int) (iskey bool) {

		for i := start + 1; i < len(in); i++ {
			var cur = in[i]
			if cur != ' ' {
				return cur == ':'
			}
		}
		return
	}

	res := make([]byte, 0, len(in))
	for i := 0; i < len(in); i++ {
		var cur = in[i]

		res = append(res, cur)

		// starting string
		if cur == '"' {
			end, str := readString(i)

			if isKey(end) {
				str = []byte(to.Join(from.Split(string(str))))
			}

			res = append(res, str...)
			res = append(res, '"')

			i = end
		}
	}
	return res
}
