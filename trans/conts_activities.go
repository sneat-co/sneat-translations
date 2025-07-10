package trans

import "strings"

func GetActivityCode(activity string) string {
	return activityCodesByTranslation[strings.ToLower(activity)]
}

// Sports activities
const (
	ActivityHiking        = "hiking"
	ActivityBasketball    = "basketball"
	ActivitySoccer        = "soccer"
	ActivityPingPong      = "ping-pong"
	ActivityTennis        = "tennis"
	ActivitySurfing       = "surfing"
	ActivityKitesurfing   = "kitesurfing"
	ActivityRunning       = "running"
	ActivityCycling       = "cycling"
	ActivitySkateboarding = "skateboarding"
	ActivityVolleyball    = "volleyball"
	ActivitySwimming      = "swimming"
	ActivityYoga          = "yoga"
	ActivityClimbing      = "climbing"
	ActivityGym           = "gym"
)

const (
	ActivityFishing = "fishing"
)

// Social activities
const (
	ActivityBookClub     = "book-club"
	ActivityCoffeeMeetup = "coffee-meetup"
	ActivityGameNight    = "game-night"
	ActivityMovieNight   = "movie-night"
	ActivityTriviaNight  = "trivia-night"
	ActivityPotluck      = "potluck"
	ActivityPicnic       = "picnic"
	ActivityBarbecue     = "barbecue"
	ActivityCrafting     = "crafting"
	ActivityKaraoke      = "karaoke"
	ActivityMusicJam     = "music-jam"
	ActivityBoardGames   = "board-games"
	ActivityArtJam       = "art-jam"
)

// Family-oriented activities
const (
	ActivityParkPlaydate    = "park-playdate"
	ActivityStoryTime       = "story-time"
	ActivityToySwap         = "toy-swap"
	ActivityHomeworkHelp    = "homework-help"
	ActivityBirthdayParty   = "birthday-party"
	ActivityBabysittingSwap = "babysitting-swap"
)

// Teen-focused activities
const (
	ActivityStudyGroup = "study-group"
	ActivityGaming     = "gaming"
	ActivitySkating    = "skating"
	ActivityDanceClass = "dance-class"
	ActivityHangout    = "hangout"
	ActivityEscapeRoom = "escape-room"
)

var AllActivityCodes = []string{
	ActivityHiking,
	ActivityBasketball,
	ActivitySoccer,
	ActivityPingPong,
	ActivityTennis,
	ActivitySurfing,
	ActivityKitesurfing,
	ActivityRunning,
	ActivityCycling,
	ActivitySkateboarding,
	ActivityVolleyball,
	ActivitySwimming,
	ActivityYoga,
	ActivityClimbing,
	ActivityGym,
	ActivityFishing,
	ActivityBookClub,
	ActivityCoffeeMeetup,
	ActivityGameNight,
	ActivityMovieNight,
	ActivityTriviaNight,
	ActivityPotluck,
	ActivityPicnic,
	ActivityBarbecue,
	ActivityCrafting,
	ActivityKaraoke,
	ActivityMusicJam,
	ActivityBoardGames,
	ActivityArtJam,
	ActivityParkPlaydate,
	ActivityStoryTime,
	ActivityToySwap,
	ActivityHomeworkHelp,
	ActivityBirthdayParty,
	ActivityBabysittingSwap,
	ActivityStudyGroup,
	ActivityGaming,
	ActivitySkating,
	ActivityDanceClass,
	ActivityHangout,
	ActivityEscapeRoom,
}

var activityCodesByTranslation map[string]string

func init() {
	activityCodesByTranslation = make(map[string]string, len(AllActivityCodes)*len(SupportedLocales))
	for _, code := range AllActivityCodes {
		if translations := TRANS[code]; translations != nil {
			for _, translation := range translations {
				translation = strings.ToLower(translation)
				if _, ok := activityCodesByTranslation[translation]; !ok {
					activityCodesByTranslation[translation] = code
				}
			}
		}
		activityCodesByTranslation[code] = code
	}
}
