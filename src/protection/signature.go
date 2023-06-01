package protection

import (
	"bytes"
	"os"
)

func AddSignature(f string) error {

	file, err := os.OpenFile(f, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	size, err := file.Stat()

	if err != nil {
		return err
	}

	offset := size.Size() / 2

	_, e := file.WriteAt([]byte("HYDROGEN"), offset)

	if e != nil {
		return e
	}

	return nil

}

func CheckSingature(f string) (bool, error) {
	file, err := os.OpenFile(f, 0, os.ModePerm)
	if err != nil {
		return false, err
	}
	defer file.Close()

	size, err := file.Stat()

	if err != nil {
		return false, err
	}

	_, err = file.Seek(size.Size()/2, 0)

	if err != nil {
		return false, err
	}

	signat := make([]byte, 8)

	_, err = file.Read(signat)

	if err != nil {
		return false, err
	}

	if bytes.Equal(signat, []byte{72, 89, 68, 82, 79, 71, 69, 78}) {
		return true, nil
	}

	return false, nil

}
