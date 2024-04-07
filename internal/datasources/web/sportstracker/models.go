package sportstracker

import (
	"errors"
	"fmt"
	"time"

	"github.com/microhod/sweaty-swapper/internal/domain"
)

const activityIDSource = "sportstracker"

var ErrActivityNotSupported = errors.New("activity type not supported")

type Workout struct {
	Key         string   `json:"workoutKey"`
	Activity    Activity `json:"activityId"`
	Created     int64    `json:"created"`
	Description string   `json:"description"`
	Photos      []Photo  `json:"photos"`
	Videos      []Video  `json:"videos"`
	GPX         GPX      `json:"gpx"`
}

func (w Workout) ToActivity() (domain.Activity, error) {
	activityType, err := w.Activity.ToActivityType()
	if err != nil {
		return domain.Activity{}, fmt.Errorf("converting activity to domain activity type: %w", err)
	}
	createdAt := time.UnixMilli(w.Created)

	activity := domain.Activity{
		ID: domain.ActivityID{
			Source: activityIDSource,
			ID:     w.Key,
		},
		Type: activityType,
		Description: w.Description,
		CreatedAt: createdAt,
		Route: domain.Route{
			Type: domain.RouteTypeGpx,
			Data: w.GPX,
		},
	}
	activity.Title = activity.DefaultTitle()

	return activity, nil
}

type Photo struct {
	Key    string `json:"key"`
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Video struct {
	Key          string `json:"key"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
}

type GPX []byte

type Activity int

const (
	ActivityWalking            Activity = 0
	ActivityRunning            Activity = 1
	ActivityCycling            Activity = 2
	ActivityCrossCountrySkiing Activity = 3
	ActivityOther1             Activity = 4
	ActivityOther2             Activity = 5
	ActivityOther3             Activity = 6
	ActivityOther4             Activity = 7
	ActivityOther5             Activity = 8
	ActivityOther6             Activity = 9
	ActivityMountainBiking     Activity = 10
	ActivityHiking             Activity = 11
	ActivityRollerSkating      Activity = 12
	ActivityAlpineSkiing       Activity = 13
	ActivityPaddling           Activity = 14
	ActivityRowing             Activity = 15
	ActivityGolf               Activity = 16
	ActivityIndoor             Activity = 17
	ActivityParkour            Activity = 18
	ActivityBallGames          Activity = 19
	ActivityOutdoorGym         Activity = 20
	ActivityPoolSwimming       Activity = 21
	ActivityTrailRunning       Activity = 22
	ActivityGym                Activity = 23
	ActivityNordicWalking      Activity = 24
	ActivityHorsebackRiding    Activity = 25
	ActivityMotorsports        Activity = 26
	ActivitySkateboarding      Activity = 27
	ActivityWaterSports        Activity = 28
	ActivityClimbing           Activity = 29
	ActivitySnowboarding       Activity = 30
	ActivitySkiTouring         Activity = 31
	ActivityFitnessClass       Activity = 32
	ActivitySoccer             Activity = 33
	ActivityTennis             Activity = 34
	ActivityBasketball         Activity = 35
	ActivityBadminton          Activity = 36
	ActivityBaseball           Activity = 37
	ActivityVolleyball         Activity = 38
	ActivityAmericanFootball   Activity = 39
	ActivityTableTennis        Activity = 40
	ActivityRacquetBall        Activity = 41
	ActivitySquash             Activity = 42
	ActivityFloorball          Activity = 43
	ActivityHandball           Activity = 44
	ActivitySoftball           Activity = 45
	ActivityBowling            Activity = 46
	ActivityCricket            Activity = 47
	ActivityRugby              Activity = 48
	ActivityIceSkating         Activity = 49
	ActivityIceHockey          Activity = 50
	ActivityYoga               Activity = 51
	ActivityIndoorCycling      Activity = 52
	ActivityTreadmill          Activity = 53
	ActivityCrossfit           Activity = 54
	ActivityCrosstrainer       Activity = 55
	ActivityRollerSkiing       Activity = 56
	ActivityIndoorRowing       Activity = 57
	ActivityStrecthing         Activity = 58
	ActivityTrackAndField      Activity = 59
	ActivityOrienteering       Activity = 60
	ActivityStandupPaddling    Activity = 61
	ActivityMartialArts        Activity = 62
	ActivityKettlebell         Activity = 63
	ActivityDancing            Activity = 64
	ActivitySnowShoeing        Activity = 65
	ActivityFrisbee            Activity = 66
	ActivityFutsal             Activity = 67
	ActivityMultisport         Activity = 68
	ActivityAerobics           Activity = 69
	ActivityTrekking           Activity = 70
	ActivitySailing            Activity = 71
	ActivityKayaking           Activity = 72
	ActivityCircuitTraining    Activity = 73
	ActivityTriathlon          Activity = 74
	ActivityPadel              Activity = 75
	ActivityCheerleading       Activity = 76
	ActivityBoxing             Activity = 77
	ActivityScubaDiving        Activity = 78
	ActivityFreeDiving         Activity = 79
	ActivityAdventureRacing    Activity = 80
	ActivityGymnastics         Activity = 81
	ActivityCanoeing           Activity = 82
	ActivityMountaineering     Activity = 83
	ActivityTelemarkSkiing     Activity = 84
	ActivityOpenwaterSwimming  Activity = 85
	ActivityWindsurfing        Activity = 86
	ActivityKitesurfing        Activity = 87
	ActivityParagliding        Activity = 88
	ActivitySnorkeling         Activity = 90
	ActivitySurfing            Activity = 91
	ActivitySwimrun            Activity = 92
	ActivityDuathlon           Activity = 93
	ActivityAquathlon          Activity = 94
	ActivityObstacleRacing     Activity = 95
	ActivityFishing            Activity = 96
	ActivityHunting            Activity = 97
	ActivityTransition         Activity = 98
	ActivityGravelCycling      Activity = 99
)

var activityToActivityType = map[Activity]domain.ActivityType{
	ActivityRunning: domain.ActivityTypeRunning,
	ActivityWalking: domain.ActivityTypeWalking,
	ActivityHiking:  domain.ActivityTypeHiking,
	ActivityCycling: domain.ActivityTypeCycling,
	ActivityGym:     domain.ActivityTypeGym,
	ActivityYoga:    domain.ActivityTypeYoga,
}

func (a Activity) ToActivityType() (domain.ActivityType, error) {
	activityType, exists := activityToActivityType[a]
	if !exists {
		return "", fmt.Errorf("%w: [%d]", ErrActivityNotSupported, a)
	}

	return activityType, nil
}
