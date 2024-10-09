package prettyprinter

import "encoding/json"

func errorToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func prettyJSON(i interface{}) ([]byte, error) {
	var prefix, indent string = "", " "
	return json.MarshalIndent(i, prefix, indent)
}

func print(bytes []byte, w Writer) error {
	_, err := w.Write(append(bytes, []byte("\n")...))
	return err
}
