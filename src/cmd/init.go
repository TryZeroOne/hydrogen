package cmd

import (
	"flag"
	"fmt"
	"os"
)

func help() {

	// fmt.Println("Usage: hydrogen <BUILD/PROTECT> <FILE TO PROTECT> -f <FLAGS>\n\nBuild examples:\nhydrogen build\t\t\t| Compiles the go code in the current directory\nhydrogen build -f jgc\t\t| Compiles the go code in the current directory with flags\n\nProtect examples:\nhydrogen protect f.bin\t\t| Protects f.bin\nhydrogen protect f.bin -f cg\t| Protects f.bin with flags\n\nBuild flags:\no\t| Obfuscate files\nt\t| Output\nj\t| Junk code\ng\t| Garbage\np\t| Custom entry point\nc\t| Compress\ne\t| Encrypt")

	fmt.Println(`
 Usage: hydrogen <BUILD/PROTECT> <FILE TO PROTECT/OUTPUT FILE> -f <FLAGS>

 Build examples:			Protect examples:
  hydrogen build build.exe		 hydrogen protect f.bin	
  hydrogen build build.exe -f jgc	 hydrogen protect f.bin -f cg


 Build flags:                           Protect flags:
  o   Obfuscate files                    c   Compress
  j   Junk code                          e   Encrypt
  g   Garbage                            g   Garbage
  p   Custom entry point
  c   Compress
  e   Encrypt
	`)

	os.Exit(1)
}

func Init() (string, string, string) {

	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	protectCmd := flag.NewFlagSet("protect", flag.ExitOnError)

	if len(os.Args) < 2 {
		help()
	}

	if os.Args[1] != "-help" && os.Args[1] != "build" && os.Args[1] != "protect" {
		help()
	}

	switch os.Args[1] {
	case "build":
		flags, output := HandleBuild(buildCmd)
		return flags, os.Args[1], output
	case "protect":
		flags, output := HandleProtect(protectCmd)
		return flags, os.Args[1], output
	case "-help":

		fmt.Println(`
 Usage: hydrogen <BUILD/PROTECT> <FILE TO PROTECT/OUTPUT FILE> -f <FLAGS>

 Build examples:			Protect examples:
  hydrogen build build.exe		 hydrogen protect f.bin	
  hydrogen build build.exe -f jgc	 hydrogen protect f.bin -f cg


 Build flags:                           Protect flags:
  o   Obfuscate files                    c   Compress
  j   Junk code                          e   Encrypt
  g   Garbage                            g   Garbage
  p   Custom entry point
  c   Compress
  e   Encrypt
	`)
		os.Exit(0)

	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
	}

	return "", "", ""
}

func HandleBuild(buildCmd *flag.FlagSet) (string, string) {

	if len(os.Args) <= 2 {
		help()
	}

	if len(os.Args) <= 4 {
		return "all", os.Args[2]
	}

	return os.Args[4], os.Args[2]

}

func HandleProtect(protectCmd *flag.FlagSet) (string, string) {
	if len(os.Args) <= 2 {
		help()
	}

	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Printf("%s does not exist. 'hydrogen -help' for help\n", os.Args[2])
		os.Exit(1)
	}
	f.Close()

	if len(os.Args) <= 4 {
		return "all", os.Args[2]
	}

	return os.Args[4], os.Args[2]

}
