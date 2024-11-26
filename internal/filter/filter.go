package filter

var validFilterTypes = [5]string{"name", "releaseDate", "group", "text", "link"}

type Filter struct {
	FType  string
	FValue string
	Skip   int
	Rows   int
}

func IsTypeEmpty(f Filter) bool {
	return f.FType == "" && f.FValue == ""
}

func IsClear(f Filter) bool {
	return f.FType == "" && f.FValue == "" && f.Rows == 0 && f.Skip == 0
}

func IsTypeValid(f Filter) bool {
	for _, filterType := range validFilterTypes {
		if f.FType == filterType {
			return true
		}
	}
	return false
}
