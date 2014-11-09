package goregexp

import (
	"regexp"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProject(t *testing.T) {
	Convey("Test new regexp res container", t, func() {

		Convey("A Regexp res container can be build", func() {
			r := regexp.MustCompile("test")
			rx := NewReres("test", r)
			So(rx.String(), ShouldEqual, "Regexp res for 'test': (0-0; len 4) [[0 4]]")
		})

		Convey("A Regexp res container can display the string on which the regexp is applied", func() {
			r := regexp.MustCompile("test")
			rx := NewReres("test2", r)
			So(rx.Text(), ShouldEqual, "test2")
		})

		Convey("A Regexp res container knows if it has any match", func() {
			r := regexp.MustCompile("test")
			rx := NewReres("test3", r)
			So(rx.HasAnyMatch(), ShouldBeTrue)
			rx = NewReres("aaa", r)
			So(rx.HasAnyMatch(), ShouldBeFalse)
		})

		Convey("A Regexp res can reference groups", func() {
			r := regexp.MustCompile("(test)+")
			rx := NewReres("testtest", r)
			So(rx.HasNext(), ShouldBeTrue)
			rx.Next()
			So(rx.HasNext(), ShouldBeFalse)
			rx.ResetNext()
			So(rx.HasNext(), ShouldBeTrue)
		})

		Convey("A Regexp res can get prefix and suffix", func() {
			r := regexp.MustCompile("(test)+")
			rx := NewReres("aaatesttestbbb", r)
			So(rx.Prefix(), ShouldEqual, "aaa")
			So(rx.Suffix(), ShouldEqual, "bbb")
		})

		Convey("A Regexp res can get the first char of the current match", func() {
			r := regexp.MustCompile("(.est)")
			rx := NewReres("aaaTestcccUestbbb", r)
			So(rx.FirstChar(), ShouldEqual, 'T')
			So(rx.HasNext(), ShouldBeTrue)
			rx.Next()
			So(rx.FirstChar(), ShouldEqual, 'U')
		})

		Convey("A Regexp res can detect if the first char of the current match is \\", func() {
			r := regexp.MustCompile("(.est)")
			rx := NewReres("aaaTestccc\\estbbb", r)
			So(rx.IsEscaped(), ShouldBeFalse)
			rx.Next()
			So(rx.IsEscaped(), ShouldBeTrue)
		})

		Convey("A Regexp res can get full match", func() {
			r := regexp.MustCompile("(.est)")
			rx := NewReres("aaa1Testccc2Uestbbb3", r)
			So(rx.FullMatch(), ShouldEqual, "Test")
			rx.Next()
			So(rx.FullMatch(), ShouldEqual, "Uest")
		})

	})
}
