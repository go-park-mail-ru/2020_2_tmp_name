CREATE TABLE "chat" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,
  "filter_id" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (filter_id) REFERENCES filters(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "comments" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,
  "time_delivery" text,
  "text" text,

  FOREIGN KEY (user_id1) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "dislikes" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,
  "filter_id" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (filter_id) REFERENCES filters(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "filter" (
  "id" SERIAL PRIMARY KEY,
  "target" varchar
);

CREATE TABLE "likes" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,
  "filter_id" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (filter_id) REFERENCES filters(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "message" (
  "id" SERIAL PRIMARY KEY,
  "text" varchar NOT NULL,
  "time_delivery" text,
  "chat_id" int4,
  "user_id" int4,

  FOREIGN KEY (chat_id) REFERENCES chat(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "photo" (
  "id" SERIAL PRIMARY KEY,
  "path" varchar NOT NULL,
  "user_id" int4,
  "mask" varchar  DEFAULT '',

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "premium_accounts" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int4,
  "date_from" date,
  "date_to" date,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "sessions" (
  "id" SERIAL PRIMARY KEY,
  "key" varchar NOT NULL,
  "value" varchar
);

CREATE TABLE "superlikes" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,
  "filter_id" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (filter_id) REFERENCES filters(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "telephone" varchar NOT NULL,
  "password" varchar NOT NULL,
  "date_birth" int4,
  "sex" varchar NOT NULL,
  "education" varchar,
  "job" varchar,
  "about_me" varchar,
  "filter_id" int4,

  FOREIGN KEY (filter_id) REFERENCES filters(id) ON DELETE CASCADE ON UPDATE CASCADE
);
