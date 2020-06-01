package logger

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogLevelText(t *testing.T) {
	Convey("test expected labels", t, func() {
		So(GetLogLevelString(LogLevelDebug), ShouldEqual, "DEBUG")
		So(GetLogLevelString(LogLevelTrace), ShouldEqual, "TRACE")
		So(GetLogLevelString(LogLevelInfo), ShouldEqual, "INFO")
		So(GetLogLevelString(LogLevelError), ShouldEqual, "ERROR")
		So(GetLogLevelString(LogLevelFatal), ShouldEqual, "FATAL")

	})
}
