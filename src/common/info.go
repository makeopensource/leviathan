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
var CommitInfo = "dev"
var BuildDate = "dev"
var Branch = "dev" // Git branch
var GoVersion = runtime.Version()

func PrintInfo() {
	// generated from https://patorjk.com/software/taag/#p=testall&t=leviathan
	var headers = []string{
		"\n ___        _______    ___      ___  ___   _________   ___  ___   ________   ________      \n|\\  \\      |\\  ___ \\  |\\  \\    /  /||\\  \\ |\\___   ___\\|\\  \\|\\  \\ |\\   __  \\ |\\   ___  \\    \n\\ \\  \\     \\ \\   __/| \\ \\  \\  /  / /\\ \\  \\\\|___ \\  \\_|\\ \\  \\\\\\  \\\\ \\  \\|\\  \\\\ \\  \\\\ \\  \\   \n \\ \\  \\     \\ \\  \\_|/__\\ \\  \\/  / /  \\ \\  \\    \\ \\  \\  \\ \\   __  \\\\ \\   __  \\\\ \\  \\\\ \\  \\  \n  \\ \\  \\____ \\ \\  \\_|\\ \\\\ \\    / /    \\ \\  \\    \\ \\  \\  \\ \\  \\ \\  \\\\ \\  \\ \\  \\\\ \\  \\\\ \\  \\ \n   \\ \\_______\\\\ \\_______\\\\ \\__/ /      \\ \\__\\    \\ \\__\\  \\ \\__\\ \\__\\\\ \\__\\ \\__\\\\ \\__\\\\ \\__\\\n    \\|_______| \\|_______| \\|__|/        \\|__|     \\|__|   \\|__|\\|__| \\|__|\\|__| \\|__| \\|__|\n                                                                                           \n                                                                                           \n                                                                                           \n",
		"\n      ___       ___           ___                       ___           ___           ___           ___           ___     \n     /\\__\\     /\\  \\         /\\__\\          ___        /\\  \\         /\\  \\         /\\__\\         /\\  \\         /\\__\\    \n    /:/  /    /::\\  \\       /:/  /         /\\  \\      /::\\  \\        \\:\\  \\       /:/  /        /::\\  \\       /::|  |   \n   /:/  /    /:/\\:\\  \\     /:/  /          \\:\\  \\    /:/\\:\\  \\        \\:\\  \\     /:/__/        /:/\\:\\  \\     /:|:|  |   \n  /:/  /    /::\\~\\:\\  \\   /:/__/  ___      /::\\__\\  /::\\~\\:\\  \\       /::\\  \\   /::\\  \\ ___   /::\\~\\:\\  \\   /:/|:|  |__ \n /:/__/    /:/\\:\\ \\:\\__\\  |:|  | /\\__\\  __/:/\\/__/ /:/\\:\\ \\:\\__\\     /:/\\:\\__\\ /:/\\:\\  /\\__\\ /:/\\:\\ \\:\\__\\ /:/ |:| /\\__\\\n \\:\\  \\    \\:\\~\\:\\ \\/__/  |:|  |/:/  / /\\/:/  /    \\/__\\:\\/:/  /    /:/  \\/__/ \\/__\\:\\/:/  / \\/__\\:\\/:/  / \\/__|:|/:/  /\n  \\:\\  \\    \\:\\ \\:\\__\\    |:|__/:/  /  \\::/__/          \\::/  /    /:/  /           \\::/  /       \\::/  /      |:/:/  / \n   \\:\\  \\    \\:\\ \\/__/     \\::::/__/    \\:\\__\\          /:/  /     \\/__/            /:/  /        /:/  /       |::/  /  \n    \\:\\__\\    \\:\\__\\        ~~~~         \\/__/         /:/  /                      /:/  /        /:/  /        /:/  /   \n     \\/__/     \\/__/                                   \\/__/                       \\/__/         \\/__/         \\/__/    \n",
		"\n          (`-')  _      (`-')  _     (`-')  _ (`-')      (`-').-> (`-')  _ <-. (`-')_ \n   <-.    ( OO).-/     _(OO ) (_)    (OO ).-/ ( OO).->   (OO )__  (OO ).-/    \\( OO) )\n ,--. )  (,------.,--.(_/,-.\\ ,-(`-')/ ,---.  /    '._  ,--. ,'-' / ,---.  ,--./ ,--/ \n |  (`-') |  .---'\\   \\ / (_/ | ( OO)| \\ /`.\\ |'--...__)|  | |  | | \\ /`.\\ |   \\ |  | \n |  |OO )(|  '--.  \\   /   /  |  |  )'-'|_.' |`--.  .--'|  `-'  | '-'|_.' ||  . '|  |)\n(|  '__ | |  .--' _ \\     /_)(|  |_/(|  .-.  |   |  |   |  .-.  |(|  .-.  ||  |\\    | \n |     |' |  `---.\\-'\\   /    |  |'->|  | |  |   |  |   |  | |  | |  | |  ||  | \\   | \n `-----'  `------'    `-'     `--'   `--' `--'   `--'   `--' `--' `--' `--'`--'  `--' \n",
		"\n ___       _______  ___      ___  __          __  ___________  __    __       __      _____  ___   \n|\"  |     /\"     \"||\"  \\    /\"  ||\" \\        /\"\"\\(\"     _   \")/\" |  | \"\\     /\"\"\\    (\\\"   \\|\"  \\  \n||  |    (: ______) \\   \\  //  / ||  |      /    \\)__/  \\\\__/(:  (__)  :)   /    \\   |.\\\\   \\    | \n|:  |     \\/    |    \\\\  \\/. ./  |:  |     /' /\\  \\  \\\\_ /    \\/      \\/   /' /\\  \\  |: \\.   \\\\  | \n \\  |___  // ___)_    \\.    //   |.  |    //  __'  \\ |.  |    //  __  \\\\  //  __'  \\ |.  \\    \\. | \n( \\_|:  \\(:      \"|    \\\\   /    /\\  |\\  /   /  \\\\  \\\\:  |   (:  (  )  :)/   /  \\\\  \\|    \\    \\ | \n \\_______)\\_______)     \\__/    (__\\_|_)(___/    \\___)\\__|    \\__|  |__/(___/    \\___)\\___|\\____\\) \n                                                                                                   \n",
		"\n██╗     ███████╗██╗   ██╗██╗ █████╗ ████████╗██╗  ██╗ █████╗ ███╗   ██╗\n██║     ██╔════╝██║   ██║██║██╔══██╗╚══██╔══╝██║  ██║██╔══██╗████╗  ██║\n██║     █████╗  ██║   ██║██║███████║   ██║   ███████║███████║██╔██╗ ██║\n██║     ██╔══╝  ╚██╗ ██╔╝██║██╔══██║   ██║   ██╔══██║██╔══██║██║╚██╗██║\n███████╗███████╗ ╚████╔╝ ██║██║  ██║   ██║   ██║  ██║██║  ██║██║ ╚████║\n╚══════╝╚══════╝  ╚═══╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝\n                                                                       \n",
		"\n         _             _     _          _        _          _                 _            _       _    _                   _          \n        _\\ \\          /\\ \\  /\\ \\    _ / /\\      /\\ \\       / /\\              /\\ \\         / /\\    / /\\ / /\\                /\\ \\     _  \n       /\\__ \\        /  \\ \\ \\ \\ \\  /_/ / /      \\ \\ \\     / /  \\             \\_\\ \\       / / /   / / // /  \\              /  \\ \\   /\\_\\\n      / /_ \\_\\      / /\\ \\ \\ \\ \\ \\ \\___\\/       /\\ \\_\\   / / /\\ \\            /\\__ \\     / /_/   / / // / /\\ \\            / /\\ \\ \\_/ / /\n     / / /\\/_/     / / /\\ \\_\\/ / /  \\ \\ \\      / /\\/_/  / / /\\ \\ \\          / /_ \\ \\   / /\\ \\__/ / // / /\\ \\ \\          / / /\\ \\___/ / \n    / / /         / /_/_ \\/_/\\ \\ \\   \\_\\ \\    / / /    / / /  \\ \\ \\        / / /\\ \\ \\ / /\\ \\___\\/ // / /  \\ \\ \\        / / /  \\/____/  \n   / / /         / /____/\\    \\ \\ \\  / / /   / / /    / / /___/ /\\ \\      / / /  \\/_// / /\\/___/ // / /___/ /\\ \\      / / /    / / /   \n  / / / ____    / /\\____\\/     \\ \\ \\/ / /   / / /    / / /_____/ /\\ \\    / / /      / / /   / / // / /_____/ /\\ \\    / / /    / / /    \n / /_/_/ ___/\\ / / /______      \\ \\ \\/ /___/ / /__  / /_________/\\ \\ \\  / / /      / / /   / / // /_________/\\ \\ \\  / / /    / / /     \n/_______/\\__\\// / /_______\\      \\ \\  //\\__\\/_/___\\/ / /_       __\\ \\_\\/_/ /      / / /   / / // / /_       __\\ \\_\\/ / /    / / /      \n\\_______\\/    \\/__________/       \\_\\/ \\/_________/\\_\\___\\     /____/_/\\_\\/       \\/_/    \\/_/ \\_\\___\\     /____/_/\\/_/     \\/_/       \n                                                                                                                                       \n",
		"\n _        _______          _________ _______ _________          _______  _       \n( \\      (  ____ \\|\\     /|\\__   __/(  ___  )\\__   __/|\\     /|(  ___  )( (    /|\n| (      | (    \\/| )   ( |   ) (   | (   ) |   ) (   | )   ( || (   ) ||  \\  ( |\n| |      | (__    | |   | |   | |   | (___) |   | |   | (___) || (___) ||   \\ | |\n| |      |  __)   ( (   ) )   | |   |  ___  |   | |   |  ___  ||  ___  || (\\ \\) |\n| |      | (       \\ \\_/ /    | |   | (   ) |   | |   | (   ) || (   ) || | \\   |\n| (____/\\| (____/\\  \\   /  ___) (___| )   ( |   | |   | )   ( || )   ( || )  \\  |\n(_______/(_______/   \\_/   \\_______/|/     \\|   )_(   |/     \\||/     \\||/    )_)\n                                                                                 \n",
		"\n                                                 \n (                          )    )               \n )\\   (    )   (      )  ( /( ( /(     )         \n((_) ))\\  /((  )\\  ( /(  )\\()))\\()) ( /(   (     \n _  /((_)(_))\\((_) )(_))(_))/((_)\\  )(_))  )\\ )  \n| |(_))  _)((_)(_)((_)_ | |_ | |(_)((_)_  _(_/(  \n| |/ -_) \\ V / | |/ _` ||  _|| ' \\ / _` || ' \\)) \n|_|\\___|  \\_/  |_|\\__,_| \\__||_||_|\\__,_||_||_|  \n                                                 \n",
	}
	const (
		colorReset  = "\033[0m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
	)

	// Get terminal width
	width := 60

	// Print header
	fmt.Println(strings.Repeat("=", width))
	fmt.Printf("%s%s %s %s\n", colorYellow, strings.Repeat(" ", (width-24)/2), headers[rand.IntN(len(headers))], colorReset)
	fmt.Println(strings.Repeat("=", width))

	// Print app info with aligned values
	printField := func(name, value string) {
		fmt.Printf("%-15s: %s%s%s\n", name, colorGreen, value, colorReset)
	}

	printField("Version", Version)
	printField("CommitInfo", CommitInfo)
	printField("BuildDate", formatTime(BuildDate))
	printField("Branch", Branch)
	printField("GoVersion", runtime.Version())

	// Add GitHub URL if repo info is available
	if Branch != "dev" && CommitInfo != "dev" {
		fmt.Println(strings.Repeat("-", width))
		fmt.Println("Links:")
		githubURL := GetGitHubURL(Branch, CommitInfo)
		fmt.Println(githubURL)
	}

	fmt.Println(strings.Repeat("=", width))
}

func GetGitHubURL(branch, commitHash string) string {
	const (
		repoOwner = "makeopensource"
		repoName  = "leviathan"
	)
	// For browsing at the specific branch
	branchURL := fmt.Sprintf("https://github.com/%s/%s/tree/%s",
		repoOwner, repoName, branch)
	// For viewing the specific commit
	commitURL := fmt.Sprintf("https://github.com/%s/%s/commit/%s",
		repoOwner, repoName, commitHash)

	return fmt.Sprintf("Branch: %s\nCommit: %s", branchURL, commitURL)
}

func formatTime(input string) string {
	buildTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		//fmt.Printf("Error parsing build time: %v\n", err)
		return input
	}
	return humanizeTime(buildTime)
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
func humanizeTime(then time.Time) string {
	return RelTime(then, time.Now(), "ago", "from now")
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

func RelTime(a, b time.Time, albl, blbl string) string {
	return CustomRelTime(a, b, albl, blbl, defaultMagnitudes)
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
	args := []interface{}{}
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
