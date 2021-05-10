# A.D.L.R.
### **A**utomating **D**ependency **L**icense **R**equirements

[![Build Status](https://www.travis-ci.com/blocky/adlr.svg?token=JczzdP6eMqmEqysZ8pDf&branch=main)](https://www.travis-ci.com/blocky/adlr)

ADLR is a project that attempts to automate fulfillment of golang module dependency requirements.

# Disclaimer
**The ADLR project offers no legal advice or license compliance guarantee. It is your responsibility to ensure compliance with licenses you interact with**

# Overview
## License Lock
ADLR creates a license lock file. This is a readable and manually edittable json file of your directly imported golang dependencies and their licenses. It is much like a `go.mod`, and you can save this file in your version control system to help monitor imports and their licenses.

## Distributable Inclusion
Serialize the contents of the license lock file and include it in your distributable for a license information command. Serialize the lock file, save it in a golang variable with the golang `-ldflags` flag, and deserialize for printout. An example of this is built in to the repo. See `Makefile`, `sh/build.sh`, and the `cmd/` folder for details. Or test out ADLR's `about license(s)` commands with `make build`.

# ADLR Process
## Your Golang Module buildlist
Using the command in your golang module:
```sh
$ go list -m -json all > buildlist.json
```
you can generate a json list of all golang modules/projects required to build your module.
If your project is complex this list can be long. Luckily, ADLR filters out only your directly imported modules.
```golang
parser := gotool.MakeBuildListParser()
mods, err := parser.ParseModuleList(buildlist)
...

direct := gotool.FilterDirectImportModules(mods)
```

## Text Mining Licenses
Unfortunately, golang does not yet have a standard for module license files. There names can be lowercase, uppercase, with or without a file extension, or not even named "license", such as "COPYLEFT". To solve this, ADLR uses text mining to prospect potential license file matches and their confidences with https://github.com/go-enry/go-license-detector.
```golang
direct := gotool.FilterDirectImportModules(mods)

prospects := adlr.MakeProspects(direct...)
mines, err := prospector.Prospect(prospects...)
...
```

## Automatically Determining License
From prospecting, one or multiple matches are returned for a golang module with license type, file name, and confidence. With preset confidence values, ADLR attempts to automatically determine the license for each golang module. If a license cannot be determined through mining, the license lock manager may be able to automatically determine it (only if a license lock file has already been created).
```golang
mines, err := prospector.Prospect(prospects...)
...

miner := adlr.MakeLicenseMiner()
locks, err := miner.Mine(mines...)
if err != nil && Verbose {
	fmt.Println(err.Error())
}
```

## Locking Dependencies and their Licenses
After mining, licenses are hopefully automatically determined. These are now ready to be locked into a file. For no pre-existing license lock, a new file is created. For an existing license lock, the new and old list of dependencies are merged.

New dependencies take priority, and will fill the lock file. But for new locks that are missing license fields, merging is attempted with pre-existing locks. For new locks that cannot be automatically resolved, the license lock manager will print them in stderr, asking for manual editting of the license lock file. These license edits will persist for that dependency.
```golang
locks, err := miner.Mine(mines...)
...

licenselock := adlr.MakeLicenseLock("./")
err = licenselock.Lock(locks)
...
```

## Auditing Locked License types
After locking, dependencies and their licenses have been written to the lock file. But unwanted license types may have slipped through. The auditing step will search through the lock file, checking license types against a whitelist. For any types not listed, an error is returned listing bad license types, and requesting whitelist inclusion or dependency removal.
```golang
licenselock := adlr.MakeLicenseLock("./")
err = licenselock.Lock(locks)
...

locks, err = licenselock.Read()
...
auditor := adlr.MakeLicenseAuditor()
err = auditor.Audit(locks)
...
```

# Dependencies for testing
Mockery - mockery v1 is used to autogenerate code for golang interfaces. Mocked interfaces are outputted to the internal/mocks/ folder. The golang binary tool can be downloaded from https://github.com/vektra/mockery
