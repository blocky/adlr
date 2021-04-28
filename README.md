A.D.L.R. **A**utomating **D**ependency **L**icense **R**equirements

[![Build Status](https://www.travis-ci.com/blocky/adlr.svg?token=JczzdP6eMqmEqysZ8pDf&branch=main)](https://www.travis-ci.com/blocky/adlr)

# Dependencies for testing
Mockery - mockery v1 is used to autogenerate code for golang interfaces. Mocked interfaces are automatically outputted to the mocks/ folder. The golang binary tool can be downloaded from https://github.com/vektra/mockery

Golang 1.16.3 - already have a different version of go installed? You can install multiple versions, https://golang.org/doc/manage-install, through the commands:
```
$ go get golang.org/dl/go1.16.3
$ go1.16.3 download
```