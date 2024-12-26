package vivi

import (
	"strings"
	"testing"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var replacer = strings.NewReplacer(
	"\033[1;38;5;75m", "",
	"\033[38;5;75m", "",
	"\033[?25l", "",
	"\033[?25h", "",
	"\033[0m", "",
	"\n", "",
)

func TestChoice(t *testing.T) {
	var sb strings.Builder
	Out = &sb

	t.Run("choosing 'a' by pressing ENTER", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		answer := Choices("a", "b", "c")

		if answer != 0 {
			t.Errorf("want index of 0, got %d", answer)
		}

		b := strings.Contains(replacer.Replace(sb.String()), "> a")

		if !b {
			t.Errorf("arrow should point at option 'a'\ngot: %q", sb.String())
		}
	})

	t.Run("choosing 'c' by pressing two DOWN and ENTER", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress(keys.Down)
			keyboard.SimulateKeyPress(keys.Down)
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		answer := Choices("a", "b", "c")

		if answer != 2 {
			t.Fatalf("want index of 2, got %d", answer)
		}

		b := strings.Contains(replacer.Replace(sb.String()), "> c")

		if !b {
			t.Errorf("arrow should point at option 'c'\ngot: %q", sb.String())
		}
	})

	t.Run("if no options", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress(keys.Down)
			keyboard.SimulateKeyPress(keys.Down)
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		Choices()

		s := replacer.Replace(sb.String())

		if s != "" {
			t.Errorf("should not output anything, got %q", s)
		}
	})
}

func TestPassword(t *testing.T) {
	var sb strings.Builder
	Out = &sb

	t.Run("with '*' placeholder and the text have a space", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress("abc")
			keyboard.SimulateKeyPress(keys.Space)
			keyboard.SimulateKeyPress("def")
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		value := Password("*")
		s := replacer.Replace(sb.String())

		if s != "*******" {
			t.Errorf("wanted string output: *******\ngot: %q", s)
		}

		if value != "abc def" {
			t.Errorf("want value to be 'abc def', got %q", value)
		}
	})

	t.Run("pressing BACKSPACE and ENTER with no text", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress(keys.Backspace)
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		value := Password("*")

		if value != "" {
			t.Errorf("want value to be empty, got %q", value)
		}
	})

	t.Run("with '$$' placeholder", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress("abc")
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		value := Password("$$")
		s := replacer.Replace(sb.String())

		if s != "$$$$$$" {
			t.Errorf("wanted string output: $$$$$$\ngot: %q", s)
		}

		if value != "abc" {
			t.Errorf("want value to be 'abc', got %q", value)
		}
	})

	t.Run("with empty placeholder", func(t *testing.T) {
		defer sb.Reset()

		go func() {
			time.Sleep(50 * time.Millisecond)
			keyboard.SimulateKeyPress("abc")
			keyboard.SimulateKeyPress(keys.Enter)
		}()

		value := Password("")
		s := replacer.Replace(sb.String())

		if s != "" {
			t.Errorf("wanted string output to be empty\ngot: %q", s)
		}

		if value != "abc" {
			t.Errorf("want value to be 'abc', got %q", value)
		}
	})
}
