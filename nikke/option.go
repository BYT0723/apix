package nikke

type (
	Option func(*params)
	params struct {
		gcount     int
		uncompress bool
		retry      int
	}
)

func WithGcount(gcount int) Option {
	return func(p *params) {
		p.gcount = gcount
	}
}

func WithUncompress() Option {
	return func(p *params) {
		p.uncompress = true
	}
}

func WithRetry(retry int) Option {
	return func(p *params) {
		p.retry = retry
	}
}
