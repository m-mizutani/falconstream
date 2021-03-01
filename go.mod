module github.com/m-mizutani/falconstream

go 1.13

require (
	github.com/aws/aws-sdk-go v1.25.4
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/m-mizutani/gofalcon v0.0.0-20191003010721-fc6517c9acd1
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli v1.22.1
	golang.org/x/net v0.0.0-20190930134127-c5a3c61f89f3 // indirect
	golang.org/x/sys v0.0.0-20191002091554-b397fe3ad8ed // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace (
        github.com/m-mizutani/gofalcon v0.0.0-20191003010721-fc6517c9acd1 => ../gofalcon
)
