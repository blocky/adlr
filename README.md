A.D.L.R. **A**utomating **D**ependency **L**icense **R**equirements

[![Build Status](https://www.travis-ci.com/blocky/adlr.svg?token=JczzdP6eMqmEqysZ8pDf&branch=main)](https://www.travis-ci.com/blocky/adlr)

ADLR is a project that attempts to automate the fulfillment of your golang module's dependency requirements.

Using the command in your golang module:
```
$ go list -m -json all
```
you can generate a json list of all the golang modules required to build it.
ADLR takes this list and attempts to create a file with all your dependency licenses.


This file is called `license.lock` and is an edittable json list of each directly imported dependency in the format:
```json
[
 {
  "name": "github.com/blocky/prettyprinter",
  "version": "v1.0.0",
  "license": {
   "kind": "MIT",
   "text": "MIT License\n\nCopyright (c) 2020-2021 Ian Hecker, David Millman and contributors..."
  }
 },
]
```
You can save this file in your version control system, and monitor not only when your project imports another dependency,
but what kind of license it brings with it.



# Dependencies for testing
Mockery - mockery v1 is used to autogenerate code for golang interfaces. Mocked interfaces are outputted to the internal/mocks/ folder. The golang binary tool can be downloaded from https://github.com/vektra/mockery
