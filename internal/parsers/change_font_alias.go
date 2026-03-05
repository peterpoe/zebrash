package parsers

import (
	"path/filepath"
	"strings"

	"github.com/ingridhq/zebrash/internal/printers"
)

// NewChangeFontAliasParser parses the ^CW command which assigns a font file to a font alias.
// Format: ^CWa,d:filename.TTF
// Example: ^CW1,E:FONT.TTF  →  alias "1" = "FONT.TTF"
func NewChangeFontAliasParser() *CommandParser {
	const code = "^CW"

	return &CommandParser{
		CommandCode: code,
		Parse: func(command string, printer *printers.VirtualPrinter) (any, error) {
			// Strip the command code and split on the first comma to get alias char
			body := strings.TrimPrefix(command, code)
			if len(body) == 0 {
				return nil, nil
			}

			alias := strings.ToUpper(string(body[0]))
			rest := body[1:]

			// rest is ",d:filename.TTF" — find the colon to strip the device prefix
			colonIdx := strings.Index(rest, ":")
			var fileName string
			if colonIdx >= 0 {
				fileName = strings.TrimSpace(rest[colonIdx+1:])
			} else {
				// No device prefix — rest starts with comma then filename
				fileName = strings.TrimSpace(strings.TrimPrefix(rest, ","))
			}

			if fileName != "" {
				printer.FontAliases[alias] = filepath.Base(fileName)
			}

			return nil, nil
		},
	}
}
