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
    popularity  INTEGER             NOT NULL,
    index latitude_longitude_popularity(latitude, longitude, popularity),
    index door_height_door_width_rent_popularity(door_height, door_width, rent, popularity),
    index door_height_rent_popularity(door_height, rent, popularity),
    index door_width_rent_popularity(door_width, rent, popularity),
    INDEX rent (rent),
    INDEX popularity (popularity),
    INDEX rent_popularity (rent, popularity)
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
    stock       INTEGER         NOT NULL,
    index kind_stock_popularity (kind, stock, popularity),
    index color_stock_popularity (color, stock, popularity),
    index height_stock_popularity (height, stock, popularity),
    index width_stock_popularity (width, stock, popularity),
    INDEX price_stock_popularity (price, stock, popularity),
    INDEX price (price),
    INDEX stock_price (stock, price)
);
