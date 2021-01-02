package starwars

// Terminal ex) iTerm2
type Terminal interface {
	Setup()
	Cmd() string
	Args() []string
}
