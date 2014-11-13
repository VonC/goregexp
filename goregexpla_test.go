package goregexp

import (
	"regexp"
	"strings"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReresLA(t *testing.T) {
	Convey("Test new regexp res can test LookAhead", t, func() {

		Convey("Regexps can simulate a lookahead at the end of a regexp", func() {
			rx, _ := regexp.Compile(`a(b*)c($|de)`)
			r := NewReresLAGroup("aabbbbcdefabbcabbbcdeabcdabbc", rx)

			So(r.HasAnyMatch(), ShouldBeTrue)
			So(len(r.matches), ShouldEqual, 3)

			So(r.Prefix(), ShouldEqual, "a")
			So(r.FullMatch(), ShouldEqual, "abbbbc")
			So(r.Group(1), ShouldEqual, "bbbb")
			So(r.Suffix(), ShouldEqual, "defabbcabbbcdeabcdabbc")

			r.Next()

			So(r.Prefix(), ShouldEqual, "defabbc")
			So(r.FullMatch(), ShouldEqual, "abbbc")
			So(r.Group(1), ShouldEqual, "bbb")
			So(r.Suffix(), ShouldEqual, "deabcdabbc")

			r.Next()

			So(r.Prefix(), ShouldEqual, "deabcd")
			So(r.FullMatch(), ShouldEqual, "abbc")
			So(r.Group(1), ShouldEqual, "bb")
			So(r.Suffix(), ShouldEqual, "")
		})
		Convey("Regexps can simulate a lookahead at the end of a regexp choice", func() {
			rx, _ := regexp.Compile(`a(b*)c|d(e*)([^f])`)
			r := NewReresLAGroup("aabbbbcdaefdeeeabbc", rx)

			So(r.HasAnyMatch(), ShouldBeTrue)
			So(len(r.matches), ShouldEqual, 4)
			//fmt.Println(r)

			So(r.Prefix(), ShouldEqual, "a")
			So(r.FullMatch(), ShouldEqual, "abbbbc")
			So(r.Group(1), ShouldEqual, "bbbb")
			So(r.Suffix(), ShouldEqual, "daefdeeeabbc")

			r.Next()

			So(r.Prefix(), ShouldEqual, "")
			So(r.FullMatch(), ShouldEqual, "d")
			So(r.Group(1), ShouldEqual, "")
			So(r.Suffix(), ShouldEqual, "aefdeeeabbc")

			r.Next()

			So(r.Prefix(), ShouldEqual, "aef")
			So(r.FullMatch(), ShouldEqual, "deee")
			So(r.Group(2), ShouldEqual, "eee")
			So(r.Suffix(), ShouldEqual, "abbc")

			r.Next()

			So(r.Prefix(), ShouldEqual, "")
			So(r.FullMatch(), ShouldEqual, "abbc")
			So(r.Group(1), ShouldEqual, "bb")
			So(r.Suffix(), ShouldEqual, "")
		})

		Convey("Regexps can simulate a lookahead at the end of a regexp referencing a previous group", func() {
			// KbdDelimiterRx = /(?:\+|,)(?=#{CC_BLANK}*[^\1])/ */
			KbdDelimiterRx, _ := regexp.Compile(`(?:\+|,)([ \t]*[^ \t])`)
			r := NewReresLAQual(`Ctrl + Alt+T`, KbdDelimiterRx, kbdla)

			So(r.HasAnyMatch(), ShouldBeTrue)
			So(len(r.matches), ShouldEqual, 2)

			So(r.FullMatch(), ShouldEqual, "+")
			r.Next()
			So(r.FullMatch(), ShouldEqual, "+")

			r = NewReresLAQual(`
   Ctrl,T`, KbdDelimiterRx, kbdla)
			So(r.HasAnyMatch(), ShouldBeTrue)
			So(len(r.matches), ShouldEqual, 1)
			So(r.FullMatch(), ShouldEqual, ",")

			r = NewReresLAQual(`Ctrl,  ,`, KbdDelimiterRx, kbdla)
			So(r.HasAnyMatch(), ShouldBeFalse)
			r = NewReresLAQual(`Ctrl +  +a`, KbdDelimiterRx, kbdla)
			So(len(r.matches), ShouldEqual, 1)
			So(r.FullMatch(), ShouldEqual, "+")
		})
	})
}

func kbdla(lh string, match []int, s string) bool {
	m := strings.TrimSpace(lh)
	g1 := string(s[match[0]:match[1]])
	//fmt.Printf("\nm='%v' vs. g1='%v'\n", m, g1)
	return m != g1
}
