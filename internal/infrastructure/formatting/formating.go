package formatting

import "fmt"

const (
	RED     = "\033[0;31m"
	GREEN   = "\033[0;32m"
	YELLOW  = "\033[0;33m"
	BLUE    = "\033[0;34m"
	MAGENTA = "\033[0;35m"
	CYAN    = "\033[0;36m"
	WHITE   = "\033[0;37m"
	NC      = "\033[0m"

	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	BLINK     = "\033[5m"
	REVERSE   = "\033[7m"
	CONCEALED = "\033[8m"

	BLACKBG   = "\033[40m"
	REDBG     = "\033[41m"
	GREENBG   = "\033[42m"
	YELLOWBG  = "\033[43m"
	BLUEBG    = "\033[44m"
	MAGENTABG = "\033[45m"
	CYANBG    = "\033[46m"
	WHITEBG   = "\033[47m"
	NCBG      = "\033[49m"

	BLACKBG_WHITE = "\033[40;37m"

	BLACKBG_YELLOW   = "\033[40;33m"
	REDBG_YELLOW     = "\033[41;33m"
	GREENBG_YELLOW   = "\033[42;33m"
	YELLOWBG_YELLOW  = "\033[43;33m"
	BLUEBG_YELLOW    = "\033[44;33m"
	MAGENTABG_YELLOW = "\033[45;33m"

	REDBG_WHITE = "\033[41;37m"
	REDBG_BLACK = "\033[41;30m"

	BLACK_UNDERLINE  = "\033[4;30m"
	RED_UNDERLINE    = "\033[4;31m"
	GREEN_UNDERLINE  = "\033[4;32m"
	YELLOW_UNDERLINE = "\033[4;33m"
	BLUE_UNDERLINE   = "\033[4;34m"

	BLACK_BOLD  = "\033[1;30m"
	RED_BOLD    = "\033[1;31m"
	GREEN_BOLD  = "\033[1;32m"
	YELLOW_BOLD = "\033[1;33m"

	BLACK_BOLD_BRIGHT  = "\033[1;90m"
	RED_BOLD_BRIGHT    = "\033[1;91m"
	GREEN_BOLD_BRIGHT  = "\033[1;92m"
	YELLOW_BOLD_BRIGHT = "\033[1;93m"
)

func LogError() string {
	return fmt.Sprintf("%sERROR%s", RED, NC)
}
