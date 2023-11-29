package ToFile

import (
	"encoding/json"
	"os"
)

type File struct {
	Path string
}

func (f *File) GetData() ([]byte, error) {
	path := f.Path

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.WriteFile(path, []byte("[]"), 0644)
			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

func (f *File) AddData(task string) error {
	data, err := f.GetData()
	if err != nil {
		return err
	}

	var decodedData []map[string]string
	err = f.SafeJsonUnmarshal(data, &decodedData)
	if err != nil {
		return err
	}

	var decodedTask map[string]string
	err = f.SafeJsonUnmarshal([]byte(task), &decodedTask)
	if err != nil {
		return err
	}

	decodedData = append(decodedData, decodedTask)

	encodedData, err := json.MarshalIndent(decodedData, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(f.Path, encodedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) DeleteData(id int) error {
	data, err := f.GetData()
	if err != nil {
		return err
	}

	var decodedData []map[string]string
	err = f.SafeJsonUnmarshal(data, &decodedData)
	if err != nil {
		return err
	}

	finalData := removeElementByIndex(decodedData, id)

	encodedData, err := json.MarshalIndent(finalData, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(f.Path, encodedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) SafeJsonUnmarshal(data []byte, writeTo any) error {
	if len(data) != 0 {
		err := json.Unmarshal(data, writeTo)
		if err != nil {
			return err
		}
	} else {
		err := os.WriteFile(f.Path, []byte("[]"), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func removeElementByIndex(slice []map[string]string, index int) []map[string]string {
    if index < 0 || index >= len(slice) {
        return slice
    }

    return append(slice[:index], slice[index+1:]...)
}