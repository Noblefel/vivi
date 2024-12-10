package vivi

import (
	"strings"
	"testing"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var replacer = strings.NewReplacer(
	"\x1b[?25l", "",
	"\x1b[?25h", "",
)

func TestChoice(t *testing.T) {
	var sb strings.Builder
	w = &sb

	go func() {
		time.Sleep(50 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	answer := Choices("a", "b", "c")

	if answer != 0 {
		t.Errorf("want index of 0 when choosing 'a', got %d", answer)
	}

	out := "\033[1m> a\033[0m\n\r  b\n\r  c\n\r"
	s := replacer.Replace(sb.String())
	s = s[len(s)-len(out):]

	if out != s {
		t.Errorf("\nwanted string output: %q\ngot: %q", out, s)
	}
	sb.Reset()

	go func() {
		time.Sleep(50 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	answer = Choices("a", "b", "c")

	if answer != 2 {
		t.Fatalf("want index of 2 when choosing 'c', got %d", answer)
	}

	out = "  a\n\r  b\n\r\033[1m> c\033[0m\n\r"
	s = replacer.Replace(sb.String())
	s = s[len(s)-len(out):]

	if out != s {
		t.Errorf("\nwanted string output: %q\ngot: %q", out, s)
	}
}

func TestPassword(t *testing.T) {
	var sb strings.Builder
	w = &sb

	go func() {
		time.Sleep(50 * time.Millisecond)
		keyboard.SimulateKeyPress("abc")
		keyboard.SimulateKeyPress(keys.Space)
		keyboard.SimulateKeyPress("def")
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	value := Password("*")
	s := replacer.Replace(sb.String())

	if s != "*******\n" {
		t.Errorf("\nwanted string output: *******\\n\ngot: %q", s)
	}
	sb.Reset()

	if value != "abc def" {
		t.Errorf("want value to be 'abc def', got %q", value)
	}

	go func() {
		time.Sleep(50 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.Backspace)
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	value = Password("*")

	if value != "" {
		t.Errorf("want value to be empty, got %q", value)
	}

	go func() {
		time.Sleep(50 * time.Millisecond)
		keyboard.SimulateKeyPress("A")
		keyboard.SimulateKeyPress(keys.Backspace)
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	value = Password("*")

	if value != "" {
		t.Errorf("\nwanted value to be empty\ngot: %q", value)
	}
}

func TestHidden(t *testing.T) {
	var sb strings.Builder
	w = &sb

	go func() {
		time.Sleep(50 * time.Millisecond)
		keyboard.SimulateKeyPress("secret")
		keyboard.SimulateKeyPress(keys.Space)
		keyboard.SimulateKeyPress("input")
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	value := Hidden()
	s := replacer.Replace(sb.String())

	if s != "\n" {
		t.Errorf("\nwant output to be empty\ngot: %q", s)
	}
	sb.Reset()

	if value != "secret input" {
		t.Errorf("want value to be 'secret input', got %q", value)
	}
}
