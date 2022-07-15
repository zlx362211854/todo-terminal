package util

type StrSlice []string

func (slice StrSlice) FindIndex(value string) int {
	for idx, v := range slice {
		if v == value {
			return idx
		}
	}
	return -1
}

func GetValueFromArgs(args StrSlice, key string) string {
	var idx = args.FindIndex(key)
	if idx >= 0 {
		return args[idx+1]
	}
	return ""
}
