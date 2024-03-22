
CREATE TABLE users (
   user_id      serial PRIMARY KEY,
   username     VARCHAR (50) UNIQUE,
   email        TEXT UNIQUE,
   password     TEXT,
   img          TEXT,
   created_at   TIMESTAMPTZ,
   updated_at   TIMESTAMPTZ
);
CREATE TABLE article (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    img         TEXT,
    owner       INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE msguser (
    id          SERIAL PRIMARY KEY,
    coming      TEXT,
    img         TEXT,
    owner       INTEGER,
    to_user     INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id),
    FOREIGN KEY (to_user) REFERENCES users(user_id)
);
CREATE TABLE groups (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    img         TEXT,
    owner       INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE msggroups (
    id          SERIAL PRIMARY KEY,
    coming      TEXT,
    img         TEXT,
    owner       INTEGER,
    to_group    INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id),
    FOREIGN KEY (to_group) REFERENCES groups(id) ON DELETE CASCADE
);
CREATE TABLE subscription (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    owner       INTEGER,
    to_user     INTEGER,
    to_group    INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_user) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_group) REFERENCES groups(id) ON DELETE CASCADE
);
