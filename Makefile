.PHONY: devserve

mdprev: mdprev.go
	GOPATH=~/gopaths/mdprev/ go build mdprev.go

devserve:
	GOPATH=~/gopaths/mdprev/ go run mdprev.go .

