package generics

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
    t.Helper()
    if got != want {
        t.Errorf("got %v, waht %v", got, want)
    }
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
    t.Helper()
    if got == want {
        t.Errorf("didn't want %v", got)
    }
}

func AssertTrue(t *testing.T, got bool) {
    t.Helper()
    if !got {
        t.Errorf("got %v, want true", got)
    }
}

func AssertFalse(t *testing.T, got bool) {
    t.Helper()
    if got {
        t.Errorf("got %v, want false", got)
    }
}

func TestAssertFunctions(t *testing.T) {
    t.Run("asserting on integers", func(t *testing.T) {
        AssertEqual(t, 1, 1)
        AssertNotEqual(t, 1, 2)
    });
    t.Run("asserting on strings", func(t *testing.T) {
        AssertEqual(t, "hello", "hello")
        AssertNotEqual(t, "hello", "Grace")
    });
}

func TestStack(t *testing.T) {
    t.Run("integer stack", func(t *testing.T) {
        stack := new(Stack[int])

        AssertTrue(t, stack.IsEmpty())

        stack.Push(123)
        AssertFalse(t, stack.IsEmpty())

        stack.Push(456)
        val, _  := stack.Pop()
        AssertEqual(t, val, 456)
        val, _  = stack.Pop()
        AssertEqual(t, val, 123)
        AssertTrue(t, stack.IsEmpty())
    })
    t.Run("string stack", func(t *testing.T) {
        stack := new(Stack[string])

        AssertTrue(t, stack.IsEmpty())

        stack.Push("123")
        AssertFalse(t, stack.IsEmpty())

        stack.Push("456")
        val, _  := stack.Pop()
        AssertEqual(t, val, "456")
        val, _  = stack.Pop()
        AssertEqual(t, val, "123")
        AssertTrue(t, stack.IsEmpty())
    })
}
