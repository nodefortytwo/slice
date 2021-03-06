package slice

type String []string

func (ss String) Index(t string) int {
	for i, v := range ss {
		if v == t {
			return i
		}
	}

	return -1
}

func (ss String) Includes(t string) bool {
	return ss.Index(t) >= 0
}

func (ss String) Next(i int) int {
	i++
	if i >= len(ss) {
		return 0
	}

	return i
}

func (ss String) NextValue(i int) string {
	return ss[ss.Next(i)]
}

func (ss String) Len() int           { return len(ss) }
func (ss String) Less(i, j int) bool { return ss[i] < ss[j] }
func (ss String) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }

func (ss String) Any(f func(string) bool) bool {
	for _, v := range ss {
		if f(v) {
			return true
		}
	}

	return false
}

func (ss String) All(f func(string) bool) bool {
	for _, v := range ss {
		if !f(v) {
			return false
		}
	}

	return true
}

func (ss String) Filter(f func(string) bool) String {
	vsf := make(String, 0)

	for _, v := range ss {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func (ss String) Map(f func(string) string) String {
	vsm := make([]string, len(ss))
	for i, v := range ss {
		vsm[i] = f(v)
	}

	return vsm
}

func (ss String) Unique() String{
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range ss {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (ss String) Slice() []string {
	return ss
}
