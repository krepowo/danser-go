package settings

var Gameplay = initGameplay()

func initGameplay() *gameplay {
	return &gameplay{
		HitErrorMeter: &hitError{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			PointFadeOutTime:     10,
			ShowPositionalMisses: true,
			PositionalMissScale:  1.5,
			ShowUnstableRate:     true,
			UnstableRateDecimals: 0,
			UnstableRateScale:    1.0,
			StaticUnstableRate:   false,
			ScaleWithSpeed:       false,
		},
		AimErrorMeter: &aimError{
			hudElementPosition: &hudElementPosition{
				hudElement: &hudElement{
					Show:    false,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XPosition: 1350,
				YPosition: 650,
			},
			PointFadeOutTime:     10,
			DotScale:             1,
			Align:                "Right",
			ShowUnstableRate:     false,
			UnstableRateScale:    1,
			UnstableRateDecimals: 0,
			StaticUnstableRate:   false,
			CapPositionalMisses:  true,
			AngleNormalized:      false,
			UseFallbackSkin:      false,
		},
		Score: &score{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			ProgressBar:     "Pie",
			ShowGradeAlways: false,
			StaticScore:     false,
			StaticAccuracy:  false,
		},
		HpBar: &hudElementOffset{
			hudElement: &hudElement{
				Show:    true,
				Scale:   1.0,
				Opacity: 1.0,
			},
			XOffset: 0,
			YOffset: 0,
		},
		ComboCounter: &comboCounter{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			Static: false,
		},
		PPCounter: &ppCounter{
			hudElementPosition: &hudElementPosition{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XPosition: 5,
				YPosition: 150,
			},
			Color: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Decimals:         0,
			Align:            "CentreLeft",
			ShowInResults:    true,
			ShowPPComponents: false,
			Static:           false,
		},
		HitCounter: &hitCounter{
			hudElementPosition: &hudElementPosition{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XPosition: 5,
				YPosition: 190,
			},
			Color300: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Color100: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Color50: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			ColorMiss: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			ColorSB: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Spacing:          48,
			FontScale:        1,
			Align:            "Left",
			ValueAlign:       "Left",
			Vertical:         false,
			Show300:          false,
			ShowSliderBreaks: false,
		},
		StrainGraph: &strainGraph{
			Show:      true,
			Opacity:   1,
			XPosition: 5,
			YPosition: 310,
			Align:     "BottomLeft",
			Width:     130,
			Height:    70,
			BgColor: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      0.2,
			},
			FgColor: &HSV{
				Hue:        297,
				Saturation: 0.4,
				Value:      0.92,
			},
			Outline: &outline{
				Show:          false,
				Width:         2,
				InnerDarkness: 0.5,
				InnerOpacity:  0.5,
			},
		},
		KeyOverlay: &keyOverlay{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			Mode:      "osu",
			HoldBar:   true,
			KeySize:   42, // + 2*Outline = 46px box, the same size as inputoverlay-key.png
			Outline:   2,
			Margin:    24, // same distance from the edge as the osu! key strip
			BarSpeed:  300,
			BarLength: 300,
			Fade:      true,
			Counter:   false,
		},
		ScoreBoard: &scoreBoard{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			Mode:           "Normal",
			ModsOnly:       false,
			AlignRight:     false,
			HideOthers:     false,
			ShowAvatars:    false,
			ExplosionScale: 1.0,
		},
		Mods: &mods{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			HideInReplays:     false,
			FoldInReplays:     false,
			ShowLazerMod:      true,
			AdditionalSpacing: 0,
		},
		Boundaries: &boundaries{
			Enabled:         true,
			BorderThickness: 1,
			BorderFill:      1,
			BorderColor: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			BorderOpacity: 1,
			BackgroundColor: &HSV{
				Hue:        0,
				Saturation: 1,
				Value:      0,
			},
			BackgroundOpacity: 0.5,
		},
		Underlay: &underlay{
			Path:       "",
			AboveHpBar: false,
		},
		Statistics:              make([]*Statistic, 0),
		SBFont:                  "",
		HUDFont:                 "",
		ShowResultsScreen:       true,
		ResultsScreenTime:       5,
		ResultsUseLocalTimeZone: false,
		ShowWarningArrows:       true,
		ShowHitLighting:         false,
		FlashlightDim:           1,
		PlayUsername:            "Guest",
		IgnoreFailsInReplays:    false,
		PPVersion:               "latest",
		LazerClassicScore:       false,
		AlwaysSkipIntro:         false,
	}
}

