package device

import (
	"github.com/vladimirvivien/go4vl/v4l2"
)

type config struct {
	useMPlane       bool
	ioType          v4l2.IOType
	pixFormat       v4l2.PixFormat
	pixFormatMPlane v4l2.PixFormatMPlane
	bufSize         uint32
	fps             uint32
	bufType         uint32
	frames          bool
}

type Option func(*config)

func WithUseMPlane(useMPlane bool) Option {
	return func(o *config) {
		o.useMPlane = useMPlane
	}
}

func WithIOType(ioType v4l2.IOType) Option {
	return func(o *config) {
		o.ioType = ioType
	}
}

func WithPixFormat(pixFmt v4l2.PixFormat) Option {
	return func(o *config) {
		o.pixFormat = pixFmt
	}
}

func WithPixFormatMPlane(pixFmtMp v4l2.PixFormatMPlane) Option {
	return func(o *config) {
		o.pixFormatMPlane = pixFmtMp
	}
}

func WithBufferSize(size uint32) Option {
	return func(o *config) {
		o.bufSize = size
	}
}

// WithFrames makes a capture device deliver Frame values, with the buffer's
// sequence number and timestamp, on GetFrames() instead of raw bytes on
// GetOutput().
func WithFrames() Option {
	return func(o *config) {
		o.frames = true
	}
}

func WithFPS(fps uint32) Option {
	return func(o *config) {
		o.fps = fps
	}
}

func WithVideoCaptureEnabled() Option {
	return func(o *config) {
		o.bufType = v4l2.BufTypeVideoCapture
	}
}

func WithVideoOutputEnabled() Option {
	return func(o *config) {
		o.bufType = v4l2.BufTypeVideoOutput
	}
}
