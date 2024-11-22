package filter

var validFilterTypes = [5]string{"name", "releaseDate", "group", "text", "link"}

type Filter struct {
	FType  string
	FValue string
	Skip   int
	Rows   int
}

func IsEmpty(f Filter) bool {
	return f.FType == "" && f.FValue == ""
}

func IsTypeValid(f Filter) bool {
	for _, filterType := range validFilterTypes {
		if f.FType == filterType {
			return true
		}
	}
	return false
}
