
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE SEQUENCE IF NOT EXISTS aphorisms_id_sequence;

CREATE TABLE IF NOT EXISTS aphorisms (
  id integer DEFAULT nextval('aphorisms_id_sequence') PRIMARY KEY,
  phrase text NOT NULL,
  language_code varchar(10) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (language_code)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS aphorisms;

DROP SEQUENCE IF EXISTS aphorisms_id_sequence;
