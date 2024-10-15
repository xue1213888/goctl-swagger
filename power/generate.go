package power

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"io/ioutil"
)

func Do(filename string, in *plugin.Plugin) error {
	p := applyGenerate(in)
	var formatted bytes.Buffer
	cw := csv.NewWriter(&formatted)
	err := cw.WriteAll(p)
	if err != nil {
		return err
	}
	cw.Flush()
	output := in.Dir + "/" + filename
	err = ioutil.WriteFile(output, formatted.Bytes(), 0666)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
