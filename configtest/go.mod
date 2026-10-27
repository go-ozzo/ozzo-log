module github.com/go-ozzo/ozzo-log/configtest

go 1.21

require (
	github.com/go-ozzo/ozzo-config v0.0.0-20260825173609-0b985a50a894
	github.com/go-ozzo/ozzo-log v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/hnakamur/jsonpreprocess v0.0.0-20171017030034-a4e954386171 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/go-ozzo/ozzo-log => ../
