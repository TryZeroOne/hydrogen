package protection

import (
	"encoding/hex"
	"errors"
	"os"
	"os/exec"
	"strings"
)

var UpxData = map[string]string{

	"496e666f3a20546869732066696c65206973207061636b6564207769746820746865205550582065786563757461626c65207061636b657220687474703a2f2f7570782e73662e6e6574": "496e666f3a20" + hex.EncodeToString([]byte(GenString(68))), // info

	"49643a2055505820342e303220436f707972696768742028432920313939362d323032332074686520555058205465616d2e20416c6c205269676874732052657365727665642e": "49643a20" + hex.EncodeToString([]byte(GenString(66))) + "2e", // id

	// // "5550784C": "4646664C",

	// "55505821": hex.EncodeToString([]byte(GenString(3))) + "21",
	"55505830": hex.EncodeToString([]byte(GenString(3))) + "30", // section
	// "557058":   hex.EncodeToString([]byte(GenString(3))),
	// "555058":   hex.EncodeToString([]byte(GenString(3))),

	// "55505831": hex.EncodeToString([]byte(GenString(3))) + "31", // section
	// "55505832": hex.EncodeToString([]byte(GenString(3))) + "32", 
}

func Compress(file string) error {

	e := exec.Command("upx", "-h").Run()
	if e != nil {
		return errors.New("Please install upx ( https://github.com/upx/upx )")
	}

	e = compress(file)

	if e != nil {
		return e
	}

	e = antiUnpack(file)

	if e != nil {
		return e
	}

	return nil
}

func antiUnpack(file string) error {

	f, err := os.ReadFile(file)
	if err != nil {
		return errors.New("Compress error: " + err.Error())
	}

	s := hex.EncodeToString(f)

	var new = s

	for x := range UpxData {
		new = strings.ReplaceAll(new, x, UpxData[x])
	}

	payload, err := hex.DecodeString(new)
	if err != nil {
		return errors.New("Compress error: " + err.Error())
	}

	err = os.WriteFile(file, payload, os.ModePerm)
	if err != nil {
		return errors.New("Compress error: " + err.Error())
	}
	cleanHeader(file)

	return nil
}

func compress(file string) error {

	cmd := exec.Command("upx", "--best", "--lzma", "--no-overlay", "--strip-relocs", file)
	var out strings.Builder
	env := os.Environ()
	cmd.Env = env
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return errors.New("Unknown compress error: " + err.Error())
	}

	e := antiUnpack(file)
	if e != nil {
		return e
	}

	return nil
}

// func GenerateHex(mode int) string {

// 	magic := GenString(3)

// 	return hex.EncodeToString([]byte(GenString(3)))

// }
