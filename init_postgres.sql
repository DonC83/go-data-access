-- CREATE USER go_user WITH PASSWORD 'example';
-- CREATE DATABASE recordings;
-- GRANT ALL PRIVILEGES ON DATABASE recordings TO go_user;

DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         serial,
  title      varchar(128) NOT NULL ,
  artist     varchar(255) NOT NULL ,
  price      numeric CHECK (price > 0),
  PRIMARY KEY (id)
);

INSERT INTO album
(title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);