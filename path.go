package SnowoUtils

import (
	"fmt"
	"os"
	"strings"
)

func GetExtension(filename string) string {
	if strings.Contains(filename,".") {
		a := strings.Split(filename,".")
		return a[len(a)-1]
	}
	return ""
}

func ConcatExtension(nameBody, ext string) string {
	if ext=="" {
		return nameBody
	}
	return fmt.Sprintf("%s.%s",nameBody,ext)
}

func FileSize(path string) (int64, error) {
	file, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return file.Size(),nil
}

