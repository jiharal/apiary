-- Postgresql
CREATE TABLE "user" (
  id             SERIAL         PRIMARY KEY,
  name           TEXT           NOT NULL,
  email          TEXT           NOT NULL,
  password       TEXT           NOT NULL,
  created_at     TIMESTAMPTZ    NOT NULL DEFAULT now(),
  created_by     SERIAL         NOT NULL,
  updated_at     TIMESTAMPTZ,
  updated_by     SERIAL
);

CREATE TABLE product (
  id             SERIAL         PRIMARY KEY,
  name           TEXT           NOT NULL,
  price          SERIAL         NOT NULL,
  imageurl       TEXT           NOT NULL,
  created_at     TIMESTAMPTZ    NOT NULL DEFAULT now(),
  created_by     SERIAL         NOT NULL,
  updated_at     TIMESTAMPTZ,
  updated_by     SERIAL
);