type gameplay struct {
	HitErrorMeter           *hitError
	AimErrorMeter           *aimError
	Score                   *score
	HpBar                   *hudElementOffset
	ComboCounter            *comboCounter
	PPCounter               *ppCounter
	HitCounter              *hitCounter
	StrainGraph             *strainGraph
	KeyOverlay              *keyOverlay
	ScoreBoard              *scoreBoard
	Mods                    *mods
	Boundaries              *boundaries
	Underlay                *underlay
	Statistics              []*Statistic `new:"InitStatistic" minSize:"0" wiki:"Help|https://github.com/Wieku/danser-go/wiki/Templates"`
	SBFont                  string       `label:"Scoreboard / Ranking font" file:"Select SBR font" filter:"TrueType/OpenType Font (*.ttf, *.otf)|ttf,otf" tooltip:"Sets the font that will be used for score board names and ranking panel (use Aller Light to match osu!)" liveedit:"false"`
	HUDFont                 string       `label:"Overlay (HUD) font" file:"Select HUD font" filter:"TrueType/OpenType Font (*.ttf, *.otf)|ttf,otf" tooltip:"Sets the font that will be used for PP/UR/hit counts" liveedit:"false"`
	ShowResultsScreen       bool         `liveedit:"false"`
	ResultsScreenTime       float64      `label:"Results screen duration" min:"1" max:"20" format:"%.1fs" liveedit:"false"`
	ResultsUseLocalTimeZone bool         `label:"Show PC's time zone instead of UTC"`
	ShowWarningArrows       bool
	ShowHitLighting         bool
	FlashlightDim           float64
	PlayUsername            string `liveedit:"false"`
	IgnoreFailsInReplays    bool
	PPVersion               string `liveedit:"false" label:"PP counter version" combo:"211112|2021 pp rework (First Xexxar),220930|2022 pp rework,241007|2024 pp rework,250306|2025 Q1 update,251020|2025 Q4 update,latest|2026 Q3 update (latest)"`
	LazerClassicScore       bool   `label:"Use \"Classic\" score for osu!lazer plays"`
	AlwaysSkipIntro         bool   `liveedit:"false"`
}

type boundaries struct {
	Enabled bool

	BorderThickness float64 `min:"0.5" max:"10" format:"%.1f o!px"`
	BorderFill      float64

	BorderColor   *HSV    `short:"true"`
	BorderOpacity float64 `scale:"100.0" format:"%.0f%%"`

	BackgroundColor   *HSV    `short:"true"`
	BackgroundOpacity float64 `scale:"100.0" format:"%.0f%%"`
}

type hudElement struct {
	Show    bool
	Scale   float64 `max:"3" scale:"100.0" format:"%.0f%%"`
	Opacity float64 `scale:"100.0" format:"%.0f%%"`
}

type hudElementOffset struct {
	*hudElement
	offset  string  `vector:"true" left:"XOffset" right:"YOffset"`
	XOffset float64 `min:"-10000" max:"10000"`
	YOffset float64 `min:"-10000" max:"10000"`
}

type hudElementPosition struct {
	*hudElement
	position  string  `vector:"true" left:"XPosition" right:"YPosition"`
	XPosition float64 `min:"-10000" max:"10000"`
	YPosition float64 `min:"-10000" max:"10000"`
}

type keyOverlay struct {
	*hudElementOffset
	Mode     string  `combo:"osu|osu! style (skin strip),stream|Stream overlay (bars travelling up)" tooltip:"osu! style uses the skin's key strip with press counters. Stream overlay mimics Blondazz/KeyOverlay: squares with a letter, bars that travel upwards while a key is held."`
	HoldBar  bool    `label:"Show key hold duration bar" tooltip:"osu! style only: draws a bar next to a held key, its length proportional to how long the key has been down"`
	KeySize   float64 `min:"20" max:"400" format:"%.0f px" label:"Key size (stream overlay)"`
	Outline   float64 `min:"0" max:"40" format:"%.0f px" label:"Border thickness (stream overlay)"`
	Margin    float64 `min:"0" max:"600" format:"%.0f px" label:"Margin from screen edge (stream overlay)"`
	BarSpeed  float64 `min:"50" max:"6000" format:"%.0f px/s" label:"Bar speed (stream overlay)"`
	BarLength float64 `min:"0" max:"2000" format:"%.0f px" label:"Bar length (stream overlay)" tooltip:"How far a bar travels before it disappears. 0 = all the way to the top of the screen."`
	Fade      bool    `label:"Fade bars near the top (stream overlay)"`
	Counter   bool    `label:"Show press counter (stream overlay)"`
}

type hitError struct {
	*hudElementOffset
	PointFadeOutTime     float64 `max:"10" format:"%.1fs"`
	ShowPositionalMisses bool
	PositionalMissScale  float64 `min:"1" max:"2" scale:"100" format:"%.0f%%"`
	ShowUnstableRate     bool
	UnstableRateDecimals int     `max:"5"`
	UnstableRateScale    float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	StaticUnstableRate   bool
	ScaleWithSpeed       bool
}

