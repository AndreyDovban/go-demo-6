-- Получение ссылок
-- SELECT *
-- FROM "links"
-- WHERE deleted_at is NULL
-- ORDER BY id ASC
-- LIMIT 5
-- OFFSET 0

-- Получение колличества активных записей
-- SELECT count(*)
-- FROM "links"
-- WHERE deleted_at is NULL


-- SELECT to_char(date, 'YYYY-MM-DD') as period, sum(clicks)
-- FROM "stats"
-- WHERE date BETWEEN '01/01/2024' AND '01/01/2025'
-- GROUP BY period 
-- ORDER BY period