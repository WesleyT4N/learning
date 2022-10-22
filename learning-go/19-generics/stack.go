package generics


type Stack[T any] struct {
    values []T
}


func (s *Stack[T]) Push(value T) {
    s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var empty_val T
        return empty_val, false
    }

    ind := len(s.values) - 1
    el := s.values[ind]
    s.values = s.values[:ind]
    return el, true
}

