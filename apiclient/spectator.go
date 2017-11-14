package apiclient

import (
	"context"
	"fmt"

	"github.com/yuhanfang/riot/constants/region"
)

type CurrentGameInfoDTO struct {
	GameID            int64                       // The ID of the game
	GameStartTime     int64                       // The game start time represented in epoch milliseconds
	PlatformID        string                      // The ID of the platform on which the game is being played
	GameMode          string                      // The game mode
	MapID             int64                       // The ID of the map
	GameType          string                      // The game type
	BannedChampions   []BannedChampionDTO         // Banned champion information
	Observers         ObserverDTO                 // The observer information
	Participants      []CurrentGameParticipantDTO // The participant information
	GameLength        int64                       // The amount of time in seconds that has passed since the game started
	GameQueueConfigID int64                       // The queue type (queue types are documented on the Game Constants page)
}

type BannedChampionDTO struct {
	PickTurn   int   // The turn during which the champion was banned
	ChampionID int64 // The ID of the banned champion
	TeamID     int64 // The ID of the team that banned the champion
}

type ObserverDTO struct {
	EncryptionKey string // Key used to decrypt the spectator grid game data for playback
}

type CurrentGameParticipantDTO struct {
	ProfileIconId int64                              // The ID of the profile icon used by this participant
	ChampionId    int64                              // The ID of the champion played by this participant
	SummonerName  string                             // The summoner name of this participant
	Runes         []CurrentGameParticipantRuneDTO    // The runes used by this participant
	Bot           bool                               // Flag indicating whether or not this participant is a bot
	TeamId        int64                              // The team ID of this participant, indicating the participant's team
	Spell2Id      int64                              // The ID of the second summoner spell used by this participant
	Masteries     []CurrentGameParticipantMasteryDTO // The masteries used by this participant
	Spell1Id      int64                              // The ID of the first summoner spell used by this participant
	SummonerId    int64                              // The summoner ID of this participant
}

type CurrentGameParticipantRuneDTO struct {
	Count  int   // The count of this rune used by the participant
	RuneId int64 // The ID of the rune
}

type CurrentGameParticipantMasteryDTO struct {
	MasteryId int64 // The ID of the mastery
	Rank      int   // The number of points put into this mastery by the user
}

func (c *client) GetCurrentGameInfoBySummoner(ctx context.Context, r region.Region, summonerID int64) (*CurrentGameInfoDTO, error) {
	var res CurrentGameInfoDTO
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/spectator/v3/active-games/by-summoner", fmt.Sprintf("/%d", summonerID), nil, &res)
	return &res, err
}

type FeaturedGamesDTO struct {
	ClientRefreshInterval int64                 // The suggested interval to wait before requesting FeaturedGames again
	GameList              []FeaturedGameInfoDTO // 	The list of featured games
}

type FeaturedGameInfoDTO struct {
	GameId            int64                        // The ID of the game
	GameStartTime     int64                        // The game start time represented in epoch milliseconds
	PlatformId        string                       // The ID of the platform on which the game is being played
	GameMode          string                       // The game mode
	MapId             int64                        // The ID of the map
	GameType          string                       // The game type
	BannedChampions   []BannedChampionDTO          // 	Banned champion information
	Observers         ObserverDTO                  // The observer information
	Participants      []FeaturedGameParticipantDTO //The participant information
	GameLength        int64                        // The amount of time in seconds that has passed since the game started
	GameQueueConfigId int64                        // The queue type (queue types are documented on the Game Constants page)
}

type FeaturedGameParticipantDTO struct {
	ProfileIconId int64  // The ID of the profile icon used by this participant
	ChampionId    int64  // The ID of the champion played by this participant
	SummonerName  string // The summoner name of this participant
	Bot           bool   // Flag indicating whether or not this participant is a bot
	Spell2Id      int64  // The ID of the second summoner spell used by this participant
	TeamId        int64  // The team ID of this participant, indicating the participant's team
	Spell1Id      int64  // The ID of the first summoner spell used by this participant
}

func (c *client) GetFeaturedGames(ctx context.Context, r region.Region) (*FeaturedGamesDTO, error) {
	var res FeaturedGamesDTO
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/spectator/v3/featured-games", "", nil, &res)
	return &res, err
}