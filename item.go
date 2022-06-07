package main

type Item struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Type         string        `json:"type"`
	Level        int           `json:"level"`
	Rarity       string        `json:"rarity"`
	VendorValue  int           `json:"vendor_value"`
	GameTypes    []string      `json:"game_types"`
	Flags        []interface{} `json:"flags"`
	Restrictions []interface{} `json:"restrictions"`
	ID           int           `json:"id"`
	ChatLink     string        `json:"chat_link"`
	Icon         string        `json:"icon"`
	Details      struct {
		Type                 string        `json:"type"`
		Flags                []string      `json:"flags"`
		InfusionUpgradeFlags []interface{} `json:"infusion_upgrade_flags"`
		AttributeAdjustment  int           `json:"attribute_adjustment"`
		InfixUpgrade         struct {
			ID   int `json:"id"`
			Buff struct {
				SkillID     int    `json:"skill_id"`
				Description string `json:"description"`
			} `json:"buff"`
			Attributes []struct {
				Attribute string `json:"attribute"`
				Modifier  int    `json:"modifier"`
			} `json:"attributes"`
		} `json:"infix_upgrade"`
		Suffix string `json:"suffix"`
	} `json:"details"`
}