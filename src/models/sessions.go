package models

import "gorm.io/gorm"

type Session struct {
	ID        string `gorm:"column:id;type:varchar(191);primaryKey;not null" json:"id"`
	AccountID uint   `gorm:"column:account_id;type:int unsigned;not null" json:"account_id"`
	Expires   uint64 `gorm:"column:expires;type:bigint unsigned;not null" json:"expires"`
}

func (Session) TableName() string {
	return "account_sessions"
}

type Account struct {
	ID                int       `gorm:"column:id;primaryKey;autoIncrement;type:int(11) unsigned" json:"id"`
	Name              string    `gorm:"column:name;type:varchar(32);not null" json:"name"`
	Password          string    `gorm:"column:password;type:text;not null" json:"password"`
	Email             string    `gorm:"column:email;type:varchar(255);not null;default:''" json:"email"`
	PremDays          int       `gorm:"column:premdays;type:int(11);not null;default:0" json:"premdays"`
	PremDaysPurchased int       `gorm:"column:premdays_purchased;type:int(11);not null;default:0" json:"premdays_purchased"`
	LastDay           uint      `gorm:"column:lastday;type:int(10) unsigned;not null;default:0" json:"lastday"`
	Type              uint8     `gorm:"column:type;type:tinyint(1) unsigned;not null;default:1" json:"type"`
	Coins             uint      `gorm:"column:coins;type:int(12) unsigned;not null;default:0" json:"coins"`
	CoinsTransferable uint      `gorm:"column:coins_transferable;type:int(12) unsigned;not null;default:0" json:"coins_transferable"`
	TournamentCoins   uint      `gorm:"column:tournament_coins;type:int(12) unsigned;not null;default:0" json:"tournament_coins"`
	Creation          uint      `gorm:"column:creation;type:int(11) unsigned;not null;default:0" json:"creation"`
	Recruiter         uint      `gorm:"column:recruiter;type:int(6);default:0" json:"recruiter"`
	HouseBidID        int       `gorm:"column:house_bid_id;type:int(11);not null;default:0" json:"house_bid_id"`
	Players           []Players `gorm:"foreignKey:AccountID" json:"players"`
}

