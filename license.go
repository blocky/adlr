package adlr

type License struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

func MakeLicense(
	kind, text string,
) License {
	return License{Kind: kind, Text: text}
}
