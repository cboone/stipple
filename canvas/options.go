package canvas

// Option is a functional option for configuring a Canvas.
type Option func(*Canvas)

// WithColor returns an option that enables per-cell ANSI color support.
func WithColor() Option {
	return func(canvas *Canvas) {
		canvas.colorEnabled = true
	}
}

// WithInvertedY returns an option that inverts the Y-axis direction.
// By default, Y increases downward (standard screen coordinates).
// With this option, Y increases upward (mathematical coordinates).
func WithInvertedY() Option {
	return func(canvas *Canvas) {
		canvas.invertY = true
	}
}