type Players struct {
	ID   uint   `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);unique;not null" json:"name"`

	GroupID   int `gorm:"column:group_id;type:int(11);not null;default:1" json:"group_id"`
	AccountID int `gorm:"column:account_id;type:int(11) unsigned;not null;default:0" json:"account_id"`

	Level     int `gorm:"column:level;type:int(11);not null;default:1" json:"level"`
	Vocation  int `gorm:"column:vocation;type:int(11);not null;default:0" json:"vocation"`
	Health    int `gorm:"column:health;type:int(11);not null;default:150" json:"health"`
	HealthMax int `gorm:"column:healthmax;type:int(11);not null;default:150" json:"healthmax"`

	Experience uint64 `gorm:"column:experience;type:bigint(20);not null;default:0" json:"experience"`

	LookBody   int `gorm:"column:lookbody;type:int(11);not null;default:0" json:"lookbody"`
	LookFeet   int `gorm:"column:lookfeet;type:int(11);not null;default:0" json:"lookfeet"`
	LookHead   int `gorm:"column:lookhead;type:int(11);not null;default:0" json:"lookhead"`
	LookLegs   int `gorm:"column:looklegs;type:int(11);not null;default:0" json:"looklegs"`
	LookType   int `gorm:"column:looktype;type:int(11);not null;default:136" json:"looktype"`
	LookAddons int `gorm:"column:lookaddons;type:int(11);not null;default:0" json:"lookaddons"`

	MagLevel int `gorm:"column:maglevel;type:int(11);not null;default:0" json:"maglevel"`
	Mana     int `gorm:"column:mana;type:int(11);not null;default:0" json:"mana"`
	ManaMax  int `gorm:"column:manamax;type:int(11);not null;default:0" json:"manamax"`

	ManaSpent uint64 `gorm:"column:manaspent;type:bigint(20) unsigned;not null;default:0" json:"manaspent"`
	Soul      uint   `gorm:"column:soul;type:int(10) unsigned;not null;default:0" json:"soul"`

	TownID int `gorm:"column:town_id;type:int(11);not null;default:1" json:"town_id"`

	PosX int `gorm:"column:posx;type:int(11);not null;default:0" json:"posx"`
	PosY int `gorm:"column:posy;type:int(11);not null;default:0" json:"posy"`
	PosZ int `gorm:"column:posz;type:int(11);not null;default:0" json:"posz"`

	Conditions []byte `gorm:"column:conditions;type:mediumblob;not null" json:"conditions"`

	Cap int `gorm:"column:cap;type:int(11);not null;default:0" json:"cap"`

	Sex       int    `gorm:"column:sex;type:int(11);not null;default:0" json:"sex"`
	Pronoun   int    `gorm:"column:pronoun;type:int(11);not null;default:0" json:"pronoun"`
	LastLogin uint64 `gorm:"column:lastlogin;type:bigint(20) unsigned;not null;default:0" json:"lastlogin"`
	LastIP    uint32 `gorm:"column:lastip;type:int(10) unsigned;not null;default:0" json:"lastip"`

	Save      bool  `gorm:"column:save;type:tinyint(1);not null;default:1" json:"save"`
	Skull     int   `gorm:"column:skull;type:tinyint(1);not null;default:0" json:"skull"`
	SkullTime int64 `gorm:"column:skulltime;type:bigint(20);not null;default:0" json:"skulltime"`

	LastLogout uint64 `gorm:"column:lastlogout;type:bigint(20) unsigned;not null;default:0" json:"lastlogout"`
	Blessings  int    `gorm:"column:blessings;type:tinyint(2);not null;default:0" json:"blessings"`

	OnlineTime int64  `gorm:"column:onlinetime;type:int(11);not null;default:0" json:"onlinetime"`
	Deletion   int64  `gorm:"column:deletion;type:bigint(15);not null;default:0" json:"deletion"`
	Balance    uint64 `gorm:"column:balance;type:bigint(20) unsigned;not null;default:0" json:"balance"`

	OfflineTrainingTime  uint16 `gorm:"column:offlinetraining_time;type:smallint(5) unsigned;not null;default:43200" json:"offlinetraining_time"`
	OfflineTrainingSkill int    `gorm:"column:offlinetraining_skill;type:tinyint(2);not null;default:-1" json:"offlinetraining_skill"`
	Stamina              uint16 `gorm:"column:stamina;type:smallint(5) unsigned;not null;default:2520" json:"stamina"`

	SkillFist       uint   `gorm:"column:skill_fist;type:int(10) unsigned;not null;default:10" json:"skill_fist"`
	SkillFistTries  uint64 `gorm:"column:skill_fist_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_fist_tries"`
	SkillClub       uint   `gorm:"column:skill_club;type:int(10) unsigned;not null;default:10" json:"skill_club"`
	SkillClubTries  uint64 `gorm:"column:skill_club_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_club_tries"`
	SkillSword      uint   `gorm:"column:skill_sword;type:int(10) unsigned;not null;default:10" json:"skill_sword"`
	SkillSwordTries uint64 `gorm:"column:skill_sword_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_sword_tries"`

	SkillAxe      uint   `gorm:"column:skill_axe;type:int(10) unsigned;not null;default:10" json:"skill_axe"`
	SkillAxeTries uint64 `gorm:"column:skill_axe_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_axe_tries"`

	SkillDist      uint   `gorm:"column:skill_dist;type:int(10) unsigned;not null;default:10" json:"skill_dist"`
	SkillDistTries uint64 `gorm:"column:skill_dist_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_dist_tries"`

	SkillShielding      uint   `gorm:"column:skill_shielding;type:int(10) unsigned;not null;default:10" json:"skill_shielding"`
	SkillShieldingTries uint64 `gorm:"column:skill_shielding_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_shielding_tries"`

	SkillFishing      uint   `gorm:"column:skill_fishing;type:int(10) unsigned;not null;default:10" json:"skill_fishing"`
	SkillFishingTries uint64 `gorm:"column:skill_fishing_tries;type:bigint(20) unsigned;not null;default:0" json:"skill_fishing_tries"`
	World             string

	Account      Account       `gorm:"foreignKey:AccountID" json:"account"`
	PlayerDeaths []PlayerDeath `gorm:"foreignKey:PlayerID;references:ID" json:"deaths"`
	PlayerItems  []PlayerItem  `gorm:"foreignKey:PlayerID;references:ID" json:"items"`
}