type aimError struct {
	*hudElementPosition
	PointFadeOutTime     float64 `max:"10" format:"%.1fs"`
	DotScale             float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	Align                string  `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	ShowUnstableRate     bool
	UnstableRateScale    float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	UnstableRateDecimals int     `max:"5"`
	StaticUnstableRate   bool
	CapPositionalMisses  bool
	AngleNormalized      bool
	UseFallbackSkin      bool `tooltip:"Useful for \"instafade\" skins" liveedit:"false"`
}

type score struct {
	*hudElementOffset
	ProgressBar     string `combo:"Pie,Bar,BottomRight,Bottom"`
	ShowGradeAlways bool   `label:"Always show grade"`
	StaticScore     bool
	StaticAccuracy  bool
}

type comboCounter struct {
	*hudElementOffset
	Static bool
}

type ppCounter struct {
	*hudElementPosition
	Color            *HSV   `short:"true"`
	Decimals         int    `max:"5"`
	Align            string `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	ShowInResults    bool
	ShowPPComponents bool `label:"Show PP breakdown"`
	Static           bool
}

type hitCounter struct {
	*hudElementPosition
	Color            []*HSV  `json:",omitempty" new:"InitHSV" label:"Color list" skip:"true"`
	Color300         *HSV    `label:"Color of 300s"`
	Color100         *HSV    `label:"Color of 100s"`
	Color50          *HSV    `label:"Color of 50s"`
	ColorMiss        *HSV    `label:"Color of misses"`
	ColorSB          *HSV    `label:"Color of slider breaks"`
	Spacing          float64 `string:"true" min:"0" max:"1366"`
	FontScale        float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	Align            string  `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	ValueAlign       string  `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	Vertical         bool
	Show300          bool `label:"Show perfect hits"`
	ShowSliderBreaks bool
}

type scoreBoard struct {
	*hudElementOffset
	Mode           string `combo:"Normal,Country,Friends,Team" tooltip:"Team mode requires Authorization Code API Mode. Country and Friends modes require osu!supporter and Authorization Code API Mode!"`
	ModsOnly       bool   `label:"Show mod leaderboard"`
	AlignRight     bool   `label:"Align to the right" label:"Simulates the second team of osu! multiplayer"`
	HideOthers     bool
	ShowAvatars    bool
	ExplosionScale float64 `min:"0.1" max:"2" scale:"100" format:"%.0f%%"`
}

type mods struct {
	*hudElementOffset
	HideInReplays     bool
	FoldInReplays     bool
	ShowLazerMod      bool
	AdditionalSpacing float64 `string:"true" min:"-1366" max:"1366"`
}

type strainGraph struct {
	Show    bool
	Opacity float64 `scale:"100.0" format:"%.0f%%"`

	position  string  `vector:"true" left:"XPosition" right:"YPosition"`
	XPosition float64 `min:"-10000" max:"10000"`
	YPosition float64 `min:"-10000" max:"10000"`

	Align string `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`

	size   string  `vector:"true" left:"Width" right:"Height"`
	Width  float64 `string:"true" min:"1" max:"10000"`
	Height float64 `string:"true" min:"1" max:"768"`

	BgColor *HSV `label:"Background color" short:"true"`
	FgColor *HSV `label:"Foreground color" short:"true"`

	Outline *outline
}

type outline struct {
	Show          bool
	Width         float64 `min:"1" max:"5"`
	InnerDarkness float64 `scale:"100.0" format:"%.0f%%" tooltip:"Darkness of filled shape, only applicable when DrawOutline is enabled"`
	InnerOpacity  float64 `scale:"100.0" format:"%.0f%%" tooltip:"Opacity of filled shape, only applicable when DrawOutline is enabled"`
}

type underlay struct {
	Path       string `file:"Select underlay image" filter:"PNG file (*.png)|png" tooltip:"PNG file that will be used as HUD background (similar to custom HP bar backgrounds). It's scaled automatically to fit the screen vertically" liveedit:"false"`
	AboveHpBar bool   `label:"Show underlay above HP bar" tooltip:"Use this if HP bar background is large"`
}

type Statistic struct {
	Show bool

	Template string `long:"true" multi:"true"`

	position  string  `vector:"true" left:"XPosition" right:"YPosition"`
	XPosition float64 `min:"-10000" max:"10000"`
	YPosition float64 `min:"-10000" max:"10000"`

	Align  string `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	Anchor string `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`

	Size    float64 `min:"3" max:"128" format:"%.0fo!px"`
	Opacity float64 `scale:"100.0" format:"%.0f%%"`
	Color   *HSV
}

func (d *defaultsFactory) InitStatistic() *Statistic {
	return &Statistic{
		Show:      true,
		Template:  "{{.pp}}PP",
		XPosition: 100,
		YPosition: 100,
		Align:     "TopLeft",
		Anchor:    "TopLeft",
		Size:      24,
		Opacity:   1,
		Color: &HSV{
			Hue:        0,
			Saturation: 0,
			Value:      1,
		},
	}
}
