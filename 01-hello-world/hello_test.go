package main

import "testing"


func TestHello(t *testing.T) {
    assertEq := func(t testing.TB, got string, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to peopl", func (t *testing.T) {
        assertEq(t, Hello("Wes"), "Hello, Wes")
    })
    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        assertEq(t, Hello(""), "Hello, World")
    })
}
