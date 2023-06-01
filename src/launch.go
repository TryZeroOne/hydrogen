package src

import (
	"strings"

	"github.com/TryZeroOne/hydrogen/src/protection"
	"github.com/TryZeroOne/hydrogen/src/styles"
)

func Launch(flags, action, output string) {

	if action == "protect" {
		protected, err := protection.CheckSingature(output)

		if err != nil {
			styles.FatalError(err.Error())
			return
		}

		if protected {
			styles.FatalError("Already protected")
			return
		}
	}

	var compile, obf, junk, garbage, compress, encrypt, entry bool

	if action == "build" {
		compile = true
	}

	if flags == "all" && action == "build" {
		compile = true
		obf = true
		junk = true
		garbage = true
		compress = true
		encrypt = true
		// entry = true
	}
	if flags == "all" && action == "protect" {
		garbage = true
		compress = true
		encrypt = true
	} else if flags != "all" {
		obf, junk, garbage, compress, encrypt, entry = ParseFlags(flags, action)
	}

	styles.ProgressMenu(compile, obf, output, junk, garbage, compress, encrypt)

	protection.CopyFiles("./")
	if obf {
		styles.Stages++
		protection.ObfuscateFiles("./" + protection.TempName)
	}

	if junk {
		styles.Stages++
		protection.GetMain("./" + protection.TempName)
		err := protection.JunkCode()
		if err != nil {
			styles.FatalError(err.Error())
		}
	}

	if compile {

		protection.GetMain("./" + protection.TempName)

		styles.Stages++
		err := protection.Compile(output, entry)
		if err != nil {
			styles.FatalError(err.Error())
		}

	}

	if compress {
		styles.Stages++

		err := protection.Compress(output)
		if err != nil {
			styles.FatalError(err.Error())
		}

	}

	if encrypt {
		styles.Stages++
		err := protection.Encryption(output)
		if err != nil {
			styles.FatalError(err.Error())
		}
	}

	if garbage {
		styles.Stages++
		err := protection.Garbage(output)
		if err != nil {
			styles.FatalError(err.Error())
		}
	}

	if action == "protect" {
		err := protection.AddSignature(output)
		if err != nil {
			styles.FatalError("Signature: " + err.Error())
		}
	}

}

// 	styles.ProgressMenu(compile, obf, output, junk, garbage, compress, encrypt)

// compile, obf, junk, garbage, compress, encrypt,junk

/*
Build flags:                           Protect flags:

	o   Obfuscate files                    c   Compress
	j   Junk code                          e   Encrypt
	g   Garbage                            g   Garbage
	p   Custom entry point
	c   Compress
	e   Encrypt
*/
func ParseFlags(flags, action string) (bool, bool, bool, bool, bool, bool) {

	var actions = make(map[string]bool, 0)
	if action == "protect" {

		if strings.Contains(flags, "c") {
			actions["compress"] = true
		}
		if strings.Contains(flags, "e") {
			actions["encrypt"] = true
		}
		if strings.Contains(flags, "g") {
			actions["garbage"] = true
		}
		return false, false, actions["garbage"], actions["compress"], actions["encrypt"], false

	} else {
		// compile, obf, junk, garbage, compress, encrypt,junk

		/*
				Build flags:                           Protect flags:

			o   Obfuscate files                    c   Compress
			j   Junk code                          e   Encrypt
			g   Garbage                            g   Garbage
			p   Custom entry point
			c   Compress
			e   Encrypt
		*/

		if strings.Contains(flags, "o") {
			actions["obfuscate"] = true
		}
		if strings.Contains(flags, "j") {
			actions["junk"] = true
		}
		if strings.Contains(flags, "g") {
			actions["garbage"] = true
		}
		if strings.Contains(flags, "p") {
			actions["entry"] = true
		}
		if strings.Contains(flags, "e") {
			actions["encrypt"] = true
		}
		if strings.Contains(flags, "c") {
			actions["compress"] = true
		}

		return actions["obfuscate"], actions["junk"], actions["garbage"], actions["compress"], actions["encrypt"], actions["entry"]
	}

}
