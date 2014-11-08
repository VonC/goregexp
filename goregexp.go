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

/* Text returns the full initial text on which the regex was applied */
func (rr *Reres) Text() string {
	return rr.s
}

/* HasAnyMatch checks if there is any match */
func (rr *Reres) HasAnyMatch() bool {
	return len(rr.matches) > 0
}

/* HasNext checks if there is one more match */
func (rr *Reres) HasNext() bool {
	return rr.i < len(rr.matches)
}

/* Next() refers to the next match, for Group() to works with */
func (rr *Reres) Next() {
	rr.previous = rr.matches[rr.i][1]
	rr.i = rr.i + 1
}

/*  ResetNext() get back to the first match */
func (rr *Reres) ResetNext() {
	rr.i = 0
	rr.previous = 0
}

/* Prefix gets the string from the last match to current one */
func (rr *Reres) Prefix() string {
	mi := rr.matches[rr.i]
	return rr.s[rr.previous:mi[0]]
}

/* Suffix gets the string from current match to the end ofthe all string */
func (rr *Reres) Suffix() string {
	mi := rr.matches[rr.i]
	res := ""
	if len(rr.s) > mi[1] {
		res = rr.s[mi[1]:]
	}
	return res
}

/* FirstChar returns the first character of the current match */
func (rr *Reres) FirstChar() uint8 {
	mi := rr.matches[rr.i]
	return rr.s[mi[0]]
}

/* Test if first character of the current match is an escape */
func (rr *Reres) IsEscaped() bool {
	mi := rr.matches[rr.i]
	return rr.s[mi[0]] == '\\'
}

/* Full string matched for the current group */
func (rr *Reres) FullMatch() string {
	mi := rr.matches[rr.i]
	return rr.s[mi[0]:mi[1]]
}

/* Check if the ith group if present in the current match */
func (rr *Reres) HasGroup(j int) bool {
	res := false
	mi := rr.matches[rr.i]
	if len(mi) > (j*2)+1 {
		if mi[j*2] > -1 {
			if mi[j*2] < mi[(j*2)+1] {
				res = true
			}
		}
	}
	return res
}

/* return the ith group string, if present in the current match */
func (rr *Reres) Group(i int) string {
	res := ""
	if rr.HasGroup(i) {
		mi := rr.matches[rr.i]
		res = rr.s[mi[i*2]:mi[(i*2)+1]]
	}
	return res
}
