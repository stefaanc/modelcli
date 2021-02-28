package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

var version string = "0.0.1"

func main() {

    nargs := len(os.Args) - 1
    if nargs == 0 {
        printHelp()

    } else {
        verb := os.Args[1]
        switch verb {

        case "version":
            if nargs > 1 {
                versionFlags := flag.NewFlagSet("version", flag.ContinueOnError)
                versionFlags.Usage = func(){}
                versionHelp := versionFlags.Bool("help", false, "get help")
                versionH := versionFlags.Bool("h", false, "get help")

                err := versionFlags.Parse(os.Args[2:])
                if err != nil {
                    log.Fatalf("\"modelcli version\": found unexpected command-line options, try \"modelcli version --help\"\n")
                }

                if *versionHelp || *versionH {
                    printVersionHelp()
                } else {
                    log.Fatalf("\"modelcli version\": found unexpected command-line arguments, try \"modelcli version --help\"\n")
                }
            } else {
                fmt.Printf("modelcli version: %s\n", version)
            }

        // case "create":
        //     switch os.Args[2] {
        //     case "application":
        //         fmt.Println("modelcli create application")

        //     default:
        //         fmt.Println("modelcli create")
        //     }

        // case "read":
        //     switch os.Args[2] {
        //     case "application":
        //         fmt.Println("modelcli read application")

        //     default:
        //         fmt.Println("modelcli read")
        //     }

        // case "update":
        //     switch os.Args[2] {
        //     case "application":
        //         fmt.Println("modelcli update application")

        //     default:
        //         fmt.Println("modelcli update")
        //     }

        // case "patch":
        //     switch os.Args[2] {
        //     case "application":
        //         fmt.Println("modelcli patch application")

        //     default:
        //         fmt.Println("modelcli patch")
        //     }

        // case "delete":
        //     switch os.Args[2] {
        //     case "application":
        //         fmt.Println("modelcli delete application")

        //     default:
        //         fmt.Println("modelcli delete")
        //     }

        default:
            printHelp()
        }
    }
}

func printHelp() {
    fmt.Printf("\"modelcli {verb} [arguments..]\"\n")
    fmt.Printf("\n")
    fmt.Printf("accepted verbs:\n")
    fmt.Printf("\n")
    fmt.Printf("   \"version\" - get version information\n")
    fmt.Printf("   \n")
    fmt.Printf("   \"create\"  - create a new object\n")
    fmt.Printf("   \"read\"    - read an object\n")
    fmt.Printf("   \"update\"  - update an object\n")
    fmt.Printf("   \"patch\"   - patch an object\n")
    fmt.Printf("   \"delete\"  - delete an object\n")
    fmt.Printf("   \n")
    fmt.Printf("   \"import\"  - import objects from another store\n")
    fmt.Printf("   \"export\"  - export objects to another store\n")
    fmt.Printf("\n")
    fmt.Printf("use \"modelcli [verb] --help\" for more information\n")
}

func printVersionHelp() {
    fmt.Printf("\"modelcli version [options..]\"\n")
    fmt.Printf("\n")
    fmt.Printf("accepted options:\n")
    fmt.Printf("   \n")
    fmt.Printf("   \"--help\" (\"-h\") - get this help information\n")
    fmt.Printf("\n")
    fmt.Printf("get version information\n")
}
