package leetcode

var palceHolder = struct{}{}

type isomorphic struct {
	mapping       map[byte]byte
	existedTarget map[byte]struct{}
}

func (o isomorphic) safeAndAdd(original byte, target byte) bool {
	pre, ok := o.mapping[original]
	if ok {
		if pre != target {
			return false
		}
		return true
	}
	_, targetExisted := o.existedTarget[target]
	if targetExisted {
		return false
	}
	o.mapping[original] = target
	o.existedTarget[target] = palceHolder
	return true
}

func newIsomorphic() *isomorphic {
	return &isomorphic{
		mapping:       make(map[byte]byte),
		existedTarget: make(map[byte]struct{}),
	}
}

func isIsomorphic(s string, t string) bool {
	sL, tL := len(s), len(t)
	if sL != tL {
		return false
	}
	mark := newIsomorphic()
	for i := 0; i < sL; i++ {
		if !mark.safeAndAdd(s[i], t[i]) {
			return false
		}
	}
	return true
}
