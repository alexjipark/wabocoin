package tx

import (
	"github.com/spf13/viper"
	"io"
	"os"
	"io/ioutil"
	"github.com/pkg/errors"
	"encoding/json"
)

func LoadJSON (template interface{}) (bool, error) {
	input := viper.GetString(InputFlag)
	if input == "" {
		return true, nil
	}

	raw, err := readInput(input)

	// parse the input
	err = json.Unmarshal(raw, template)
	if err != nil {
		return true, err
	}
	return true, nil


}


func readInput(file string) ([]byte, error) {
	var reader io.Reader
	if file == "-" {
		reader = os.Stdin
	} else {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		reader = f
	}

	data, err := ioutil.ReadAll(reader)
	return data, errors.WithStack(err)
}