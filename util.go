package topfee

import (
	"io/ioutil"
	"strings"
)

func numFilesInDir(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return -1
	}
	return len(files)
}

func readColonFile(path string) (map[string]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")

	values := make(map[string]string)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		//fmt.Println(strconv.Itoa(i) + "++ " + line)
		parts := strings.Split(string(line), ":")
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		values[key] = value
	}

	return values, nil

}
