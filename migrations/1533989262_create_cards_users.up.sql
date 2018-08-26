CREATE TABLE cards_users (
  id SERIAL PRIMARY KEY,
  card_id INTEGER REFERENCES cards(id) ON DELETE CASCADE ON UPDATE CASCADE,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  correct BOOLEAN DEFAULT FALSE NOT NULL
);