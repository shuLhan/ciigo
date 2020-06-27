module github.com/shuLhan/ciigo

go 1.13

require (
	github.com/bytesparadise/libasciidoc v0.2.1-0.20190929071746-6c57b552cca9
	github.com/onsi/ginkgo v1.12.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shuLhan/share v0.16.1-0.20200627020222-d13579896847
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/yuin/goldmark v1.1.32
	github.com/yuin/goldmark-meta v0.0.0-20191126180153-f0638e958b60
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/tools v0.0.0-20200604183345-4d5ea46c79fe // indirect
)

exclude github.com/bytesparadise/libasciidoc v0.3.0

exclude github.com/bytesparadise/libasciidoc v0.4.0

replace github.com/bytesparadise/libasciidoc => github.com/bytesparadise/libasciidoc v0.2.1-0.20190929071746-6c57b552cca9
