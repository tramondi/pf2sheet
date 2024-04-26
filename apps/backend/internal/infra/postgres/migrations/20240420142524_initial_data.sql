-- migrate:up
INSERT INTO "classes"("code", "title") VALUES
  ('bard', 'Бард'),
  ('fighter', 'Воин'),
  ('wizard', 'Волшебник'),
  ('druid', 'Друид'),
  ('cleric', 'Жрец'),
  ('monk', 'Монах'),
  ('rouge', 'Плут');

INSERT INTO "ancestries"("code", "title") VALUES
  ('orc', 'Орк'),
  ('gnome', 'Гном'),
  ('human', 'Человек'),
  ('elf', 'Эльф'),
  ('halfling', 'Халфлинг');

-- migrate:down

