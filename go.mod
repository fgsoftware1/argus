module babywolf.io/argus

go 1.22.7

require babywolf.io/utils v0.0.0-00010101000000-000000000000

require (
	github.com/jedib0t/go-pretty/v6 v6.6.0 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
)

require (
	github.com/fatih/color v1.17.0
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/exp v0.0.0-20241004190924-225e2abe05e6
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/term v0.25.0
)

replace babywolf.io/utils => ./utils