type PlayerItem struct {
	// Clave primaria compuesta
	PlayerID int `gorm:"column:player_id;not null;default:0;primaryKey;index:idx_player_id"` // INDEX player_id
	PID      int `gorm:"column:pid;not null;default:0;primaryKey"`
	SID      int `gorm:"column:sid;not null;default:0;primaryKey;index:idx_sid"` // INDEX sid

	// Campos normales (DDL usa int(11) para ambos)
	ItemType int `gorm:"column:itemtype;not null;default:0"`
	Count    int `gorm:"column:count;not null;default:0"`
	// BLOB NOT NULL
	Attributes []byte `gorm:"column:attributes;type:blob;not null"`

	// Relación con players(id) + ON DELETE CASCADE
	Player Players `gorm:"foreignKey:PlayerID;references:ID;constraint:OnDelete:CASCADE;"`
}

func (PlayerItem) TableName() string { return "player_items" }
func (Players) TableName() string {
	return "players"
}

type PlayersOnline struct {
	PlayerID uint `gorm:"column:player_id"`
}

func (PlayersOnline) TableName() string {
	return "players_online"
}

type PlayerDeath struct {
	// FK a players(id)
	PlayerID int     `gorm:"column:player_id;not null;index:idx_player_id"` // INDEX player_id
	Player   Players `gorm:"foreignKey:PlayerID;references:ID;constraint:OnDelete:CASCADE;"`

	// bigint(20) unsigned NOT NULL DEFAULT 0
	Time uint64 `gorm:"column:time;type:bigint unsigned;not null;default:0"`

	// int(11) NOT NULL DEFAULT 1
	Level int `gorm:"column:level;type:int;not null;default:1"`

	// varchar(255) NOT NULL, INDEX
	KilledBy string `gorm:"column:killed_by;type:varchar(255);not null;index:idx_killed_by"`

	// tinyint(1) NOT NULL DEFAULT 1
	IsPlayer bool `gorm:"column:is_player;type:tinyint(1);not null;default:1"`

	// varchar(100) NOT NULL, INDEX
	MostDamageBy string `gorm:"column:mostdamage_by;type:varchar(100);not null;index:idx_mostdamage_by"`

	// tinyint(1) NOT NULL DEFAULT 0
	MostDamageIsPlayer bool `gorm:"column:mostdamage_is_player;type:tinyint(1);not null;default:0"`

	// tinyint(1) NOT NULL DEFAULT 0
	Unjustified bool `gorm:"column:unjustified;type:tinyint(1);not null;default:0"`

	// tinyint(1) NOT NULL DEFAULT 0
	MostDamageUnjustified bool `gorm:"column:mostdamage_unjustified;type:tinyint(1);not null;default:0"`

	// TEXT NOT NULL
	Participants string `gorm:"column:participants;type:text;not null"`
}

func (PlayerDeath) TableName() string { return "player_deaths" }

type PlayersNames struct {
	gorm.Model
	Name      string `gorm:"unique"`
	World     string
	AccountID int
}

type Towns struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name;type:varchar(255)"`
	Pos_x int    `gorm:"column:posx"`
	Pos_y int    `gorm:"column:posy"`
	Pos_z int    `gorm:"column:posz"`
}
