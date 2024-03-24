package prettyld

import "github.com/piprate/json-gold/ld"

func Unmarshal(b any, dest any, options *ld.JsonLdOptions) error {
	list, err := Parse(b, options)
	if err != nil {
		return err
	}
	list.UnmarshalTo(&dest)
	return nil
}
