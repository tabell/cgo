package cpals

import (
    "log"
    "os"
    "bufio"
)


func StringArrayFromFile(filename string) (e error, result []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
        return err, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        return err, nil
		log.Fatal(err)
	}

	return
}
