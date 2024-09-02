# A.D.L.R.
### **A**utomating **D**ependency **L**icense **R**equirements

[![GoDoc](https://godoc.org/github.com/blocky/adlr?status.svg)](https://godoc.org/github.com/blocky/adlr)
[![Build Status](https://www.travis-ci.com/blocky/adlr.svg?token=JczzdP6eMqmEqysZ8pDf&branch=main)](https://www.travis-ci.com/blocky/adlr)
[![Go Report Card](https://goreportcard.com/badge/github.com/blocky/adlr)](https://goreportcard.com/report/github.com/blocky/adlr)

ADLR is a project that attempts to automate fulfillment of golang module dependency license requirements in files suitable for vcs.

# Disclaimer
**The ADLR project offers no legal advice or license compliance guarantee.
It is your responsibility to ensure compliance with licenses you interact with**

# Overview
## ADLR's License Lock
ADLR creates a license lock file.
This is a readable and manually edittable json file of your directly imported golang dependencies and their licenses.
It is much like a `go.mod`, and you can save this file in your version control system.
Some benefits of this:
+ monitor imports' licenses across versions
+ automate listing *copyrights*|*permissions*|*warranties* for licenses in your source code

## Get ADLR
`go get github.com/blocky/adlr/...`

## ADLR and Distributable Inclusion
Automate a license information command for your distributable with your license lock file

### Linker Flag
1. Serialize the lock file _(Go Linker flag requires strings to have no spaces or newlines)_
2. Pass to a variable in your code with the `-ldflags` build flag
3. Deserialize and unmarshal for license information

### Go 1.16 File Embedding
1. Embed your verified license file with an embed directive:
```golang
\\go:embed verified-licenses.json
var LicensesBytes []byte
```
2. Unmarshal for license information
```golang
var licenses []adlr.DependencyLock
err := json.Unmarshal(licenses, &LicensesBytes)
```

# ADLR Process
The ADLR process consists of 4 steps:
1. Generate a list of module dependencies required to build your go project
2. Locate the license files for each of the dependencies
3. Identify the license types for each of the dependencies
4. Verify the license types against a whitelist of allowed licenses

This functionality is realized in the `adlr/cmd` source code which is compiled as the `adlr-cli` tool.

## Your Golang Module buildlist
Use the `adlr-cli` tool in your golang module:
```bash
# cd to go project directory (that contains mod file)
# Default output file is `./buildlist.json` 
$ adlr-cli license buildlist

# Or output to specific file
$ adlr-cli license buildlist -b /path/to/my-buildlist.json
```
This will generate a json list of all golang modules/projects required to build your module.

## Text Mining Licenses
Unfortunately, golang does not yet have a standard for module license files.
There names can be lowercase, uppercase, with or without a file extension, or not even named "license", such as "COPYLEFT".
To solve this, ADLR uses [go-license-detector](https://github.com/go-enry/go-license-detector) to text mine projects for potential license files.
For licenses that cannot be located, they will be appended to a list and output to stderr.
```bash
# Assumes existence of `./buildlist.json` file
# Default output file is `./located-licenses.json`
$ adlr-cli license locate

# Or provide custom input and output files
$ adlr-cli license locate \
-b /path/to/my-buildlist.json \
-l /path/to/my-located-licenses.json
```

## Automatically Determining License
From prospecting, one or multiple matches are returned for a golang module with license type, file name, and confidence.
With preset confidence values, ADLR attempts to automatically determine the license for each golang module.
For licenses that cannot be determined, they will be appended to a list and output to stderr.
```bash
# Assumes existence of `./located-licenses.json` file
# Default output file is `./identified-licenses.json`
$ adlr-cli license identify

# Or provide custom input and output files
$ adlr-cli license identify \
-l /path/to/my-located-licenses.json \
-i /path/to/my-identified-licenses.json
```

## Auditing Locked License types
Finally, identified licenses are compared to a whitelist of approved license types and those that pass are written to a final file.
Licenses that are not on the whitelist are appended to a list and output to stderr.
```bash
# Assumes existence of `./identified-licenses.json` file
# Default output file is `./verified-licenses.json`
$ adlr-cli license verify

# Or provide custom input and output files
$ adlr-cli license verify \
-i /path/to/my-located-licenses.json \
-v /path/to/my-verified-licenses.json
```

# Development
Contributions are welcome! Contact BLOCKY through our website [www.blocky.rocks](www.blocky.rocks).

## Branch Practices
### Branches
+ **feature/**: Used for adding features, increments semver x.**y**.z
+ **bugfix/**: Used for fixing bugs, increments semver x.y.**z**
+ **chore/**: Used for small chores, tasks, etc and does not usually result in a semver increase/release

### Main & Develop
Due to recent errors in PR merges to the main branch, *all PR's must initially merge into the* **develop branch**, **checked for bugs**, *then a PR merging* **develop's** *changes into* **main**

### Squash Merging
We use squash merging for PR's. Therefore, not all of your commits are required to pass testing **besides the last commit**
