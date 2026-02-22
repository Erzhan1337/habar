package internal

const (
	QueryMarketStats = `
		SELECT 
			ROUND(100.0 * COUNT(DISTINCT CASE WHEN f.flag_type IN ('HIGH_SUGAR', 'VERY_HIGH_SUGAR') THEN m.product_id END) / NULLIF(COUNT(DISTINCT m.product_id), 0), 2) as sugar_pct,
			ROUND(100.0 * COUNT(DISTINCT CASE WHEN f.flag_type = 'CONTAINS_ARTIFICIAL_SWEETENER' THEN m.product_id END) / NULLIF(COUNT(DISTINCT m.product_id), 0), 2) as sweetener_pct,
			ROUND(100.0 * COUNT(DISTINCT CASE WHEN f.flag_type = 'CONTAINS_CAFFEINE' THEN m.product_id END) / NULLIF(COUNT(DISTINCT m.product_id), 0), 2) as caffeine_pct,
			ROUND(100.0 * COUNT(DISTINCT CASE WHEN f.flag_type = 'CONTAINS_MULTIPLE_ADDITIVES' THEN m.product_id END) / NULLIF(COUNT(DISTINCT m.product_id), 0), 2) as heavy_additive_pct
		FROM beverage_master m
		LEFT JOIN beverage_flags f ON m.product_id = f.product_id;`

	QueryCategoryDistribution = `
		SELECT category_lvl_1, COUNT(*) as beverage_count
		FROM beverage_master
		GROUP BY category_lvl_1
		ORDER BY beverage_count DESC;`

	QueryAdditivesHeatmap = `
		SELECT m.category_lvl_1, COALESCE(i.additive_category, 'UNKNOWN') AS additive_category, COUNT(*) AS usage_count
		FROM beverage_master m
		JOIN beverage_ingredients i ON m.product_id = i.product_id
		WHERE i.e_number IS NOT NULL
		GROUP BY m.category_lvl_1, COALESCE(i.additive_category, 'UNKNOWN')
		ORDER BY m.category_lvl_1, usage_count DESC;`

	QueryETopChart = `
		SELECT e_number, ingredient_name, COUNT(*) as usage_count
		FROM beverage_ingredients
		WHERE e_number IS NOT NULL
		GROUP BY e_number, ingredient_name
		ORDER BY usage_count DESC
		LIMIT 10;`

	QuerySweetenerMatrix = `
		SELECT m.category_lvl_1, i.sweetener, COUNT(DISTINCT m.product_id) as beverage_count
		FROM beverage_master m
		JOIN beverage_ingredients i ON m.product_id = i.product_id
		WHERE i.sweetener IS NOT NULL
		GROUP BY m.category_lvl_1, i.sweetener
		ORDER BY m.category_lvl_1, beverage_count DESC;`

	QuerySugarDistribution = `
		SELECT
			CASE
				WHEN n.sugar_g_100ml < 0.5 THEN 'ZERO_OR_TRACE'
				WHEN n.sugar_g_100ml <= 5 THEN 'LOW'
				WHEN n.sugar_g_100ml <= 10 THEN 'MEDIUM'
				ELSE 'HIGH'
			END AS sugar_segment,
			COUNT(*) AS beverage_count
		FROM beverage_master m
		JOIN beverage_nutrition n ON m.product_id = n.product_id
		GROUP BY sugar_segment
		ORDER BY beverage_count DESC;`

	QueryCleanLabelShare = `
		SELECT
			ROUND(100.0 * COUNT(CASE WHEN NOT EXISTS (
				SELECT 1 FROM beverage_flags f
				WHERE f.product_id = m.product_id
				AND f.flag_type IN ('ULTRA_PROCESSED', 'CONTAINS_SYNTHETIC_COLOR', 'CONTAINS_ARTIFICIAL_SWEETENER')
			) THEN 1 END) / NULLIF(COUNT(*), 0), 2) AS clean_label_pct
		FROM beverage_master m;`

	QueryCleanLabel = `
		SELECT * FROM beverage_master m
		WHERE NOT EXISTS (
			SELECT 1 FROM beverage_flags f 
			WHERE f.product_id = m.product_id 
			AND f.flag_type IN ('ULTRA_PROCESSED', 'CONTAINS_SYNTHETIC_COLOR', 'CONTAINS_ARTIFICIAL_SWEETENER')
		);`

	QueryCompositionClassification = `
		SELECT
			m.product_id,
			m.name,
			m.category_lvl_1,
			m.base_type,
			m.caffeine_type,
			m.carbonation,
			COALESCE(string_agg(DISTINCT i.sweetener, ','), '') AS sweeteners,
			COALESCE(string_agg(DISTINCT i.additive_category, ','), '') AS additive_categories,
			COALESCE(string_agg(DISTINCT i.allergen, ','), '') AS allergens
		FROM beverage_master m
		LEFT JOIN beverage_ingredients i ON m.product_id = i.product_id
		GROUP BY m.product_id, m.name, m.category_lvl_1, m.base_type, m.caffeine_type, m.carbonation
		ORDER BY m.product_id;`
)
