CREATE TABLE "products" (
  "id" VARCHAR PRIMARY KEY,
  "title" VARCHAR,
  "description" VARCHAR,
  "price" FLOAT,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "products_categories" (
  "id" VARCHAR PRIMARY KEY,
  "product_id" VARCHAR,
  "category_id" INT
);

CREATE TABLE "categories" (
  "id" INT PRIMARY KEY,
  "title" VARCHAR UNIQUE
);

CREATE TABLE "users" (
  "id" VARCHAR PRIMARY KEY,
  "username" VARCHAR UNIQUE,
  "password" VARCHAR,
  "email" VARCHAR UNIQUE,
  "role_id" INT,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "oauth" (
  "id" VARCHAR PRIMARY KEY,
  "user_id" VARCHAR,
  "access_token" VARCHAR,
  "refresh_token" VARCHAR,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "roles" (
  "id" INTREGER PRIMARY KEY,
  "title" VARCHAR
);

CREATE TABLE "images" (
  "id" VARCHAR PRIMARY KEY,
  "filename" VARCHAR,
  "url" VARCHAR,
  "product_id" VARCHAR,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" VARCHAR PRIMARY KEY,
  "user_id" VARCHAR,
  "contact" VARCHAR,
  "address" VARCHAR,
  "tranfer_slip" JSONB,
  "status" VARCHAR,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "products_orders" (
  "id" VARCHAR PRIMARY KEY,
  "order_id" VARCHAR,
  "qty" INT,
  "product_id" JSONB
);

ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
