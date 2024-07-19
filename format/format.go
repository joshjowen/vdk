package format

import (
	"github.com/joshjowen/vdk/av/avutil"
	"github.com/joshjowen/vdk/format/aac"
	"github.com/joshjowen/vdk/format/flv"
	"github.com/joshjowen/vdk/format/mp4"
	"github.com/joshjowen/vdk/format/rtmp"
	"github.com/joshjowen/vdk/format/rtsp"
	"github.com/joshjowen/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
