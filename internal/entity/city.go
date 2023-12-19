package entity

type City struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name    string `json:"name" gorm:"not null"`
	StateID string `json:"stateId" gorm:"not null"`
	IBGE    string `json:"ibge" gorm:"not null"`
	LatLon  string `json:"latLon"`
	CodTom  uint   `json:"codTom"`

	State State `json:"state" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
