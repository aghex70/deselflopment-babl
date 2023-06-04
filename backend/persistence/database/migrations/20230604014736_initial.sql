-- +goose Up
CREATE TABLE deselflopment-babl_calendars (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(128),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

CREATE TABLE deselflopment-babl_entries (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(128),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

CREATE TABLE deselflopment-babl_users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(128),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

