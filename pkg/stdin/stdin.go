package stdin

import (
	"bufio"
	"os"
)

func ReadLines(callback func([]byte) error) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := callback(scanner.Bytes())
		if err != nil {
			return err
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
