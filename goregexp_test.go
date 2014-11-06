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
	})
}
