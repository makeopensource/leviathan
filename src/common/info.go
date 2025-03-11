package common

import (
	"fmt"
	"math"
	"math/rand/v2"
	"runtime"
	"sort"
	"strings"
	"time"
)

// build args to modify these vars
//
// go build -ldflags "\
//  -X github.com/makeopensource/leviathan/common.Version=0.1.0 \
//  -X github.com/makeopensource/leviathan/common.CommitInfo=$(git rev-parse HEAD) \
//  -X github.com/makeopensource/leviathan/common.BuildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
//  -X github.com/makeopensource/leviathan/common.Platform=$(go env GOOS)/$(go env GOARCH) \
//  -X github.com/makeopensource/leviathan/common.Branch=$(git rev-parse --abbrev-ref HEAD) \
//  "
// main.go

var Version = "dev"
var CommitInfo = "unknown"
var BuildDate = "unknown"
var Branch = "unknown" // Git branch
var SourceHash = "unknown"
var GoVersion = runtime.Version()

func PrintInfo() {
	// generated from https://patorjk.com/software/taag/#p=testall&t=leviathan
	var headers = []string{
		// contains some characters that mess with multiline strings leave this alone
		"\n          (`-')  _      (`-')  _     (`-')  _ (`-')      (`-').-> (`-')  _ <-. (`-')_ \n   <-.    ( OO).-/     _(OO ) (_)    (OO ).-/ ( OO).->   (OO )__  (OO ).-/    \\( OO) )\n ,--. )  (,------.,--.(_/,-.\\ ,-(`-')/ ,---.  /    '._  ,--. ,'-' / ,---.  ,--./ ,--/ \n |  (`-') |  .---'\\   \\ / (_/ | ( OO)| \\ /`.\\ |'--...__)|  | |  | | \\ /`.\\ |   \\ |  | \n |  |OO )(|  '--.  \\   /   /  |  |  )'-'|_.' |`--.  .--'|  `-'  | '-'|_.' ||  . '|  |)\n(|  '__ | |  .--' _ \\     /_)(|  |_/(|  .-.  |   |  |   |  .-.  |(|  .-.  ||  |\\    | \n |     |' |  `---.\\-'\\   /    |  |'->|  | |  |   |  |   |  | |  | |  | |  ||  | \\   | \n `-----'  `------'    `-'     `--'   `--' `--'   `--'   `--' `--' `--' `--'`--'  `--' \n",
		"\n __       ______   __   __   ________  ________   _________  ___   ___   ________   ___   __      \n/_/\\     /_____/\\ /_/\\ /_/\\ /_______/\\/_______/\\ /________/\\/__/\\ /__/\\ /_______/\\ /__/\\ /__/\\    \n\\:\\ \\    \\::::_\\/_\\:\\ \\\\ \\ \\\\__.::._\\/\\::: _  \\ \\\\__.::.__\\/\\::\\ \\\\  \\ \\\\::: _  \\ \\\\::\\_\\\\  \\ \\   \n \\:\\ \\    \\:\\/___/\\\\:\\ \\\\ \\ \\  \\::\\ \\  \\::(_)  \\ \\  \\::\\ \\   \\::\\/_\\ .\\ \\\\::(_)  \\ \\\\:. `-\\  \\ \\  \n  \\:\\ \\____\\::___\\/_\\:\\_/.:\\ \\ _\\::\\ \\__\\:: __  \\ \\  \\::\\ \\   \\:: ___::\\ \\\\:: __  \\ \\\\:. _    \\ \\ \n   \\:\\/___/\\\\:\\____/\\\\ ..::/ //__\\::\\__/\\\\:.\\ \\  \\ \\  \\::\\ \\   \\: \\ \\\\::\\ \\\\:.\\ \\  \\ \\\\. \\`-\\  \\ \\\n    \\_____\\/ \\_____\\/ \\___/_( \\________\\/ \\__\\/\\__\\/   \\__\\/    \\__\\/ \\::\\/ \\__\\/\\__\\/ \\__\\/ \\__\\/\n                                                                                                  \n",
		`
 ___        _______    ___      ___  ___   _________   ___  ___   ________   ________      
|\  \      |\  ___ \  |\  \    /  /||\  \ |\___   ___\|\  \|\  \ |\   __  \ |\   ___  \    
\ \  \     \ \   __/| \ \  \  /  / /\ \  \\|___ \  \_|\ \  \\\  \\ \  \|\  \\ \  \\ \  \   
 \ \  \     \ \  \_|/__\ \  \/  / /  \ \  \    \ \  \  \ \   __  \\ \   __  \\ \  \\ \  \  
  \ \  \____ \ \  \_|\ \\ \    / /    \ \  \    \ \  \  \ \  \ \  \\ \  \ \  \\ \  \\ \  \ 
   \ \_______\\ \_______\\ \__/ /      \ \__\    \ \__\  \ \__\ \__\\ \__\ \__\\ \__\\ \__\
    \|_______| \|_______| \|__|/        \|__|     \|__|   \|__|\|__| \|__|\|__| \|__| \|__|
`,
		`
      ___       ___           ___                       ___           ___           ___           ___           ___     
     /\__\     /\  \         /\__\          ___        /\  \         /\  \         /\__\         /\  \         /\__\    
    /:/  /    /::\  \       /:/  /         /\  \      /::\  \        \:\  \       /:/  /        /::\  \       /::|  |   
   /:/  /    /:/\:\  \     /:/  /          \:\  \    /:/\:\  \        \:\  \     /:/__/        /:/\:\  \     /:|:|  |   
  /:/  /    /::\~\:\  \   /:/__/  ___      /::\__\  /::\~\:\  \       /::\  \   /::\  \ ___   /::\~\:\  \   /:/|:|  |__ 
 /:/__/    /:/\:\ \:\__\  |:|  | /\__\  __/:/\/__/ /:/\:\ \:\__\     /:/\:\__\ /:/\:\  /\__\ /:/\:\ \:\__\ /:/ |:| /\__\
 \:\  \    \:\~\:\ \/__/  |:|  |/:/  / /\/:/  /    \/__\:\/:/  /    /:/  \/__/ \/__\:\/:/  / \/__\:\/:/  / \/__|:|/:/  /
  \:\  \    \:\ \:\__\    |:|__/:/  /  \::/__/          \::/  /    /:/  /           \::/  /       \::/  /      |:/:/  / 
   \:\  \    \:\ \/__/     \::::/__/    \:\__\          /:/  /     \/__/            /:/  /        /:/  /       |::/  /  
    \:\__\    \:\__\        ~~~~         \/__/         /:/  /                      /:/  /        /:/  /        /:/  /   
     \/__/     \/__/                                   \/__/                       \/__/         \/__/         \/__/    
`,
		`
██╗     ███████╗██╗   ██╗██╗ █████╗ ████████╗██╗  ██╗ █████╗ ███╗   ██╗
██║     ██╔════╝██║   ██║██║██╔══██╗╚══██╔══╝██║  ██║██╔══██╗████╗  ██║
██║     █████╗  ██║   ██║██║███████║   ██║   ███████║███████║██╔██╗ ██║
██║     ██╔══╝  ╚██╗ ██╔╝██║██╔══██║   ██║   ██╔══██║██╔══██║██║╚██╗██║
███████╗███████╗ ╚████╔╝ ██║██║  ██║   ██║   ██║  ██║██║  ██║██║ ╚████║
╚══════╝╚══════╝  ╚═══╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
`,
		`
 _        _______          _________ _______ _________          _______  _       
( \      (  ____ \|\     /|\__   __/(  ___  )\__   __/|\     /|(  ___  )( (    /|
| (      | (    \/| )   ( |   ) (   | (   ) |   ) (   | )   ( || (   ) ||  \  ( |
| |      | (__    | |   | |   | |   | (___) |   | |   | (___) || (___) ||   \ | |
| |      |  __)   ( (   ) )   | |   |  ___  |   | |   |  ___  ||  ___  || (\ \) |
| |      | (       \ \_/ /    | |   | (   ) |   | |   | (   ) || (   ) || | \   |
| (____/\| (____/\  \   /  ___) (___| )   ( |   | |   | )   ( || )   ( || )  \  |
(_______/(_______/   \_/   \_______/|/     \|   )_(   |/     \||/     \||/    )_) 
`,
	}

	const (
		width      = 90
		colorReset = "\033[0m"
		// Nord color palette ANSI equivalents
		nord4  = "\033[38;5;188m" // Snow Storm (darkest) - main text color
		nord8  = "\033[38;5;110m" // Frost - light blue
		nord9  = "\033[38;5;111m" // Frost - blue
		nord10 = "\033[38;5;111m" // Frost - deep blue
		nord15 = "\033[38;5;139m" // Aurora - purple
	)

	// Print header
	dividerContent := strings.Repeat("=", width)
	divider := nord9 + dividerContent + colorReset

	fmt.Println(divider)
	fmt.Printf("%s%s %s %s\n", nord15, strings.Repeat(" ", (width-24)/2), headers[rand.IntN(len(headers))], colorReset)
	fmt.Println(divider)

	// Print app info with aligned values
	printField := func(name, value string) {
		fmt.Printf("%s%-15s: %s%s%s\n", nord4, name, nord8, value, colorReset)
	}

	printField("Version", Version)
	printField("BuildDate", formatTime(BuildDate))
	printField("Branch", Branch)
	printField("CommitInfo", CommitInfo)
	printField("Source Hash", SourceHash)
	printField("GoVersion", GoVersion)

	//nolint
	if Branch != "unknown" && CommitInfo != "unknown" {
		fmt.Println(nord10 + strings.Repeat("-", width) + colorReset)
		var baserepo = fmt.Sprintf("https://github.com/makeopensource/leviathan")
		branchURL := fmt.Sprintf("%s/tree/%s", baserepo, Branch)
		commitURL := fmt.Sprintf("%s/commit/%s", baserepo, CommitInfo)

		printField("Repo", baserepo)
		printField("Branch", branchURL)
		printField("Commit", commitURL)
	}

	fmt.Println(divider)
}

