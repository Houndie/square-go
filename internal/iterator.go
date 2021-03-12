package internal

type Iterator struct {
	newSegment    func(string) (int, string, error)
	err           error
	idx           int
	segmentLength int
	cursor        string
	done          bool
}

func NewIterator(newSegment func(string) (int, string, error)) *Iterator {
	return &Iterator{
		newSegment:    newSegment,
		idx:           -1,
		segmentLength: 0,
		cursor:        "",
		done:          false,
	}
}

func (i *Iterator) Error() error {
	return i.err
}

func (i *Iterator) Value() int {
	return i.idx
}

func (i *Iterator) Next() bool {
	i.idx = i.idx + 1
	if i.idx < i.segmentLength {
		return true
	}

	if i.done {
		return false
	}

	var err error
	i.segmentLength, i.cursor, err = i.newSegment(i.cursor)
	if err != nil {
		i.err = err
		return false
	}

	if i.segmentLength == 0 {
		return false
	}
	i.idx = 0
	if i.cursor == "" {
		i.done = true
	}
	return true
}
