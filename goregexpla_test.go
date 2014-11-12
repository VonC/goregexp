package goregexp

import (
	"regexp"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReresLA(t *testing.T) {
	Convey("Test new regexp res can test LookAhead", t, func() {

		Convey("Regexps can simulate a lookahead at the end of a regexp", t, func() {
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
		Convey("Regexps can simulate a lookahead at the end of a regexp choice", t, func() {
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

	})
}
