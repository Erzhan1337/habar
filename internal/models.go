package internal

import (
	"fmt"
	"strings"
	"time"
)

type Category string
type BaseType string
type Sweetener string
type IngredientType string
type AdditiveCategory string
type FunctionalAdditive string
type Allergen string
type FlagType string
type CaffeineType string
type CarbonationType string

const (
	CatWater           Category = "WATER"
	CatSoftDrink       Category = "SOFT_DRINK"
	CatJuice           Category = "JUICE"
	CatDairyDrink      Category = "DAIRY_DRINK"
	CatPlantBased      Category = "PLANT_BASED"
	CatFunctionalDrink Category = "FUNCTIONAL_DRINK"
	CatEnergyDrink     Category = "ENERGY_DRINK"
	CatTeaDrink        Category = "TEA_DRINK"
	CatCoffeeDrink     Category = "COFFEE_DRINK"
	CatAlcoholic       Category = "ALCOHOLIC"
	CatFermented       Category = "FERMENTED"
)

const (
	WATER_BASED     BaseType = "WATER_BASED"
	MILK_BASED      BaseType = "MILK_BASED"
	PLANT_BASED     BaseType = "PLANT_BASED"
	JUICE_BASED     BaseType = "JUICE_BASED"
	ALCOHOL_BASED   BaseType = "ALCOHOL_BASED"
	FERMENTED_BASED BaseType = "FERMENTED_BASED"
)

const (
	SUGAR                  Sweetener = "SUGAR"
	GLUCOSE_FRUCTOSE_SYRUP Sweetener = "GLUCOSE_FRUCTOSE_SYRUP"
	HONEY                  Sweetener = "HONEY"
	STEVIA                 Sweetener = "STEVIA"
	ASPARTAME              Sweetener = "ASPARTAME"
	SUCRALOSE              Sweetener = "SUCRALOSE"
	ACESULFAME_K           Sweetener = "ACESULFAME_K"
	SACCHARIN              Sweetener = "SACCHARIN"
	NO_SWEETENER           Sweetener = "NO_SWEETENER"
)

const (
	INGREDIENT_BASE       IngredientType = "BASE"
	INGREDIENT_SWEETENER  IngredientType = "SWEETENER"
	INGREDIENT_ADDITIVE   IngredientType = "ADDITIVE"
	INGREDIENT_FUNCTIONAL IngredientType = "FUNCTIONAL"
	INGREDIENT_ALLERGEN   IngredientType = "ALLERGEN"
	INGREDIENT_OTHER      IngredientType = "OTHER"
)

const (
	COLORANT        AdditiveCategory = "COLORANT"
	PRESERVATIVE    AdditiveCategory = "PRESERVATIVE"
	ANTIOXIDANT     AdditiveCategory = "ANTIOXIDANT"
	ACIDIFIER       AdditiveCategory = "ACIDIFIER"
	EMULSIFIER      AdditiveCategory = "EMULSIFIER"
	STABILIZER      AdditiveCategory = "STABILIZER"
	FLAVOR_ENHANCER AdditiveCategory = "FLAVOR_ENHANCER"
	SWEETENER       AdditiveCategory = "SWEETENER"
	THICKENER       AdditiveCategory = "THICKENER"
)

const (
	VITAMINS     FunctionalAdditive = "VITAMINS"
	MINERALS     FunctionalAdditive = "MINERALS"
	ELECTROLYTES FunctionalAdditive = "ELECTROLYTES"
	PROTEIN      FunctionalAdditive = "PROTEIN"
	PROBIOTICS   FunctionalAdditive = "PROBIOTICS"
	PREBIOTICS   FunctionalAdditive = "PREBIOTICS"
	COLLAGEN     FunctionalAdditive = "COLLAGEN"
	BCAA         FunctionalAdditive = "BCAA"
	TAURINE      FunctionalAdditive = "TAURINE"
	L_CARNITINE  FunctionalAdditive = "L_CARNITINE"
)
const (
	HIGH_SUGAR                    FlagType = "HIGH_SUGAR"
	VERY_HIGH_SUGAR               FlagType = "VERY_HIGH_SUGAR"
	ZERO_SUGAR                    FlagType = "ZERO_SUGAR"
	CONTAINS_MULTIPLE_ADDITIVES   FlagType = "CONTAINS_MULTIPLE_ADDITIVES"
	CONTAINS_SYNTHETIC_COLOR      FlagType = "CONTAINS_SYNTHETIC_COLOR"
	CONTAINS_ARTIFICIAL_SWEETENER FlagType = "CONTAINS_ARTIFICIAL_SWEETENER"
	CONTAINS_CAFFEINE             FlagType = "CONTAINS_CAFFEINE"
	ENERGY_CATEGORY               FlagType = "ENERGY_CATEGORY"
	FERMENTED_FLAG                FlagType = "FERMENTED"
	ULTRA_PROCESSED               FlagType = "ULTRA_PROCESSED"
)

