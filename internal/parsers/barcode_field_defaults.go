package parsers

import (
	"strconv"
	"strings"

	"github.com/peterpoe/zebrash/internal/printers"
)

func NewBarcodeFieldDefaults() *CommandParser {
	const code = "^BY"

	return &CommandParser{
		CommandCode: code,
		Parse: func(command string, printer *printers.VirtualPrinter) (any, error) {
			parts := splitCommand(command, code, 0)
			if len(parts) > 0 {
				if v, err := strconv.Atoi(parts[0]); err == nil {
					printer.DefaultBarcodeDimensions.ModuleWidth = v
				}
			}

			if len(parts) > 1 {
				if v, err := strconv.ParseFloat(strings.Trim(parts[1], " "), 32); err == nil {
					if v >= 2.0 && v <= 3.0 {
						printer.DefaultBarcodeDimensions.WidthRatio = v
					}
					// else: out-of-range value, keep the current default (3.0)
				}
			}

			if len(parts) > 2 {
				if v, err := strconv.Atoi(strings.Trim(parts[2], " ")); err == nil {
					printer.DefaultBarcodeDimensions.Height = v
				}
			}

			return nil, nil
		},
	}
}
