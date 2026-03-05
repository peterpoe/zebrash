package drawers

type DrawerOptions struct {
	LabelWidthMm  float64
	LabelHeightMm float64
	Dpmm          int
	// Render labels with inverted orientation upside-down
	EnableInvertedLabels bool
	// Custom fonts keyed by filename
	CustomFonts map[string][]byte
	// Scale factor applied to all font sizes (default 1.0 = no scaling)
	FontScale float64
}

func (d DrawerOptions) WithDefaults() DrawerOptions {
	res := d

	// by default produce 4x8 inches 203 dpi label

	if res.LabelWidthMm == 0 {
		res.LabelWidthMm = 101.6
	}

	if res.LabelHeightMm == 0 {
		res.LabelHeightMm = 203.2
	}

	if res.Dpmm == 0 {
		res.Dpmm = 8
	}

	return res
}
