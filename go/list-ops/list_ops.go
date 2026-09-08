package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accumulator := initial

	for i := range s {
		accumulator = fn(accumulator, s[i])
	}

	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial

	for i := range s {
		accumulator = fn(s[s.Length()-1-i], accumulator)
	}

	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := make(IntList, 0)
	offset := 0

	for i := range s {
		if fn(s[i]) {
			out = out.Append(IntList{s[i]})
			offset++
		}
	}

	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := make(IntList, s.Length())

	for i := range s {
		out[i] = fn(s[i])
	}

	return out
}

func (s IntList) Reverse() IntList {
	out := make(IntList, s.Length())

	for i := range s {
		out[i] = s[s.Length()-1-i]
	}

	return out
}

func (s IntList) Append(lst IntList) IntList {
	out := make(IntList, s.Length()+lst.Length())

	for i := range s {
		out[i] = s[i]
	}

	for i := range lst {
		out[i+s.Length()] = lst[i]
	}

	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	total := s.Length()

	for i := range lists {
		total += lists[i].Length()
	}

	out := make(IntList, total)
	offset := 0

	for i := range lists {
		for j := range lists[i] {
			out[offset] = lists[i][j]
			offset++
		}
	}

	return out
}
