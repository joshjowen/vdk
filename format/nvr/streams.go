package nvr

import "github.com/joshjowen/vdk/av"

type Stream struct {
	codec av.CodecData
	idx   int
}
