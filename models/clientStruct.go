package models

type ClientSession struct {
	SessionKey                    string `json:"sessionkey"`
	LastLoginTime                 uint32 `json:"lastlogintime"`
	IsPremium                     bool   `json:"ispremium"`
	PremiumUntil                  uint64 `json:"premiumuntil"`
	OptionTracking                bool   `json:"optiontracking"`
	Status                        string `json:"status"`
	ReturnerNotification          bool   `json:"returnernotification"`
	ShowRewardNews                bool   `json:"showrewardnews"`
	IsReturner                    bool   `json:"isreturner"`
	FpsTracking                   bool   `json:"fpstracking"`
	TournamentTicketPurchaseState uint8  `json:"tournamentticketpurchasestate"`
	EmailCodeRequest              bool   `json:"emailcoderequest"`
}

type ClientWorld struct {
	ID                         uint8  `json:"id"`
	Name                       string `json:"name"`
	ExternalAddress            string `json:"externaladdress"`
	ExternalPort               uint16 `json:"externalport"`
	PreviewState               uint   `json:"previewstate"`
	Location                   string `json:"location"`
	AntiCheatProtection        bool   `json:"anticheatprotection"`
	ExternalAddRessUnProtected string `json:"externaladdressunprotected"`
	ExternalAddressProtected   string `json:"externaladdressprotected"`
	PvpType                    uint8  `json:"pvptype"`
	ExternalPortProtected      uint16 `json:"externalportprotected"`
	ExternalPortUnprotected    uint16 `json:"externalportunprotected"`
	IsTournamentWorld          bool   `json:"istournamentworld"`
	RestrictedStore            bool   `json:"restrictedstore"`
	CurrentTournamentPhase     uint8  `json:"currenttournamentphase"`
}

type ClientCharacters struct {
	WorldID                          uint   `json:"worldid"`
	Name                             string `json:"name"`
	IsMale                           bool   `json:"ismale"`
	Tutorial                         bool   `json:"tutorial"`
	Level                            uint16 `json:"level"`
	Vocation                         uint8  `json:"vocation"`
	OutfitID                         uint16 `json:"outfitid"`
	HeadColor                        uint16 `json:"headcolor"`
	TorsoColor                       uint16 `json:"torsocolor"`
	LegsColor                        uint16 `json:"legscolor"`
	DetailColor                      uint16 `json:"detailcolor"`
	AddonsFlags                      uint8  `json:"addonsflags"`
	IsHidden                         bool   `json:"ishidden"`
	IsTournamentParticipant          bool   `json:"istournamentparticipant"`
	RemainIngDailyTournamentPlayTime uint16 `json:"remainingdailytournamentplaytime"`
	IsMainCharacter                  bool   `json:"ismaincharacter"`
	DailyRewardState                 uint16 `json:"dailyrewardstate"`
}
