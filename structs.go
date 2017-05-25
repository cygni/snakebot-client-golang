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
	GameTick          int    `json:"gameTick"`
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
	MaxNoofPlayers                     int  `json:"maxNoofPlayers"`
	StartSnakeLength                   int  `json:"startSnakeLength"`
	TimeInMsPerTick                    int  `json:"timeInMsPerTick"`
	ObstaclesEnabled                   bool `json:"obstaclesEnabled"`
	FoodEnabled                        bool `json:"foodEnabled"`
	HeadToTailConsumes                 bool `json:"headToTailConsumes"`
	TailConsumeGrows                   bool `json:"tailConsumeGrows"`
	AddFoodLikelihood                  int  `json:"addFoodLikelihood"`
	RemoveFoodLikelihood               int  `json:"removeFoodLikelihood"`
	SpontaneousGrowthEveryNWorldTick   int  `json:"spontaneousGrowthEveryNWorldTick"`
	TrainingGame                       bool `json:"trainingGame"`
	PointsPerLength                    int  `json:"pointsPerLength"`
	PointsPerFood                      int  `json:"pointsPerFood"`
	PointsPerCausedDeath               int  `json:"pointsPerCausedDeath"`
	PointsPerNibble                    int  `json:"pointsPerNibble"`
	NoofRoundsTailProtectedAfterNibble int  `json:"noofRoundsTailProtectedAfterNibble"`
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
	GameTick          int    `json:"gameTick"`
	Map               Map    `json:"map"`
}

type InvalidPlayerName struct {
	ReasonCode int `json:"reasonCode"`
}

type GameEnded struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
	PlayerWinnerId    string `json:"playerWinnerId"`
	GameId            string `json:"gameId"`
	GameTick          int    `json:"gameTick"`
	Map               Map    `json:"map"`
}

type SnakeDead struct {
	PlayerId    string `json:"playerId"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	GameId      string `json:"gameId"`
	GameTick    int    `json:"gameTick"`
	DeathReason string `json:"deathReason"`
}

type GameStarting struct {
	ReceivingPlayerId string `json:"receivingPlayerId"`
	GameId            string `json:"gameId"`
	NoofPlayers       int    `json:"noofPlayers"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
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
	Rank       int    `json:"rank"`
	Points     int    `json:"points"`
	Alive      bool   `json:"alive"`
}

type Map struct {
	Width             int         `json:"width"`
	Height            int         `json:"height"`
	WorldTick         int         `json:"worldTick"`
	SnakeInfos        []SnakeInfo `json:"snakeInfos"`
	FoodPositions     []int       `json:"foodPositions"`
	ObstaclePositions []int       `json:"obstaclePositions"`
}

type SnakeInfo struct {
	Name                      string `json:"name"`
	Points                    int    `json:"points"`
	Positions                 []int  `json:"positions"`
	TailProtectedForGameTicks int    `json:"tailProtectedForGameTicks"`
	Id                        string `json:"id"`
}
