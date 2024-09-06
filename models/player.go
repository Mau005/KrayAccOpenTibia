package models

type Players struct {
	ID                   int     `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	Name                 string  `gorm:"column:name;size:255" json:"name"`
	GroupID              int     `gorm:"column:group_id;default:1" json:"group_id"`
	AccountID            int     `gorm:"column:account_id;default:0" json:"account_id"`
	Level                int     `gorm:"column:level;default:1" json:"level" yaml:"Level"`
	Vocation             int     `gorm:"column:vocation;default:0" json:"vocation" yaml:"Vocation"`
	Health               int     `gorm:"column:health;default:150" json:"health"`
	HealthMax            int     `gorm:"column:healthmax;default:150" json:"healthmax" yaml:"HealthMax"`
	Experience           uint64  `gorm:"column:experience;default:0" json:"experience" yaml:"Experience"`
	LookBody             int     `gorm:"column:lookbody;default:0" json:"lookbody"`
	LookFeet             int     `gorm:"column:lookfeet;default:0" json:"lookfeet"`
	LookHead             int     `gorm:"column:lookhead;default:0" json:"lookhead"`
	LookLegs             int     `gorm:"column:looklegs;default:0" json:"looklegs"`
	LookType             int     `gorm:"column:looktype;default:136" json:"looktype"`
	LookAddons           int     `gorm:"column:lookaddons;default:0" json:"lookaddons"`
	LookMount            int     `gorm:"column:lookmount;default:0" json:"lookmount"`
	LookMountHead        int     `gorm:"column:lookmounthead;default:0" json:"lookmounthead"`
	LookMountBody        int     `gorm:"column:lookmountbody;default:0" json:"lookmountbody"`
	LookMountLegs        int     `gorm:"column:lookmountlegs;default:0" json:"lookmountlegs"`
	LookMountFeet        int     `gorm:"column:lookmountfeet;default:0" json:"lookmountfeet"`
	RandomizeMount       bool    `gorm:"column:randomizemount;default:0" json:"randomizemount"`
	Direction            uint8   `gorm:"column:direction;default:2" json:"direction"`
	MagLevel             int     `gorm:"column:maglevel;default:0" json:"maglevel"`
	Mana                 int     `gorm:"column:mana;default:0" json:"mana"`
	ManaMax              int     `gorm:"column:manamax;default:0" json:"manamax" yaml:"ManaMax"`
	ManaSpent            uint64  `gorm:"column:manaspent;default:0" json:"manaspent"`
	Soul                 uint    `gorm:"column:soul;default:0" json:"soul"`
	TownID               int     `gorm:"column:town_id;default:1" json:"town_id" yaml:"TownID"`
	PosX                 int     `gorm:"column:posx;default:0" json:"posx"`
	PosY                 int     `gorm:"column:posy;default:0" json:"posy"`
	PosZ                 int     `gorm:"column:posz;default:0" json:"posz"`
	Conditions           *[]byte `gorm:"column:conditions;type:blob" json:"conditions"` // Utilizar puntero para permitir valores NULL
	Cap                  int     `gorm:"column:cap;default:400" json:"cap" yaml:"Cap"`
	Sex                  int     `gorm:"column:sex;default:0" json:"sex"`
	LastLogin            uint64  `gorm:"column:lastlogin;default:0" json:"lastlogin"`
	LastIP               []byte  `gorm:"column:lastip;size:16;default:0" json:"lastip"` // Utilizar []byte para varbinary
	Save                 bool    `gorm:"column:save;default:1" json:"save"`
	Skull                int     `gorm:"column:skull;default:0" json:"skull"`
	SkullTime            int64   `gorm:"column:skulltime;default:0" json:"skulltime"`
	LastLogout           uint64  `gorm:"column:lastlogout;default:0" json:"lastlogout"`
	Blessings            int     `gorm:"column:blessings;default:0" json:"blessings"`
	OnlineTime           int64   `gorm:"column:onlinetime;default:0" json:"onlinetime"`
	Deletion             int64   `gorm:"column:deletion;default:0" json:"deletion"`
	Balance              uint64  `gorm:"column:balance;default:0" json:"balance"`
	OfflineTrainingTime  uint16  `gorm:"column:offlinetraining_time;default:43200" json:"offlinetraining_time"`
	OfflineTrainingSkill int     `gorm:"column:offlinetraining_skill;default:-1" json:"offlinetraining_skill"`
	Stamina              uint16  `gorm:"column:stamina;default:2520" json:"stamina"`
	SkillFist            uint    `gorm:"column:skill_fist;default:10" json:"skill_fist"`
	SkillFistTries       uint64  `gorm:"column:skill_fist_tries;default:0" json:"skill_fist_tries"`
	SkillClub            uint    `gorm:"column:skill_club;default:10" json:"skill_club"`
	SkillClubTries       uint64  `gorm:"column:skill_club_tries;default:0" json:"skill_club_tries"`
	SkillSword           uint    `gorm:"column:skill_sword;default:10" json:"skill_sword"`
	SkillSwordTries      uint64  `gorm:"column:skill_sword_tries;default:0" json:"skill_sword_tries"`
	SkillAxe             uint    `gorm:"column:skill_axe;default:10" json:"skill_axe"`
	SkillAxeTries        uint64  `gorm:"column:skill_axe_tries;default:0" json:"skill_axe_tries"`
	SkillDist            uint    `gorm:"column:skill_dist;default:10" json:"skill_dist"`
	SkillDistTries       uint64  `gorm:"column:skill_dist_tries;default:0" json:"skill_dist_tries"`
	SkillShielding       uint    `gorm:"column:skill_shielding;default:10" json:"skill_shielding"`
	SkillShieldingTries  uint64  `gorm:"column:skill_shielding_tries;default:0" json:"skill_shielding_tries"`
	SkillFishing         uint    `gorm:"column:skill_fishing;default:10" json:"skill_fishing"`
	SkillFishingTries    uint64  `gorm:"column:skill_fishing_tries;default:0" json:"skill_fishing_tries"`
	Account              Account `gorm:"foreignKey:AccountID" json:"account"`
}

func (Players) TableName() string {
	return "players"
}
