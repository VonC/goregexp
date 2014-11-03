package goregexp

import (
	"fmt"
	"regexp"
)

// Reres encapsulate a regex, a string, and the results
// from applying the regexp on the string
// internally managing results from FindAllStringSubmatchIndex
type Reres struct {
	r        *regexp.Regexp
	s        string
	matches  [][]int
	i        int
	previous int
}

func (r *Reres) String() string {
	msg := fmt.Sprintf("Regexp res for '%v': (%v-%v; len %v) %v", r.r, r.i, r.previous, len(r.s), r.matches)
	return msg
}

// NewReres builds a new regexp result
// (internally using FindAllStringSubmatchIndex on a string)
func NewReres(s string, r *regexp.Regexp) *Reres {
	matches := r.FindAllStringSubmatchIndex(s, -1)
	return &Reres{r, s, matches, 0, 0}
}
