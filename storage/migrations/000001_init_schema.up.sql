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
  "group_id" int4
);

CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "for_love" boolean,
  "for_friends" boolean,
  "for_sex" boolean
);

CREATE TABLE "likes" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id)
);

CREATE TABLE "dislikes" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id)
);

CREATE TABLE "chat" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,

  FOREIGN KEY (user_id1) REFERENCES users(id)
);

CREATE TABLE "message" (
  "id" SERIAL PRIMARY KEY,
  "text" varchar NOT NULL,
  "time_delivery" text,
  "chat_id" int4,
  "user_id" int4,

  FOREIGN KEY (chat_id) REFERENCES chat(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE "photo" (
  "id" SERIAL PRIMARY KEY,
  "path" varchar NOT NULL,
  "user_id" int4, 

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE "comments" (
  "id" SERIAL PRIMARY KEY,
  "user_id1" int4,
  "user_id2" int4,
  "time_delivery" text,
  "text" text,

  FOREIGN KEY (user_id1) REFERENCES users(id)
);

CREATE TABLE "sessions" (
  "id" SERIAL PRIMARY KEY,
  "key" varchar NOT NULL,
  "value" varchar
);