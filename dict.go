package main

type Dictionary struct {
	converter       *Converter
	unitTypes       map[int]UnitType
	unitTypesBySlug map[string]UnitType
	unitSystems     map[int]UnitSystem
}

type UnitType struct {
	Name string
	Slug string
}

type UnitSystem struct {
	Name string
}

func NewDictionary(c *Converter) Dictionary {
	d := Dictionary{}
	d.converter = c

	d.unitTypes = map[int]UnitType{}

	d.unitTypes[UNIT_TYPE_LENGTH] = UnitType{
		"Длинна, расстояния",
		"length",
	}

	d.unitTypes[UNIT_TYPE_WEIGHT] = UnitType{
		"Веса",
		"weight",
	}

	d.unitTypesBySlug = map[string]UnitType{}
	for _, ut := range d.unitTypes {
		d.unitTypesBySlug[ut.Slug] = ut
	}

	d.unitSystems = map[int]UnitSystem{}
	d.unitSystems[LENGTH_METRIC] = UnitSystem{
		"Метричкеская",
	}

	d.unitSystems[LENGTH_BRITISH] = UnitSystem{
		"Американская и британская",
	}

	d.unitSystems[LENGTH_NAUTICAL] = UnitSystem{
		"Международные морские еденицы",
	}

	d.unitSystems[LENGTH_SLAVIC] = UnitSystem{
		"Древнерусская",
	}

	d.unitSystems[WEIGHT_METRIC] = UnitSystem{
		"Метрическая система",
	}

	d.unitSystems[WEIGHT_TROY] = UnitSystem{
		"Тройская система",
	}

	return d
}

func (d *Dictionary) getUnitTypesBySlug(slug string) UnitType {
	return d.unitTypesBySlug[slug]
}

func (d *Dictionary) getTypeBySlug(t string) int {
	for k, ut := range d.unitTypes {
		if ut.Slug == t {
			return k
		}
	}

	return -1
}
