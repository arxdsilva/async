package async

type Option func(*config)

type config struct {
	chanSize int
}

// WithChanSize changes the default channel size to the assigned number
func WithChanSize(size int) Option {
	return func(c *config) {
		c.chanSize = size
	}
}
