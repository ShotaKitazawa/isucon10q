DROP DATABASE IF EXISTS isuumo;
CREATE DATABASE isuumo;

DROP TABLE IF EXISTS isuumo.estate;
DROP TABLE IF EXISTS isuumo.chair;

CREATE TABLE isuumo.estate
(
    id          INTEGER             NOT NULL PRIMARY KEY,
    name        VARCHAR(64)         NOT NULL,
    description VARCHAR(4096)       NOT NULL,
    thumbnail   VARCHAR(128)        NOT NULL,
    address     VARCHAR(128)        NOT NULL,
    latitude    DOUBLE PRECISION    NOT NULL,
    longitude   DOUBLE PRECISION    NOT NULL,
    rent        INTEGER             NOT NULL,
    door_height INTEGER             NOT NULL,
    door_width  INTEGER             NOT NULL,
    features    VARCHAR(64)         NOT NULL,
    popularity  INTEGER             NOT NULL
);

CREATE TABLE isuumo.chair
(
    id          INTEGER         NOT NULL PRIMARY KEY,
    name        VARCHAR(64)     NOT NULL,
    description VARCHAR(4096)   NOT NULL,
    thumbnail   VARCHAR(128)    NOT NULL,
    price       INTEGER         NOT NULL,
    height      INTEGER         NOT NULL,
    width       INTEGER         NOT NULL,
    depth       INTEGER         NOT NULL,
    color       VARCHAR(64)     NOT NULL,
    features    VARCHAR(64)     NOT NULL,
    kind        VARCHAR(64)     NOT NULL,
    popularity  INTEGER         NOT NULL,
    stock       INTEGER         NOT NULL
);

ALTER TABLE isuumo.chair add index price_id_index(price,id);
ALTER TABLE isuumo.estate add index price_id_index(rent,id);

ALTER TABLE isuumo.chair add index price_height_width(price,height,width);
ALTER TABLE isuumo.chair add index height_width(height,width);
ALTER TABLE isuumo.chair add index price_width(price,width);
ALTER TABLE isuumo.chair add index price_height(price,height);
ALTER TABLE isuumo.chair add index price(price);
ALTER TABLE isuumo.chair add index height(height);
ALTER TABLE isuumo.chair add index width(width);

ALTER TABLE isuumo.estate add index height_width_rent(door_height,door_width,rent);
ALTER TABLE isuumo.estate add index width_rent(door_width,rent);
ALTER TABLE isuumo.estate add index height_rent(door_height,rent);
ALTER TABLE isuumo.estate add index height_width(door_height,door_width);
ALTER TABLE isuumo.estate add index height(door_height);
ALTER TABLE isuumo.estate add index width(door_width);
ALTER TABLE isuumo.estate add index rent(rent);
