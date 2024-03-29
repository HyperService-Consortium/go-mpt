module github.com/HyperService-Consortium/go-mpt

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190422183909-d864b10871cd
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190422233926-fe54fb35175b
)

require (
	github.com/HyperService-Consortium/go-hexutil v1.0.1-0.20190901115112-895ca17550ce
	github.com/HyperService-Consortium/go-rlp v1.0.1-0.20190903144357-b5693c05a6b8
	github.com/syndtr/goleveldb v1.0.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)
