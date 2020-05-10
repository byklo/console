package console

import (
    "bufio"
    "fmt"
    "os"
    "github.com/labstack/gommon/color"
    "strings"
)

const (
    DEFAULT_EXIT_COMMAND = "exit"
    DEFAULT_HELP_COMMAND = "help"
    DEFAULT_PROMPT_SYMBOL = ">"
)

type Console struct {
    commands map[string]func([]string)
    helps []string
    exit_command string
    help_command string
    prompt_symbol string
    scanner *bufio.Scanner
}

func NewConsole() *Console {
    return &Console{
        make(map[string]func([]string)),
        make([]string, 0),
        DEFAULT_EXIT_COMMAND,
        DEFAULT_HELP_COMMAND,
        DEFAULT_PROMPT_SYMBOL,
        bufio.NewScanner(os.Stdin),
    }
}

func (console *Console) Run() {
    debug("console start")

    for {
        command_tokens := console.prompt()
        // debugf("got tokens: %q", command_tokens)

        if len(command_tokens) <= 0 {
            continue
        } else if command_tokens[0] == console.help_command {
            fmt.Println(strings.Join(console.helps, "\n"))
            continue
        } else if command_tokens[0] == console.exit_command {
            return
        }

        if action, exists := console.commands[command_tokens[0]]; exists {
            action(command_tokens[1:])
        } else {
            fmt.Printf("%s: command not found\n", command_tokens[0])
        }
    }
}

func (console *Console) Register(command string, help string, action func([]string)) {
    console.commands[command] = action
    console.helps = append(console.helps, help)
}

func (console *Console) prompt() []string {
    fmt.Printf("%s ", console.prompt_symbol)
    console.scanner.Scan()
    raw := console.scanner.Text()
    tokens := strings.Fields(raw)
    return tokens
}







func debug(message string) {
    fmt.Println(color.Yellow(fmt.Sprintf("[DEBUG] %s", message)))
}

func debugf(format string, a ...interface{}) {
    debug(fmt.Sprintf(format, a...))
}

