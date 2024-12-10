// Package vivi have some helper to get user input from terminal.
// Built on top of [atomicgo.dev/keyboard]
package vivi

import (
	"fmt"
	"io"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var w io.Writer = os.Stdout

// Choices list out the options and return the selected index after enter keypress.
//
//	// "what is 5 * 5  ?"
//	answer := vivi.Choices(
//		"[1] 25",
//		"[2] 10",
//		"[3] 50",
//	)
//
//	if answer == 0 {
//		fmt.Println("✔️  Correct")
//	} else {
//		fmt.Println("❌ Try Again")
//	}
func Choices(options ...string) int {
	fmt.Fprint(w, "\033[?25l")
	printChoices(0, options)
	current := 0

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			os.Exit(0)
		}

		if key.Code == keys.Enter || key.Code == keys.Space {
			return true, nil
		}

		if key.Code == keys.Down {
			if current+1 == len(options) {
				current = 0
			} else {
				current++
			}
		}

		if key.Code == keys.Up {
			if current == 0 {
				current = len(options) - 1
			} else {
				current--
			}
		}

		fmt.Fprintf(w, "\033[%dA", len(options))
		fmt.Fprint(w, "\033[0J")
		printChoices(current, options)
		return false, nil
	})

	fmt.Fprintf(w, "\033[?25h")
	return current
}

func printChoices(current int, options []string) {
	for i, o := range options {
		if current == i {
			fmt.Fprintf(w, "\033[1m> %s\033[0m\n\r", o)
		} else {
			fmt.Fprintf(w, "  %s\n\r", o)
		}
	}
}

// Password hides input by replacing it with the given placeholder.
//
//	secret := vivi.Password("*")
func Password(placeholder string) string {
	if placeholder == "" {
		panic("empty placeholder")
	}

	fmt.Fprint(w, "\033[?25l")
	var buf []rune

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Enter {
			return true, nil
		}

		if key.Code == keys.CtrlC {
			os.Exit(0)
		}

		if key.Code == keys.Backspace && len(buf) > 0 {
			buf = buf[:len(buf)-1]
			fmt.Fprint(w, "\033[1D\033[K")
			return false, nil
		}

		if key.Code == keys.Space {
			buf = append(buf, ' ')
			fmt.Fprint(w, placeholder)
		}

		if key.Code == keys.RuneKey {
			buf = append(buf, key.Runes...)
			fmt.Fprint(w, placeholder)
		}

		return false, nil
	})

	fmt.Fprintf(w, "\n\033[?25h")
	return string(buf)
}

// Hidden hides input completely
func Hidden() string {
	var buf []rune

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Enter {
			return true, nil
		}

		if key.Code == keys.CtrlC {
			os.Exit(0)
		}

		if key.Code == keys.Backspace && len(buf) > 0 {
			buf = buf[:len(buf)-1]
			return false, nil
		}

		if key.Code == keys.Space {
			buf = append(buf, ' ')
		}

		if key.Code == keys.RuneKey {
			buf = append(buf, key.Runes...)
		}

		return false, nil
	})

	fmt.Fprint(w, "\n")
	return string(buf)
}
