package domain

type SimplifyMenu struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

type API struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:255;" json:"name"`
	Url     string `gorm:"size:255;" json:"url"`
	Methods string `gorm:"size:255;" json:"methods"`
}

type UpdateCasbinPolicyReq struct {
	OldPolicy []string `gorm:"size:255;" json:"old_policy"`
	NewPolicy []string `gorm:"size:255;" json:"new_policy"`
}

type AddCasbinRulePolicyReq struct {
	NewPolicy []string `gorm:"size:255;" json:"new_policy"`
}

type RemoveCasbinPolicyReq struct {
	RemovePolicy []string `gorm:"size:255;" json:"remove_policy"`
}
