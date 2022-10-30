module example.com/hello

go 1.18

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000

require (
	github.com/yuin/goldmark v1.4.13 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/tools v0.2.0 // indirect
)
