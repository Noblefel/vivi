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

var Out io.Writer = os.Stdout

func init() { fmt.Fprint(Out, "\033[?25l") }

// Choices list out the options and return the selected index after enter keypress.
func Choices(options ...string) int {
	printChoices(0, options)
	current := 0

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			os.Exit(0)
		}

		if key.Code == keys.Enter || key.Code == keys.Space {
			return true, nil
		}

		if len(options) == 0 {
			return false, nil
		}

		if key.Code == keys.Down {
			if current+1 >= len(options) {
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

		fmt.Fprintf(Out, "\033[%dA", len(options))
		fmt.Fprint(Out, "\033[0J")
		printChoices(current, options)
		return false, nil
	})

	return current
}

func printChoices(current int, options []string) {
	for i, o := range options {
		if current == i {
			fmt.Fprint(Out, "\033[1;38;5;75m>\033[0m ")
			fmt.Fprintf(Out, "\033[38;5;75m%s\033[0m\n\r", o)
		} else {
			fmt.Fprintf(Out, "  %s\n\r", o)
		}
	}
}

// Password hides input by replacing it with the given placeholder.
func Password(placeholder string) string {
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
			fmt.Fprint(Out, "\033[1D\033[K")
			return false, nil
		}

		if key.Code == keys.Space {
			buf = append(buf, ' ')
			fmt.Fprint(Out, placeholder)
		}

		if key.Code == keys.RuneKey {
			buf = append(buf, key.Runes...)
			fmt.Fprint(Out, placeholder)
		}

		return false, nil
	})

	fmt.Fprintf(Out, "\n")
	return string(buf)
}
