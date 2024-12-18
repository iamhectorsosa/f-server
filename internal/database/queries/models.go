// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

type Match struct {
	ID    int `json:"id"`
	TeamA int `json:"team_a"`
	TeamB int `json:"team_b"`
}

type Player struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Team struct {
	ID      int    `json:"id"`
	Player1 string `json:"player_1"`
	Player2 string `json:"player_2"`
	Score   int    `json:"score"`
}
