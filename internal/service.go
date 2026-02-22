package internal

import "strings"

func AnalyzeBeverage(b *Beverage) []Flag {
	var flags []Flag

	if b.Nutrition.SugarG100ml > 10 {
		flags = append(flags, Flag{FlagType: VERY_HIGH_SUGAR, Value: true})
	} else if b.Nutrition.SugarG100ml > 5 {
		flags = append(flags, Flag{FlagType: HIGH_SUGAR, Value: true})
	} else if b.Nutrition.SugarG100ml < 0.5 {
		flags = append(flags, Flag{FlagType: ZERO_SUGAR, Value: true})
	}

	if b.Nutrition.CaffeineMg100ml > 0 || b.CaffeineType != NO_CAFFEINE {
		flags = append(flags, Flag{FlagType: CONTAINS_CAFFEINE, Value: true})
	}

	if b.CategoryLvl1 == CatEnergyDrink {
		flags = append(flags, Flag{FlagType: ENERGY_CATEGORY, Value: true})
	}

	if b.CategoryLvl1 == CatFermented || b.BaseType == FERMENTED_BASED {
		flags = append(flags, Flag{FlagType: FERMENTED_FLAG, Value: true})
	}

	eCount := 0
	hasSyntheticColor := false
	hasArtificialSweetener := false

	for _, ing := range b.Ingredients {
		if ing.Sweetener != nil {
			s := *ing.Sweetener
			if s == ASPARTAME || s == SUCRALOSE || s == ACESULFAME_K || s == SACCHARIN {
				hasArtificialSweetener = true
			}
		}
		if ing.ENumber != nil && *ing.ENumber != "" {
			eCount++
			if isSyntheticColor(ing) {
				hasSyntheticColor = true
			}
		}
	}

	if hasArtificialSweetener {
		flags = append(flags, Flag{FlagType: CONTAINS_ARTIFICIAL_SWEETENER, Value: true})
	}
	if eCount > 5 {
		flags = append(flags, Flag{FlagType: CONTAINS_MULTIPLE_ADDITIVES, Value: true})
	}
	if hasSyntheticColor {
		flags = append(flags, Flag{FlagType: CONTAINS_SYNTHETIC_COLOR, Value: true})
	}
	if eCount > 5 || hasArtificialSweetener {
		flags = append(flags, Flag{FlagType: ULTRA_PROCESSED, Value: true})
	}

	return flags
}

func isSyntheticColor(ing Ingredient) bool {
	if ing.AdditiveCategory == nil || *ing.AdditiveCategory != COLORANT || ing.ENumber == nil {
		return false
	}
	// Typical synthetic food colorants; natural ones are intentionally excluded.
	synthetic := map[string]struct{}{
		"E102": {}, "E104": {}, "E110": {}, "E122": {}, "E123": {},
		"E124": {}, "E127": {}, "E129": {}, "E131": {}, "E132": {},
		"E133": {}, "E142": {}, "E151": {}, "E155": {},
	}
	e := strings.ToUpper(strings.TrimSpace(*ing.ENumber))
	_, ok := synthetic[e]
	return ok
}
