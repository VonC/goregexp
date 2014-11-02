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
	})
}
