DROP TABLE IF EXISTS record CASCADE;
DROP TABLE IF EXISTS story CASCADE;
DROP TABLE IF EXISTS source CASCADE;
DROP TABLE IF EXISTS pref CASCADE;
DROP TABLE IF EXISTS run CASCADE;
DROP TABLE IF EXISTS search CASCADE;
DROP TABLE IF EXISTS account CASCADE;
DROP TABLE IF EXISTS auth CASCADE;

CREATE TABLE IF NOT EXISTS source (
    name varchar NOT NULL PRIMARY KEY,
    state JSONB
);

CREATE TABLE IF NOT EXISTS story (
    id varchar NOT NULL,
    source varchar NOT NULL,
    title varchar NOT NULL,
    author varchar NOT NULL,
    profile varchar,
    created timestamp NOT NULL,
    url varchar,
    category varchar NOT NULL,
    tokens varchar[] NOT NULL,
    magnitude decimal NOT NULL,
    score decimal NOT NULL,
    tags JSONB,
    PRIMARY KEY (id, source),
    CONSTRAINT story_source_fk FOREIGN KEY(source) REFERENCES source(name) ON DELETE CASCADE
);

CREATE INDEX ON story (id,source,lower(title),lower(author),category,tokens,created,magnitude,score);


CREATE TABLE IF NOT EXISTS record (
    story_id varchar NOT NULL,
    source varchar NOT NULL,
    created timestamp NOT NULL,
    PRIMARY KEY (story_id, source, created),
    CONSTRAINT record_story_fk FOREIGN KEY(story_id,source) REFERENCES story(id,source) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS account (
    email varchar NOT NULL PRIMARY KEY,
    picture varchar NOT NULL,
    org varchar NOT NULL,
    updated timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS auth (
    id varchar NOT NULL PRIMARY KEY,
    created timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS pref (
    email varchar NOT NULL PRIMARY KEY,
    content JSONB,
    CONSTRAINT pref_account_fk FOREIGN KEY(email) REFERENCES account(email) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS search (
    email varchar NOT NULL,
    query_name varchar NOT NULL,
    content JSONB,
    PRIMARY KEY (email, query_name),
    CONSTRAINT search_account_fk FOREIGN KEY(email) REFERENCES account(email) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS run (
    email varchar NOT NULL,
    query_name varchar NOT NULL,
    last_run timestamp NOT NULL,
    PRIMARY KEY (email, query_name, last_run),
    CONSTRAINT run_account_fk FOREIGN KEY(email) REFERENCES account(email) ON DELETE CASCADE,
    CONSTRAINT run_search_fk FOREIGN KEY(email,query_name) REFERENCES search(email,query_name) ON DELETE CASCADE
);


ALTER TABLE source OWNER TO dev;
ALTER TABLE story OWNER TO dev;
ALTER TABLE record OWNER TO dev;
ALTER TABLE account OWNER TO dev;
ALTER TABLE auth OWNER TO dev;
ALTER TABLE pref OWNER TO dev;
ALTER TABLE search OWNER TO dev;
ALTER TABLE run OWNER TO dev;


GRANT CONNECT ON DATABASE news TO dev;
GRANT ALL PRIVILEGES ON DATABASE news TO dev;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dev;