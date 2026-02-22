package internal

import "gorm.io/gorm"

func EnsureConstraints(db *gorm.DB) error {
	stmts := []string{
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_beverage_ingredients_product') THEN
				ALTER TABLE beverage_ingredients
				ADD CONSTRAINT fk_beverage_ingredients_product
				FOREIGN KEY (product_id) REFERENCES beverage_master(product_id) ON DELETE CASCADE;
			END IF;
		END $$;`,
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_beverage_nutrition_product') THEN
				ALTER TABLE beverage_nutrition
				ADD CONSTRAINT fk_beverage_nutrition_product
				FOREIGN KEY (product_id) REFERENCES beverage_master(product_id) ON DELETE CASCADE;
			END IF;
		END $$;`,
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_beverage_flags_product') THEN
				ALTER TABLE beverage_flags
				ADD CONSTRAINT fk_beverage_flags_product
				FOREIGN KEY (product_id) REFERENCES beverage_master(product_id) ON DELETE CASCADE;
			END IF;
		END $$;`,
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'chk_beverage_master_category_lvl_1') THEN
				ALTER TABLE beverage_master
				ADD CONSTRAINT chk_beverage_master_category_lvl_1
				CHECK (category_lvl_1 IN ('WATER','SOFT_DRINK','JUICE','DAIRY_DRINK','PLANT_BASED','FUNCTIONAL_DRINK','ENERGY_DRINK','TEA_DRINK','COFFEE_DRINK','ALCOHOLIC','FERMENTED'));
			END IF;
		END $$;`,
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'chk_beverage_master_base_type') THEN
				ALTER TABLE beverage_master
				ADD CONSTRAINT chk_beverage_master_base_type
				CHECK (base_type IN ('WATER_BASED','MILK_BASED','PLANT_BASED','JUICE_BASED','ALCOHOL_BASED','FERMENTED_BASED'));
			END IF;
		END $$;`,
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'chk_beverage_nutrition_non_negative') THEN
				ALTER TABLE beverage_nutrition
				ADD CONSTRAINT chk_beverage_nutrition_non_negative
				CHECK (sugar_g_100ml >= 0 AND caffeine_mg_100ml >= 0 AND sodium_mg >= 0 AND alcohol_percent >= 0);
			END IF;
		END $$;`,
	}

	for _, stmt := range stmts {
		if err := db.Exec(stmt).Error; err != nil {
			return err
		}
	}
	return nil
}
