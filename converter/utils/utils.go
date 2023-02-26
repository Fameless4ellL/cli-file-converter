package utils

func GetFileExtension(filename string) string {
	ext := ""
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			break
		}
		ext = string(filename[i]) + ext
	}
	return ext
}