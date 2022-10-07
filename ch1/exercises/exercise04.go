// dup1 shows any line of its input
// that occurs more than once, preceded by the
// count of occurrences. It reads from a
// sequence of files, specified by its
// command line arguments, or, if none,
// reads from the standard input.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filenames := os.Args[1:]
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	if len(filenames) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, filename := range filenames {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(f, counts, foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v\t%d\t%s\n", foundIn[line], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		foundIn[line] = append(foundIn[line], f.Name())
	}
	// WARNING: ignoring potential errors of input.Err()
}
