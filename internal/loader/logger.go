package loader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func write(fileName string, b []byte) error {
	res := bytes.Buffer{}

	err := json.Indent(&res, b, "  ", "  ")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fmt.Sprintf("%s.json", fileName), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(res.Bytes())

	return nil
}
