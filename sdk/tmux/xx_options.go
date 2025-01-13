package tmux

func applyOptions[Opts any, Opt ~func(*Opts)](opts *Opts, opt []Opt) {
	for _, o := range opt {
		o(opts)
	}
}
