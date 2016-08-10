package collections

type Set map[interface{}]bool

func NewSet() Set {
	return make(map[interface{}]bool)
}

func (s *Set) Contains(e interface{}) bool {
	return (*s)[e]
}

func (s *Set) Add(elements ...interface{}) {
	for _, e := range elements {
		if !(*s).Contains(e) {
			(*s)[e] = true
		}
	}
}

func (s *Set) addAll(elements []interface{}) {
	s.Add(elements...)
}

func (s *Set) remove(e interface{}) {
	delete(*s, e)
}
