package exists

import "os"

func File(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
