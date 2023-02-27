package async

type Option func(*config)

type config struct {
	chanSizeData int
	chanSizeErr  int
}

// WithChanSizeData changes the default channel size to the assigned number
func WithChanSizeData(size int) Option {
	return func(c *config) {
		c.chanSizeData = size
	}
}

// WithChanSizeErr changes the default channel size to the assigned number
func WithChanSizeErr(size int) Option {
	return func(c *config) {
		c.chanSizeErr = size
	}
}
