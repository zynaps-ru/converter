package main

const (
	LENGTH_METRIC = iota
	LENGTH_BRITISH
	LENGTH_NAUTICAL

	WEIGHT_METRIC = iota
	WEIGHT_TROY

	UNIT_TYPE_LENGTH = iota
	UNIT_TYPE_WEIGHT
)

type Converter struct {
	mainUnits map[int]Unit
	units     map[int]map[string]Unit
}

func (c *Converter) Convert(v float64, from, to string, unitType int) float64 {
	fromUnit := c.units[unitType][from]
	toUnit := c.units[unitType][to]

	return fromUnit.FromMain(toUnit.ToMain(v))
}

func NewConverter() Converter {
	c := Converter{}
	c.units = map[int]map[string]Unit{}

	millimeter := Unit{
		"милиметр",
		false,
		10000,
		UNIT_TYPE_LENGTH,
		LENGTH_METRIC,
	}

	centimeter := Unit{
		"сантиметр",
		true,
		100,
		UNIT_TYPE_LENGTH,
		LENGTH_METRIC,
	}

	meter := Unit{
		"метр",
		false,
		1,
		UNIT_TYPE_LENGTH,
		LENGTH_METRIC,
	}

	kiloMeter := Unit{
		"километр",
		false,
		0.001,
		UNIT_TYPE_LENGTH,
		LENGTH_METRIC,
	}

	decimeter := Unit{
		"дециметр",
		false,
		10,
		UNIT_TYPE_LENGTH,
		LENGTH_METRIC,
	}

	mile := Unit{
		"миля",
		false,
		0.000621371,
		UNIT_TYPE_LENGTH,
		LENGTH_BRITISH,
	}

	foot := Unit{
		"фут",
		false,
		3.28084,
		UNIT_TYPE_LENGTH,
		LENGTH_BRITISH,
	}

	inch := Unit{
		"дюйм",
		false,
		39.3701,
		UNIT_TYPE_LENGTH,
		LENGTH_BRITISH,
	}

	nauticalMile := Unit{
		"морская лига",
		false,
		0.000539957,
		UNIT_TYPE_LENGTH,
		LENGTH_NAUTICAL,
	}

	nauticalLeague := Unit{
		"морская лига",
		false,
		0.00018,
		UNIT_TYPE_LENGTH,
		LENGTH_NAUTICAL,
	}

	cable := Unit{
		"кабельт",
		false,
		0.00018,
		UNIT_TYPE_LENGTH,
		LENGTH_NAUTICAL,
	}

	yard := Unit{
		"ярд",
		false,
		0.9144,
		UNIT_TYPE_LENGTH,
		LENGTH_BRITISH,
	}

	kilogram := Unit{
		"килограмм",
		true,
		1,
		UNIT_TYPE_WEIGHT,
		WEIGHT_METRIC,
	}

	gramm := Unit{
		"грамм",
		false,
		1000,
		UNIT_TYPE_WEIGHT,
		WEIGHT_METRIC,
	}

	millgramm := Unit{
		"милиграмм",
		false,
		1000000,
		UNIT_TYPE_WEIGHT,
		WEIGHT_METRIC,
	}

	ton := Unit{
		"тонн",
		false,
		0.001,
		UNIT_TYPE_WEIGHT,
		WEIGHT_METRIC,
	}
	funt := Unit{
		"фунт",
		false,
		2.20462,
		UNIT_TYPE_WEIGHT,
		WEIGHT_TROY,
	}
	uncia := Unit{
		"унция",
		false,
		35.274,
		UNIT_TYPE_WEIGHT,
		WEIGHT_TROY,
	}

	units := []Unit{millimeter, centimeter, meter, kiloMeter, yard, kilogram, ton, decimeter, mile, foot, inch,
		nauticalMile, nauticalLeague, cable, gramm, millgramm, uncia, funt,
	}

	c.mainUnits = map[int]Unit{}
	for _, u := range units {
		if u.isMain {
			c.mainUnits[u.unitType] = u
		}

		_, ok := c.units[u.unitType]
		if !ok {
			c.units[u.unitType] = map[string]Unit{}
		}

		c.units[u.unitType][u.Key] = u

	}

	return c
}
