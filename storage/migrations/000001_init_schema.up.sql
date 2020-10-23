CREATE TABLE "users" (
  "id" int4 PRIMARY KEY,
  "name" varchar NOT NULL,
  "telephone" varchar NOT NULL,
  "password" varchar NOT NULL,
  "date_birth" date NOT NULL DEFAULT (now()),
  "sex" varchar NOT NULL,
  "studing" varchar,
  "working" varchar,
  "about_me" varchar,
  "group_id" int4 NOT NULL
);

CREATE TABLE "groups" (
  "id" int4 PRIMARY KEY,
  "for_love" boolean,
  "for_friends" boolean,
  "for_sex" boolean
);

CREATE TABLE "likes" (
  "id" int4 PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4
);

CREATE TABLE "chat" (
  "id" int4 PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4
);

CREATE TABLE "message" (
  "id" int4 PRIMARY KEY,
  "path" varchar NOT NULL,
  "time_delivery" timestamptz NOT NULL DEFAULT (now()),
  "chat_id" int4,
  "user_id" int4
);

CREATE TABLE "photo" (
  "id" int4 PRIMARY KEY,
  "path" varchar NOT NULL,
  "user_id" int4
);
