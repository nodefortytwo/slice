package slice

type StrSlice []string

func (ss StrSlice) Index(t string) int {
    for i, v := range ss {
        if v == t {
            return i
        }
    }
    return -1
}

func (ss StrSlice) Contains(t string) bool {
    return ss.Index(t) >= 0
}

func (ss StrSlice) Any(f func(string) bool) bool {
    for _, v := range ss {
        if f(v) {
            return true
        }
    }
    return false
}

func (ss StrSlice) All(f func(string) bool) bool {
    for _, v := range ss {
        if !f(v) {
            return false
        }
    }
    return true
}

func (ss StrSlice) Filter(f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range ss {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func (ss StrSlice) Map(f func(string) string) []string {
    vsm := make([]string, len(ss))
    for i, v := range ss {
        vsm[i] = f(v)
    }
    return vsm
}