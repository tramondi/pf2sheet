package dto

type Sheet struct {
	ID         int       `json:"id"`
	Ancestry   *Ancestry `json:"ancestry,omitempty"`
	Class      *Class    `json:"class,omitempty"`
	Background *string   `json:"background,omitempty"`
	FullName   *string   `json:"full_name,omitempty"`
	Level      *int16    `json:"level,omitempty"`
	HpCurrent  *int16    `json:"hp_current,omitempty"`
	HpMax      *int16    `json:"hp_max,omitempty"`
}
