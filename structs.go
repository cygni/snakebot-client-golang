package main

type TypedMessage struct {
	Type string `json:"type"`
}

// Outbound
type RegisterPlayer struct {
	Type       string `json:"type"`
	PlayerName string `json:"playerName"`
}

type StartGame struct {
	Type string `json:"type"`
}

type RegisterMove struct {
	Type              string `json:"type"`
	Direction         string `json:"direction"`
	GameTick          int32  `json:"gameTick"`
	ReceivingPlayerId string `json:"receivingPlayerId"`
	GameId            string `json:"gameId"`
}

type HeartBeat struct {
	Type              string `json:"type"`
	ReceivingPlayerId string `json:"receivingPlayerId"`
}

type ClientInfo struct {
	Type                   string `json:"type"`
	Language               string `json:"language"`
	LanguageVersion        string `json:"languageVersion"`
	OperatingSystem        string `json:"operatingSystem"`
	OperatingSystemVersion string `json:"operatingSystemVersion"`
	ClientVersion          string `json:"clientVersion"`
}

type GameSettings struct {
	MaxNoofPlayers                     int32 `json:"maxNoofPlayers"`
	StartSnakeLength                   int32 `json:"startSnakeLength"`
	TimeInMsPerTick                    int32 `json:"timeInMsPerTick"`
	ObstaclesEnabled                   bool  `json:"obstaclesEnabled"`
	FoodEnabled                        bool  `json:"foodEnabled"`
	HeadToTailConsumes                 bool  `json:"headToTailConsumes"`
	TailConsumeGrows                   bool  `json:"tailConsumeGrows"`
	AddFoodLikelihood                  int32 `json:"addFoodLikelihood"`
	RemoveFoodLikelihood               int32 `json:"removeFoodLikelihood"`
	SpontaneousGrowthEveryNWorldTick   int32 `json:"spontaneousGrowthEveryNWorldTick"`
	TrainingGame                       bool  `json:"trainingGame"`
	PointsPerLength                    int32 `json:"pointsPerLength"`
	PointsPerFood                      int32 `json:"pointsPerFood"`
	PointsPerCausedDeath               int32 `json:"pointsPerCausedDeath"`
	PointsPerNibble                    int32 `json:"pointsPerNibble"`
	NoofRoundsTailProtectedAfterNibble int32 `json:"noofRoundsTailProtectedAfterNibble"`
}

// Inbound
type PlayerRegistered struct {
	GameId            string       `json:"gameId"`
	GameMode          string       `json:"gameMode"`
	ReceivingPlayerId string       `json:"receivingPlayerId"`
	Name              string       `json:"name"`
	GameSettings      GameSettings `json:"gameSettings"`
}

type MapUpdate struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
	GameId            string `json:"gameId"`
	GameTick          int32  `json:"gameTick"`
	Map               Map    `json:"map"`
}

type InvalidPlayerName struct {
	ReasonCode int32 `json:"reasonCode"`
}

type GameEnded struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
	PlayerWinnerId    string `json:"playerWinnerId"`
	GameId            string `json:"gameId"`
	GameTick          int32  `json:"gameTick"`
	Map               Map    `json:"map"`
}

type SnakeDead struct {
	PlayerId    string `json:"playerId"`
	X           int32  `json:"x"`
	Y           int32  `json:"y"`
	GameId      string `json:"gameId"`
	GameTick    int32  `json:"gameTick"`
	DeathReason string `json:"deathReason"`
}

type GameStarting struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
	GameId            string `json:"gameId"`
	NoofPlayers       int32  `json:"noofPlayers"`
	Width             int32  `json:"width"`
	Height            int32  `json:"height"`
}

type HeartBeatResponse struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
}

type GameLink struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
	GameId            string `json:"gameId"`
	Url               string `json:"url"`
}

type TournamentEnded struct {
	ReceivingPlayerId string       `json:"receivingPlayerId"`
	TournamentId      string       `json:"tournamentId"`
	TournamentName    string       `json:"tournamentName"`
	GameResult        []GameResult `json:"gameResult"`
	GameId            string       `json:"gameId"`
	PlayerWinnerId    string       `json:"playerWinnerId"`
}

type GameResult struct {
	GameId            string       `json:"gameId"`
	ReceivingPlayerId string       `json:"receivingPlayerId"`
	PlayerRanks       []PlayerRank `json:"playerRanks"`
}

type PlayerRank struct {
	PlayerName string `json:"playerName"`
	PlayerId   string `json:"playerId"`
	Rank       int32  `json:"rank"`
	Points     int32  `json:"points"`
	Alive      bool   `json:"alive"`
}

type Map struct {
	Width             int32       `json:"width"`
	Height            int32       `json:"height"`
	WorldTick         int32       `json:"worldTick"`
	SnakeInfos        []SnakeInfo `json:"snakeInfos"`
	FoodPositions     []int32     `json:"foodPositions"`
	ObstaclePositions []int32     `json:"obstaclePositions"`
}

type SnakeInfo struct {
	Name                      string  `json:"name"`
	Points                    int32   `json:"points"`
	Positions                 []int32 `json:"positions"`
	TailProtectedForGameTicks int32   `json:"tailProtectedForGameTicks"`
	Id                        string  `json:"id"`
}
