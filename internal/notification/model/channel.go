package model

func (c Channel) IsValid() bool {
	switch c {

	case Email:
		return true

	case SMS:
		return true

	case Push:
		return true
	}
	return false
}
