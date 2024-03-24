package prettyld

func Unmarshal(b any, dest any) error {
	list, err := Parse(b, nil)
	if err != nil {
		return err
	}
	list.UnmarshalTo(&dest)
	return nil
}
