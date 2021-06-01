package internal

// License holds license details of a golang module
// derived by licensedb and solved by Miner
type License struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

// Create a License
func MakeLicense(
	kind, text string,
) License {
	return License{Kind: kind, Text: text}
}
