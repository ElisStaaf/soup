package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type SearchDebug struct {
	Workers int
}

const (
	LITERAL = iota
	REGEX
)

type SearchOptions struct {
	Kind   int
	Lines  bool
	Regex  *regexp.Regexp
	Finder *stringFinder
}

type searchJob struct {
	path string
	opts *SearchOptions
}

func Search(paths []string, opts *SearchOptions, debug *SearchDebug) {
	searchJobs := make(chan *searchJob)

	var wg sync.WaitGroup
	for w := 0; w < debug.Workers; w++ {
		go searchWorker(searchJobs, &wg)
	}
	for _, path := range paths {
		dirTraversal(path, opts, searchJobs, &wg)
	}
	wg.Wait()
}

func dirTraversal(path string, opts *SearchOptions, searchJobs chan *searchJob, wg *sync.WaitGroup) {
	info, err := os.Lstat(path)
	if err != nil {
		log.Fatalf("couldn't lstat path %s: %s\n", path, err)
	}

	if !info.IsDir() {
		wg.Add(1)
		searchJobs <- &searchJob{
			path,
			opts,
		}
		return
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("couldn't open path %s: %s\n", path, err)
	}
	dirNames, err := f.Readdirnames(-1)
	if err != nil {
		log.Fatalf("couldn't read dir names for path %s: %s\n", path, err)
	}

	for _, deeperPath := range dirNames {
		dirTraversal(filepath.Join(path, deeperPath), opts, searchJobs, wg)
	}
}

func searchWorker(jobs chan *searchJob, wg *sync.WaitGroup) {
	for job := range jobs {
		f, err := os.Open(job.path)
		if err != nil {
			log.Fatalf("couldn't open path %s: %s\n", job.path, err)
		}

		scanner := bufio.NewScanner(f)
		isBinary := false

		line := 1
		for scanner.Scan() {
			text := scanner.Bytes()

			// Check the first buffer for NUL
			if line == 1 {
				isBinary = bytes.IndexByte(text, 0) != -1
			}

			if job.opts.Kind == LITERAL {
				if job.opts.Finder.next(text) != -1 {
					if isBinary {
						fmt.Printf("Binary file %s matches\n", job.path)
						break
					} else if job.opts.Lines {
						fmt.Printf("%s:%d %s\n", job.path, line, text)
					} else {
						fmt.Printf("%s %s\n", job.path, text)
					}
				}
			} else if job.opts.Kind == REGEX {
				if job.opts.Regex.Find(scanner.Bytes()) != nil {
					if isBinary {
						fmt.Printf("Binary file %s matches\n", job.path)
						break
					} else if job.opts.Lines {
						fmt.Printf("%s:%d %s\n", job.path, line, text)
					} else {
						fmt.Printf("%s %s\n", job.path, text)
					}
				}
			}
			line++
		}
		wg.Done()
	}
}

type stringFinder struct {
	pattern []byte
	badCharSkip [256]int
	goodSuffixSkip []int
}

func MakeStringFinder(pattern []byte) *stringFinder {
	f := &stringFinder{
		pattern:        pattern,
		goodSuffixSkip: make([]int, len(pattern)),
	}
	last := len(pattern) - 1

	for i := range f.badCharSkip {
		f.badCharSkip[i] = len(pattern)
	}
	for i := 0; i < last; i++ {
		f.badCharSkip[pattern[i]] = last - i
	}

	lastPrefix := last
	for i := last; i >= 0; i-- {
		if bytes.HasPrefix(pattern, pattern[i+1:]) {
			lastPrefix = i + 1
		}
		f.goodSuffixSkip[i] = lastPrefix + last - i
	}
	for i := 0; i < last; i++ {
		lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
			f.goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i
		}
	}

	return f
}

func longestCommonSuffix(a, b []byte) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}

func (f *stringFinder) next(text []byte) int {
	i := len(f.pattern) - 1
	for i < len(text) {
		j := len(f.pattern) - 1
		for j >= 0 && text[i] == f.pattern[j] {
			i--
			j--
		}
		if j < 0 {
			return i + 1
		}
		i += max(f.badCharSkip[text[i]], f.goodSuffixSkip[j])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
