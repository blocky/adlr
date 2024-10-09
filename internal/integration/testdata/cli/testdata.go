package testdata

const LocateErrOut = `[{
  "Name": "github.com/go-git/gcfg",
  "Dir": "../..//vendor/github.com/go-git/gcfg",
  "Version": "v1.5.1-0.20230307220236-3a3c6141e376",
  "ErrStr": "no license file was found"
 },
 {
  "Name": "github.com/poy/onpar",
  "Dir": "../..//vendor/github.com/poy/onpar",
  "Version": "v0.3.3",
  "ErrStr": "could not clone repo from ../..//vendor/github.com/poy/onpar: repository not found"
 }]`

const IdentifyErrOut = `[
 {
  "Name": "golang.org/x/mod",
  "Dir": "../..//vendor/golang.org/x/mod",
  "Version": "v0.17.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "dario.cat/mergo",
  "Dir": "../..//vendor/dario.cat/mergo",
  "Version": "v1.0.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "github.com/ProtonMail/go-crypto",
  "Dir": "../..//vendor/github.com/ProtonMail/go-crypto",
  "Version": "v0.0.0-20230828082145-3c4c8a2d2371",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "github.com/cloudflare/circl",
  "Dir": "../..//vendor/github.com/cloudflare/circl",
  "Version": "v1.3.7",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   },
   {
    "license": "BSD-2-Clause",
    "confidence": 0.8385417,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-No-Military-License",
    "confidence": 0.7960199,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-Clear",
    "confidence": 0.7910448,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-No-Nuclear-License-2014",
    "confidence": 0.78109455,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "github.com/cyphar/filepath-securejoin",
  "Dir": "../..//vendor/github.com/cyphar/filepath-securejoin",
  "Version": "v0.2.4",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "github.com/hhatto/gorst",
  "Dir": "../..//vendor/github.com/hhatto/gorst",
  "Version": "v0.0.0-20181029133204-ca9f730cac5b",
  "ErrStr": "does not meet minimum lead: 0.930348 - 0.886667 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9303483,
    "file": "LICENSE"
   },
   {
    "license": "MIT",
    "confidence": 0.88666666,
    "file": "LICENSE"
   },
   {
    "license": "BSD-2-Clause",
    "confidence": 0.8333333,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-No-Military-License",
    "confidence": 0.800995,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-Clear",
    "confidence": 0.7910448,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-No-Nuclear-License-2014",
    "confidence": 0.7761194,
    "file": "LICENSE"
   },
   {
    "license": "MIT-0",
    "confidence": 0.55813956,
    "file": "LICENSE"
   },
   {
    "license": "X11-distribute-modifications-variant",
    "confidence": 0.45454544,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "github.com/pmezard/go-difflib",
  "Dir": "../..//vendor/github.com/pmezard/go-difflib",
  "Version": "v1.0.0",
  "ErrStr": "does not meet minimum lead: 0.956522 - 0.924623 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.95652175,
    "file": "LICENSE"
   },
   {
    "license": "BSD-2-Clause",
    "confidence": 0.92462313,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-No-Military-License",
    "confidence": 0.8309179,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-Clear",
    "confidence": 0.82608694,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause-No-Nuclear-License-2014",
    "confidence": 0.8115942,
    "file": "LICENSE"
   },
   {
    "license": "BSD-1-Clause",
    "confidence": 0.78571427,
    "file": "LICENSE"
   },
   {
    "license": "BSD-2-Clause-Views",
    "confidence": 0.75376886,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "github.com/spf13/pflag",
  "Dir": "../..//vendor/github.com/spf13/pflag",
  "Version": "v1.0.5",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/crypto",
  "Dir": "../..//vendor/golang.org/x/crypto",
  "Version": "v0.23.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/exp",
  "Dir": "../..//vendor/golang.org/x/exp",
  "Version": "v0.0.0-20240205201215-2c58cdc269a3",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/net",
  "Dir": "../..//vendor/golang.org/x/net",
  "Version": "v0.25.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/sync",
  "Dir": "../..//vendor/golang.org/x/sync",
  "Version": "v0.7.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/sys",
  "Dir": "../..//vendor/golang.org/x/sys",
  "Version": "v0.20.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/text",
  "Dir": "../..//vendor/golang.org/x/text",
  "Version": "v0.16.0",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "golang.org/x/tools",
  "Dir": "../..//vendor/golang.org/x/tools",
  "Version": "v0.21.1-0.20240508182429-e35e4ccd0d2d",
  "ErrStr": "does not meet minimum lead: 0.930693 - 0.884422 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-3-Clause",
    "confidence": 0.9306931,
    "file": "LICENSE"
   },
   {
    "license": "BSD-Source-Code",
    "confidence": 0.8844221,
    "file": "LICENSE"
   }
  ]
 },
 {
  "Name": "gopkg.in/warnings.v0",
  "Dir": "../..//vendor/gopkg.in/warnings.v0",
  "Version": "v0.1.2",
  "ErrStr": "does not meet minimum lead: 0.936782 - 0.900585 lt 0.050000",
  "Matches": [
   {
    "license": "BSD-2-Clause",
    "confidence": 0.9367816,
    "file": "LICENSE"
   },
   {
    "license": "BSD-1-Clause",
    "confidence": 0.9005848,
    "file": "LICENSE"
   },
   {
    "license": "BSD-3-Clause",
    "confidence": 0.7643678,
    "file": "LICENSE"
   }
  ]
 }]`

const VerifyErrOut = `[{
  "name": "github.com/davecgh/go-spew",
  "version": "v1.1.1",
  "err": "non-whitelisted license: 0BSD",
  "license": {
   "kind": "0BSD",
   "text": "ISC License\n\nCopyright (c) 2012-2016 Dave Collins \u003cdave@davec.name\u003e\n\nPermission to use, copy, modify, and/or distribute this software for any\npurpose with or without fee is hereby granted, provided that the above\ncopyright notice and this permission notice appear in all copies.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\" AND THE AUTHOR DISCLAIMS ALL WARRANTIES\nWITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF\nMERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR\nANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES\nWHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN\nACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF\nOR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.\n"
  }
 }]`
