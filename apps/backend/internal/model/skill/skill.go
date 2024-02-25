package skill

type Lvl string

const (
	LvlNone    Lvl = "none"
	LvlTrained Lvl = "trained"
	LvlExpert  Lvl = "expert"
	LvlMaster  Lvl = "master"
	LvlLegend  Lvl = "legend"
)

type Skill struct {
	Lvl

	Title           string
	KeyAbilityBonus int
	Item            int
	Armor           int

	charLvlBonus int
}
