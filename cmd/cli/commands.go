package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/viktorfrom/d7024e-kademlia/internal/kademlia"
)

const (
	errNoArg       string = "No argument!"
	errInvalidCmd  string = "Invalid command!"
	errNoFileFound string = "Could not find or open file: "
)

var (
	osExit   = os.Exit
	logFatal = log.Fatal
	helpFile = Prompt()
)

// Commands handles the commands of the CLI. `output` is the io.Writer to output data to.
// `node` is the Kademlia node this CLI runs for. `commands` a list of program commands.
func Commands(output io.Writer, node *kademlia.Node, commands []string) {

	switch commands[0] {
	case "put":
		if len(commands) == 2 {
			Put(*node, commands[1])
		} else {
			fmt.Fprintln(output, errNoArg)
		}
	case "p":
		if len(commands) == 2 {
			Put(*node, commands[1])
		} else {
			fmt.Fprintln(output, errNoArg)
		}
	case "get":
		if len(commands) == 2 {
			Get(*node, commands[1])
		} else {
			fmt.Fprintln(output, errNoArg)
		}
	case "g":
		if len(commands) == 2 {
			Get(*node, commands[1])
		} else {
			fmt.Fprintln(output, errNoArg)
		}
	case "info":
		fmt.Println("ID: ", node.RT.GetMeID())
	case "exit":
		Exit()
	case "e":
		Exit()
	case "help":
		Help(output)
	case "h":
		Help(output)
	default:
		fmt.Fprintln(output, errInvalidCmd)
	}
}

func Put(node kademlia.Node, input string) {
	hash := node.StoreValue(input)
	println("Hash = ", hash)
}

func Get(node kademlia.Node, hash string) {
	value, err := node.FindValue(hash)

	if err != nil {
		println(err.Error())
	} else {
		println("value = ", value)
	}
}

func Exit() {
	osExit(3)
}

func Help(output io.Writer) {
	text := Prompt()
	fmt.Fprintln(output, text)
}
