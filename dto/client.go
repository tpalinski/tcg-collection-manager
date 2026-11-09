package dto

type PokemonCardDTO struct {
	Category       string            `json:"category"`
	ID             string            `json:"id"`
	Illustrator    string            `json:"illustrator"`
	Image          string            `json:"image"`
	LocalID        string            `json:"localId"`
	Name           string            `json:"name"`
	Rarity         string            `json:"rarity"`
	Set            CardSetDTO        `json:"set"`
	Variants       CardVariantsDTO   `json:"variants"`
	HP             *int              `json:"hp"`
	Types          []string          `json:"types"`
	EvolveFrom     *string           `json:"evolveFrom"`
	Description    *string           `json:"description"`
	Stage          *string           `json:"stage"`
	Attacks        []CardAttackDTO   `json:"attacks"`
	Weaknesses     []CardWeaknessDTO `json:"weaknesses"`
	Retreat        *int              `json:"retreat"`
	RegulationMark *string           `json:"regulationMark"`
	Legal          CardLegalDTO      `json:"legal"`
}

type CardSetDTO struct {
	CardCount CardCountDTO `json:"cardCount"`
	ID        string       `json:"id"`
	Logo      string       `json:"logo"`
	Name      string       `json:"name"`
	Symbol    string       `json:"symbol"`
}

type CardCountDTO struct {
	Official int `json:"official"`
	Total    int `json:"total"`
}

type CardVariantsDTO struct {
	FirstEdition bool `json:"firstEdition"`
	Holo         bool `json:"holo"`
	Normal       bool `json:"normal"`
	Reverse      bool `json:"reverse"`
	WPromo       bool `json:"wPromo"`
}

type CardAttackDTO struct {
	Cost   []string `json:"cost"`
	Name   string   `json:"name"`
	Effect string   `json:"effect"`
	Damage *int     `json:"damage"`
}

type CardWeaknessDTO struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CardLegalDTO struct {
	Standard bool `json:"standard"`
	Expanded bool `json:"expanded"`
}
