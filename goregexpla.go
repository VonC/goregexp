package goregexp

import (
	"bytes"
	"regexp"
)

// Qualifier qualifies what the regexp is looking ahead for
type Qualifier func(lh string, match []int, s string) bool

// NewReresLAGroup builds new result from FindAllStringSubmatchIndex
// on a string, validated by last group being a lookahead after each match
func NewReresLAGroup(s string, r *regexp.Regexp) *Reres {
	return newReresLA(s, r, nil)
}

// NewReresLAQual builds new result from FindAllStringSubmatchIndex
// on a string, validated by last group being a lookahead after each match,
// if that last group match qualifies
func NewReresLAQual(s string, r *regexp.Regexp, q Qualifier) *Reres {
	return newReresLA(s, r, q)
}

// newReresLA builds new result from FindAllStringSubmatchIndex on a string,
// validated by last group being a lookahead after each match
func newReresLA(s string, r *regexp.Regexp, q Qualifier) *Reres {
	bf := bytes.NewBufferString(s)
	by := bf.Bytes()
	m := [][]int{}
	lg := []int{}
	res := &Reres{r, s, m, 0, 0}
	shift := 0
	for match := r.FindSubmatchIndex(by); match != nil && len(match) > 0; match = r.FindSubmatchIndex(by) {
		if len(match) > 0 {
			//fmt.Printf("\nMatch '%v' '%v'\n-------\n", match, string(by))
			match, lg = match[:len(match)-2], match[len(match)-2:]
			for i, mi := range match {
				match[i] = mi + shift
				/*
					ss := ""
					if mi > -1 {
						ss = s[mi+shift:]
					}
					fmt.Printf("\ni=%v: shift=%v match[i]=%v: s[match[%v]]='%v'\n", i, shift, mi, mi+shift, ss)
				*/
			}
			//fmt.Printf("\nAppend '%v'\n===\n", match)
			if lg[0] <= lg[1] && lg[0] > -1 {
				lh := string(by[lg[0]:lg[1]])
				by = by[lg[0]:]
				shift = shift + lg[0]
				delta := lg[1] - lg[0]
				match[1] = match[1] - delta
				//fmt.Printf("\nBYQ (%v-%v)'%v' '%v' => lh='%v'\n'%v'\n-------\n", len(by), match[1], by, string(by), lh, match)
				if q == nil || q(lh, match, s) {
					m = append(m, match)
				}
			} else {
				m = append(m, match)
				//fmt.Printf("\nBY (%v-%v)'%v' '%v'\n'%v'\n-------\n", len(by), match[1]-shift, by, string(by), match)
				by = by[match[1]-shift:]
				shift = shift + (match[1] - shift)
			}
		}
	}
	res.matches = m
	//fmt.Printf("\n*** %v======\n", res.matches)
	return res
}
