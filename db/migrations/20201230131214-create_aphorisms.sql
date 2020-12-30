-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS aphorisms_id_sequence;

CREATE TABLE IF NOT EXISTS aphorisms (
  id integer DEFAULT nextval('aphorisms_id_sequence') PRIMARY KEY,
  phrase text NOT NULL,
  language_code varchar(10) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (language_code)
);

-- +migrate Down
DROP TABLE IF EXISTS aphorisms;

DROP SEQUENCE IF EXISTS aphorisms_id_sequence;
