# KrunkerAPI
A Simple Krunker api with GoLang

## Player Profile Request Map

### How to use API
Simply just type
```
go get github.com/somonox/KrunkerAPI@latest
```
in your project directory.


### Main Fields

- `player_name`: Player's name.
- `player_clan`: Clan name.
- `player_kills`: Total number of kills by the player.
- `player_deaths`: Total number of deaths.
- `player_score`: Player's total score.
- `player_timeplayed`: Total playtime in milliseconds.
- `player_games_played`: Total number of games played.
- `player_wins`: Total games won.
- `player_funds`: KR (in-game currency) owned by the player.
- `player_skinvalue`: Total value of owned skins.
- `player_datenew`: Date the player registered.
- `player_followed`: Number of followers.
- `player_following`: Number of people the player is following.
- `player_elo4`: ranked (season 8) elo (1 2 3 was Yendis ranked).
- `player_hack`: Flags a player for hacking 1 is hacker tagged, 0 is not hacker tagged.

### Unclear Fields

- `clan_rank`: Possibly the rank within the clan.
- `partner_approved`: Related to verified badges.
- `player_alias`: Premium alias, if applicable.
- `player_badge`: Represents player's badges.
- `player_calls`: Total reports to KPD.
- `player_chal`: Old challenge system (lvl 30 max).
- `player_elo`: Old ranked 1v1.
- `player_elo2`: Old ranked 2v2.
- `player_eventcount`: Unclear purpose.
- `player_featured`: Unknown usage.
- `player_id`: Player ID, UID.
- `player_infected`: Old feature, if someone knife you, you got infected.
- `player_jobrating`: total calls (helpful + unhelpful).
- `player_jobratingpositive`: Helpful reports.
- `player_premium`: Premium badge status.
- `player_region`: A flag selected by the user.
- `player_taggedaccounts`: "Bans from calls", if more than 1 player call the same cheater but got banned, everyone got this value added.
- `player_twitchname`: Linked Twitch account name, if any.
- `player_type`: Unclear usage.

### Player Stats

- `n`: Number of nukes.
- `s`: Total shots fired.
- `h`: Total hits.
- `hs`: Total headshots.
- `ls`: Total legshots.
- `wb`: Total wall bangs.
- `mk`: Total melee kills.
- `tmk`: Total throwing melee kills.
- `fk`: Total fist kills.
- `spry`: Total sprays used.
- `crc`: Total crouches.
- `sl`: Total slimers.

### Unknown `player_stats` Keys

- `c`, `c1`, `c2`, `ast`, `c5`, `c8`, `r2`, `c0`, `c4`, `c12`, `c11`, `r4`, `r3`, `c3`, `c6`, `flg`, `c7`, `c9`, `c15`, `sad`, `c13`, `r1`, `bdg` (likely matches `player_badge`), `ad`, `cad`, `c10`, `c14`, `r5`: Unclear or unknown usages.

---

## API Usage Example

Initialize using `NewKrunkerAPI`.  
Return value: `KrunkerAPI`, `error`.

To reserve structure destruction, use `defer api.Close()`.

### Fetching Player Profile

`GetProfile(username string)`

Return value: `*Profile`, `message`.  
This returns a pointer to the `Profile` object along with the original `decodedMessage`.
If the profile does not exist, it returns nill.


The keys and values from `decodedMessage` are outlined above, but some fields remain unclear. Help would be appreciated for those.

### Example Code

```go
package main

import (
	"log"

	"github.com/somonox/KrunkerAPI"
)

func main() {
	api, err := KrunkerAPI.NewKrunkerAPI()
	if err != nil {
		log.Fatal(err)
	}
	defer api.Close()

	profile, rawData := api.GetProfile("a6a6")
	if profile == nil {
		log.Fatal("Failed to get profile")
	}

	log.Println("Profile:", *profile)
	log.Println("Raw data:", *rawData)
}
```

# Disclaimer
This project is developed purely for fun and educational purposes.

I do not plan to maintain this project regularly.
If you find any bugs or issues, feel free to open an issue — I’ll try to address it when I can.
Contributions are also welcome!
