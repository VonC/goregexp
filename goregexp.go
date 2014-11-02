package goregexp

import (
	"fmt"
	"regexp"
)

/* Encapsulate a regex and a string,
for managing results from FindAllStringSubmatchIndex */
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

/* Build new result from FindAllStringSubmatchIndex on a string */
func NewReres(s string, r *regexp.Regexp) *Reres {
	matches := r.FindAllStringSubmatchIndex(s, -1)
	return &Reres{r, s, matches, 0, 0}
}