const (
	NO_CAFFEINE      CaffeineType = "NO_CAFFEINE"
	NATURAL_CAFFEINE CaffeineType = "NATURAL_CAFFEINE"
	ADDED_CAFFEINE   CaffeineType = "ADDED_CAFFEINE"
)
const (
	STILL              CarbonationType = "STILL"
	CARBONATED         CarbonationType = "CARBONATED"
	LIGHTLY_CARBONATED CarbonationType = "LIGHTLY_CARBONATED"
)

const (
	MILK   Allergen = "MILK"
	SOY    Allergen = "SOY"
	GLUTEN Allergen = "GLUTEN"
	NUTS   Allergen = "NUTS"
)

type Flag struct {
	ID        uint     `gorm:"primaryKey;column:id"`
	ProductID uint     `json:"product_id" gorm:"column:product_id;index;not null"`
	FlagType  FlagType `json:"flag_type" gorm:"column:flag_type;type:varchar(64);not null"`
	Value     bool     `json:"value" gorm:"column:value;not null"`
}

type Beverage struct {
	ProductID    uint            `json:"product_id" gorm:"primaryKey;column:product_id"`
	Barcode      string          `json:"barcode" gorm:"column:barcode;uniqueIndex;not null"`
	Brand        string          `json:"brand" gorm:"column:brand;not null"`
	Name         string          `json:"name" gorm:"column:name;not null"`
	VolumeML     int             `json:"volume_ml" gorm:"column:volume_ml;not null"`
	CategoryLvl1 Category        `json:"category_lvl_1" gorm:"column:category_lvl_1;type:varchar(50);not null"`
	BaseType     BaseType        `json:"base_type" gorm:"column:base_type;type:varchar(50);not null"`
	CaffeineType CaffeineType    `json:"caffeine_type" gorm:"column:caffeine_type;type:varchar(50);not null"`
	Carbonation  CarbonationType `json:"carbonation" gorm:"column:carbonation;type:varchar(50);not null"`

	Nutrition   Nutrition    `json:"nutrition" gorm:"foreignKey:ProductID;references:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Ingredients []Ingredient `json:"ingredients" gorm:"foreignKey:ProductID;references:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Flags       []Flag       `json:"flags" gorm:"foreignKey:ProductID;references:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsVerified  bool         `json:"is_verified" gorm:"default:false"`
	DataStatus  string       `json:"data_status" gorm:"type:varchar(20);default:'READY'"`
	SourceURL   string       `json:"source_url"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type Ingredient struct {
	ID               uint                `gorm:"primaryKey;column:id"`
	ProductID        uint                `json:"product_id" gorm:"column:product_id;index;not null"`
	IngredientName   string              `json:"ingredient_name" gorm:"column:ingredient_name;not null"`
	Type             IngredientType      `json:"type" gorm:"column:type;type:varchar(32);not null"`
	ENumber          *string             `json:"e_number" gorm:"column:e_number"`
	Category         string              `json:"category" gorm:"column:category;type:varchar(64)"`
	AdditiveName     string              `json:"additive_name,omitempty" gorm:"column:additive_name;type:varchar(128)"`
	Sweetener        *Sweetener          `json:"sweetener,omitempty" gorm:"column:sweetener;type:varchar(64)"`
	Allergen         *Allergen           `json:"allergen,omitempty" gorm:"column:allergen;type:varchar(32)"`
	Functional       *FunctionalAdditive `json:"functional_additive,omitempty" gorm:"column:functional_additive;type:varchar(32)"`
	AdditiveCategory *AdditiveCategory   `json:"additive_category,omitempty" gorm:"column:additive_category;type:varchar(32)"`
}

type Nutrition struct {
	ID                uint     `gorm:"primaryKey;column:id"`
	ProductID         uint     `json:"product_id" gorm:"column:product_id;uniqueIndex;not null"`
	SugarG100ml       float64  `json:"sugar_g_100ml" gorm:"column:sugar_g_100ml"`
	TotalSugarG100    *float64 `json:"total_sugar_g_100ml,omitempty" gorm:"-"`
	CaffeineMg100ml   float64  `json:"caffeine_mg_100ml" gorm:"column:caffeine_mg_100ml"`
	SodiumMg          float64  `json:"sodium_mg" gorm:"column:sodium_mg"`
	AlcoholPercent    float64  `json:"alcohol_%" gorm:"column:alcohol_percent"`
	AlcoholPercentAlt *float64 `json:"alcohol_percent,omitempty" gorm:"-"`
}

func (Beverage) TableName() string {
	return "beverage_master"
}

func (Ingredient) TableName() string {
	return "beverage_ingredients"
}

func (Nutrition) TableName() string {
	return "beverage_nutrition"
}

func (Flag) TableName() string {
	return "beverage_flags"
}

func (b *Beverage) NormalizeAndValidate() error {
	b.Barcode = strings.TrimSpace(b.Barcode)
	b.Brand = normalizeText(b.Brand)
	b.Name = normalizeText(b.Name)
	if b.VolumeML <= 0 {
		return fmt.Errorf("volume_ml must be > 0")
	}
	if b.Barcode == "" || b.Brand == "" || b.Name == "" {
		return fmt.Errorf("barcode, brand and name are required")
	}

	if !isAllowedCategory(b.CategoryLvl1) {
		return fmt.Errorf("invalid category_lvl_1: %s", b.CategoryLvl1)
	}
	if !isAllowedBaseType(b.BaseType) {
		return fmt.Errorf("invalid base_type: %s", b.BaseType)
	}
	if !isAllowedCaffeineType(b.CaffeineType) {
		return fmt.Errorf("invalid caffeine_type: %s", b.CaffeineType)
	}
	if !isAllowedCarbonation(b.Carbonation) {
		return fmt.Errorf("invalid carbonation: %s", b.Carbonation)
	}
	if b.Nutrition.SugarG100ml < 0 || b.Nutrition.CaffeineMg100ml < 0 || b.Nutrition.SodiumMg < 0 || b.Nutrition.AlcoholPercent < 0 {
		return fmt.Errorf("nutrition values cannot be negative")
	}

	if b.Nutrition.TotalSugarG100 != nil {
		b.Nutrition.SugarG100ml = *b.Nutrition.TotalSugarG100
	}
	if b.Nutrition.AlcoholPercentAlt != nil {
		b.Nutrition.AlcoholPercent = *b.Nutrition.AlcoholPercentAlt
	}

	for i := range b.Ingredients {
		ing := &b.Ingredients[i]
		ing.IngredientName = normalizeIngredientName(ing.IngredientName)
		if ing.IngredientName == "" {
			return fmt.Errorf("ingredient_name is required")
		}
		ing.Category = normalizeText(ing.Category)
		ing.AdditiveName = normalizeText(ing.AdditiveName)
		if !isAllowedIngredientType(ing.Type) {
			return fmt.Errorf("invalid ingredient type: %s", ing.Type)
		}
		if ing.ENumber != nil {
			v := normalizeText(*ing.ENumber)
			if v == "" {
				ing.ENumber = nil
			} else {
				ing.ENumber = &v
			}
		}
		if ing.Sweetener != nil && !isAllowedSweetener(*ing.Sweetener) {
			return fmt.Errorf("invalid sweetener: %s", *ing.Sweetener)
		}
		if ing.Allergen != nil && !isAllowedAllergen(*ing.Allergen) {
			return fmt.Errorf("invalid allergen: %s", *ing.Allergen)
		}
		if ing.Functional != nil && !isAllowedFunctional(*ing.Functional) {
			return fmt.Errorf("invalid functional_additive: %s", *ing.Functional)
		}
		if ing.AdditiveCategory != nil && !isAllowedAdditiveCategory(*ing.AdditiveCategory) {
			return fmt.Errorf("invalid additive_category: %s", *ing.AdditiveCategory)
		}
		if ing.AdditiveCategory != nil && ing.AdditiveName == "" {
			ing.AdditiveName = ing.IngredientName
		}
	}
	return nil
}

func normalizeText(v string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(v)), " ")
}

func normalizeIngredientName(v string) string {
	clean := strings.ToLower(normalizeText(v))
	if clean == "" {
		return ""
	}
	aliases := map[string]string{
		"sugar":                  "SUGAR",
		"sucrose":                "SUGAR",
		"table sugar":            "SUGAR",
		"glucose-fructose syrup": "GLUCOSE_FRUCTOSE_SYRUP",
		"glucose fructose syrup": "GLUCOSE_FRUCTOSE_SYRUP",
		"hfcs":                   "GLUCOSE_FRUCTOSE_SYRUP",
		"honey":                  "HONEY",
		"stevia":                 "STEVIA",
		"aspartame":              "ASPARTAME",
		"sucralose":              "SUCRALOSE",
		"acesulfame k":           "ACESULFAME_K",
		"acesulfame-k":           "ACESULFAME_K",
		"saccharin":              "SACCHARIN",
		"water":                  "WATER",
		"carbon dioxide":         "CARBON_DIOXIDE",
		"co2":                    "CARBON_DIOXIDE",
		"caffeine":               "CAFFEINE",
		"taurine":                "TAURINE",
		"l-carnitine":            "L_CARNITINE",
	}
	if canonical, ok := aliases[clean]; ok {
		return canonical
	}
	return strings.ToUpper(strings.ReplaceAll(clean, " ", "_"))
}

func isAllowedCategory(v Category) bool {
	switch v {
	case CatWater, CatSoftDrink, CatJuice, CatDairyDrink, CatPlantBased, CatFunctionalDrink, CatEnergyDrink, CatTeaDrink, CatCoffeeDrink, CatAlcoholic, CatFermented:
		return true
	default:
		return false
	}
}

func isAllowedBaseType(v BaseType) bool {
	switch v {
	case WATER_BASED, MILK_BASED, PLANT_BASED, JUICE_BASED, ALCOHOL_BASED, FERMENTED_BASED:
		return true
	default:
		return false
	}
}

func isAllowedSweetener(v Sweetener) bool {
	switch v {
	case SUGAR, GLUCOSE_FRUCTOSE_SYRUP, HONEY, STEVIA, ASPARTAME, SUCRALOSE, ACESULFAME_K, SACCHARIN, NO_SWEETENER:
		return true
	default:
		return false
	}
}

func isAllowedIngredientType(v IngredientType) bool {
	switch v {
	case INGREDIENT_BASE, INGREDIENT_SWEETENER, INGREDIENT_ADDITIVE, INGREDIENT_FUNCTIONAL, INGREDIENT_ALLERGEN, INGREDIENT_OTHER:
		return true
	default:
		return false
	}
}

func isAllowedAdditiveCategory(v AdditiveCategory) bool {
	switch v {
	case COLORANT, PRESERVATIVE, ANTIOXIDANT, ACIDIFIER, EMULSIFIER, STABILIZER, FLAVOR_ENHANCER, SWEETENER, THICKENER:
		return true
	default:
		return false
	}
}

func isAllowedFunctional(v FunctionalAdditive) bool {
	switch v {
	case VITAMINS, MINERALS, ELECTROLYTES, PROTEIN, PROBIOTICS, PREBIOTICS, COLLAGEN, BCAA, TAURINE, L_CARNITINE:
		return true
	default:
		return false
	}
}

func isAllowedAllergen(v Allergen) bool {
	switch v {
	case MILK, SOY, GLUTEN, NUTS:
		return true
	default:
		return false
	}
}

func isAllowedCaffeineType(v CaffeineType) bool {
	switch v {
	case NO_CAFFEINE, NATURAL_CAFFEINE, ADDED_CAFFEINE:
		return true
	default:
		return false
	}
}

func isAllowedCarbonation(v CarbonationType) bool {
	switch v {
	case STILL, CARBONATED, LIGHTLY_CARBONATED:
		return true
	default:
		return false
	}
}