func formatTime(input string) string {
	buildTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		//fmt.Printf("Error parsing build time: %v\n", err)
		return input
	}
	// Get the local timezone
	localLocation, err := time.LoadLocation("Local")
	if err != nil {
		return input
	}
	// Convert the time to the local timezone
	localBuildTime := buildTime.In(localLocation)
	return fmt.Sprintf("%s (%s)", localBuildTime.Format("2006-01-02 3:04 PM MST"), timeago(localBuildTime))
}

// Seconds-based time units
const (
	Day      = 24 * time.Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 37 * Year
)

// Time formats a time into a relative string.
//
// Time(someT) -> "3 weeks ago"
//
// stolen from -> https://github.com/dustin/go-humanize/blob/master/times.go
func timeago(then time.Time) string {
	return CustomRelTime(then, time.Now(), "ago", "from now", defaultMagnitudes)
}

type RelTimeMagnitude struct {
	D      time.Duration
	Format string
	DivBy  time.Duration
}

var defaultMagnitudes = []RelTimeMagnitude{
	{time.Second, "now", time.Second},
	{2 * time.Second, "1 second %s", 1},
	{time.Minute, "%d seconds %s", time.Second},
	{2 * time.Minute, "1 minute %s", 1},
	{time.Hour, "%d minutes %s", time.Minute},
	{2 * time.Hour, "1 hour %s", 1},
	{Day, "%d hours %s", time.Hour},
	{2 * Day, "1 day %s", 1},
	{Week, "%d days %s", Day},
	{2 * Week, "1 week %s", 1},
	{Month, "%d weeks %s", Week},
	{2 * Month, "1 month %s", 1},
	{Year, "%d months %s", Month},
	{18 * Month, "1 year %s", 1},
	{2 * Year, "2 years %s", 1},
	{LongTime, "%d years %s", Year},
	{math.MaxInt64, "a long while %s", 1},
}

func CustomRelTime(a, b time.Time, albl, blbl string, magnitudes []RelTimeMagnitude) string {
	lbl := albl
	diff := b.Sub(a)

	if a.After(b) {
		lbl = blbl
		diff = a.Sub(b)
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].D > diff
	})

	if n >= len(magnitudes) {
		n = len(magnitudes) - 1
	}
	mag := magnitudes[n]
	var args []interface{}
	escaped := false
	for _, ch := range mag.Format {
		if escaped {
			switch ch {
			case 's':
				args = append(args, lbl)
			case 'd':
				args = append(args, diff/mag.DivBy)
			}
			escaped = false
		} else {
			escaped = ch == '%'
		}
	}
	return fmt.Sprintf(mag.Format, args...)
}
