package stringset

const testVersion = 4

type Set struct {
	elements map[string]bool
}

func New() Set {
	return Set{elements: make(map[string]bool)}
}

func NewFromSlice(ss []string) Set {
	elements := make(map[string]bool)

	if ss == nil {
		return Set{elements: elements}
	}

	for _, s := range ss {
		elements[s] = true
	}

	return Set{elements: elements}
}

func (s Set) String() string {
	str := "{"

	l := len(s.elements)
	var i int

	for k := range s.elements {
		str += "\"" + k + "\""
		if i < l-1 {
			str += ", "
		}
		i++
	}

	str += "}"

	return str
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(k string) bool {
	_, ok := s.elements[k]
	return ok
}

func (s Set) Add(str string) {
	s.elements[str] = true
}

func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}

	if !s1.IsEmpty() && s2.IsEmpty() {
		return false
	}

	for k := range s1.elements {
		if _, ok := s2.elements[k]; !ok {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}

	if !s1.IsEmpty() && s2.IsEmpty() {
		return true
	}

	for k := range s1.elements {
		if _, ok := s2.elements[k]; ok {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.elements) != len(s2.elements) {
		return false
	}

	for k := range s1.elements {
		if _, ok := s2.elements[k]; !ok {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	commonSet := Set{elements: make(map[string]bool)}

	for k := range s1.elements {
		if _, ok := s2.elements[k]; ok {
			commonSet.elements[k] = true
		}
	}

	return commonSet
}

func Difference(s1, s2 Set) Set {
	diff := Set{elements: make(map[string]bool)}

	if s1.IsEmpty() && s2.IsEmpty() {
		return diff
	}

	if s1.IsEmpty() && !s2.IsEmpty() {
		return s1
	}

	if !s1.IsEmpty() && s2.IsEmpty() {
		return s1
	}

	for k := range s1.elements {
		if _, ok := s2.elements[k]; !ok {
			diff.elements[k] = true
		}
	}

	return diff
}

func Union(s1, s2 Set) Set {
	for k := range s1.elements {
		if _, ok := s2.elements[k]; !ok {
			s2.elements[k] = true
		}
	}
	return s2
}
