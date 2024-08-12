CREATE TABLE "colors"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT 'now()',
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT 'now()',
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "value" TEXT NULL,
        "product_id" BIGINT NULL
);
CREATE TABLE "sellers"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "user_id" BIGINT NULL,
        "name" TEXT NULL,
        "image" TEXT NULL,
        "phone" TEXT NULL,
        "description" TEXT NULL
);
ALTER TABLE
    "sellers" ADD PRIMARY KEY("id");
CREATE INDEX "sellers_deleted_at_index" ON
    "sellers"("deleted_at");
CREATE TABLE "sizes"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT 'now()',
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT 'now()',
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "value" TEXT NULL,
        "product_id" BIGINT NULL
);
CREATE TABLE "checkouts"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT 'now()',
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT 'now()',
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "cart_id" BIGINT NULL,
        "user_id" BIGINT NULL,
        "delivery" TEXT NOT NULL,
        "summary" DECIMAL(10, 2) NOT NULL
);
ALTER TABLE
    "checkouts" ADD PRIMARY KEY("id");
CREATE TABLE "categories"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "name" TEXT NULL,
        "image" TEXT NULL,
        "slug" TEXT NULL
);
ALTER TABLE
    "categories" ADD PRIMARY KEY("id");
CREATE INDEX "categories_deleted_at_index" ON
    "categories"("deleted_at");
CREATE TABLE "products"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "category_id" BIGINT NULL,
        "name" TEXT NULL,
        "seller_id" BIGINT NULL,
        "price" DECIMAL(10, 2) NULL,
        "stock" BIGINT NULL,
        "image" TEXT NULL,
        "size" BIGINT NULL,
        "color" TEXT NULL,
        "rating" BIGINT NULL
);
CREATE INDEX "products_deleted_at_index" ON
    "products"("deleted_at");
CREATE TABLE "images"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT 'now()',
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT 'now()',
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "url" TEXT NULL,
        "product_id" BIGINT NULL
);
CREATE TABLE "users"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "email" TEXT NULL,
        "password" TEXT NULL,
        "role" TEXT NULL,
        "verify" VARCHAR(255) NULL
);
CREATE INDEX "users_deleted_at_index" ON
    "users"("deleted_at");
CREATE TABLE "orders"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT 'now()',
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT 'now()',
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "user_id" BIGINT NULL,
        "checkout_id" BIGINT NULL,
        "address_id" BIGINT NULL,
        "status" TEXT NULL,
        "total_price" DECIMAL(10, 2) NULL
);
ALTER TABLE
    "orders" ADD PRIMARY KEY("id");
CREATE TABLE "carts"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "product_id" BIGINT NULL,
        "seller_id" BIGINT NULL,
        "user_id" BIGINT NULL,
        "quantity" BIGINT NOT NULL
);
CREATE INDEX "carts_deleted_at_index" ON
    "carts"("deleted_at");
CREATE TABLE "user_verifications"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "user_id" BIGINT NULL,
        "token" TEXT NULL
);
CREATE INDEX "user_verifications_deleted_at_index" ON
    "user_verifications"("deleted_at");
CREATE TABLE "addresses"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT 'now()',
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT 'now()',
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "user_id" BIGINT NULL,
        "main_address" TEXT NULL,
        "name" TEXT NOT NULL,
        "phone" TEXT NOT NULL,
        "detail_address" TEXT NOT NULL,
        "postal_code" TEXT NOT NULL,
        "city" TEXT NOT NULL,
        "primary" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "addresses" ADD PRIMARY KEY("id");
CREATE TABLE "customers"(
    "id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(0)
    WITH
        TIME zone NULL DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "user_id" BIGINT NULL,
        "name" TEXT NULL,
        "image" TEXT NULL,
        "phone" TEXT NULL,
        "date_of_birth" TIMESTAMP(0)
    WITH
        TIME zone NULL,
        "gender" TEXT NULL DEFAULT '' other ''
);
ALTER TABLE
    "customers" ADD PRIMARY KEY("id");
CREATE INDEX "customers_deleted_at_index" ON
    "customers"("deleted_at");
ALTER TABLE
    "carts" ADD CONSTRAINT "carts_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "products"("id");
ALTER TABLE
    "carts" ADD CONSTRAINT "carts_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_checkout_id_foreign" FOREIGN KEY("checkout_id") REFERENCES "checkouts"("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_id_foreign" FOREIGN KEY("id") REFERENCES "sellers"("user_id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_address_id_foreign" FOREIGN KEY("address_id") REFERENCES "addresses"("id");
ALTER TABLE
    "products" ADD CONSTRAINT "products_seller_id_foreign" FOREIGN KEY("seller_id") REFERENCES "sellers"("id");
ALTER TABLE
    "checkouts" ADD CONSTRAINT "checkouts_cart_id_foreign" FOREIGN KEY("cart_id") REFERENCES "carts"("id");
ALTER TABLE
    "products" ADD CONSTRAINT "products_category_id_foreign" FOREIGN KEY("category_id") REFERENCES "categories"("id");
ALTER TABLE
    "checkouts" ADD CONSTRAINT "checkouts_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "colors" ADD CONSTRAINT "colors_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "products"("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "customers" ADD CONSTRAINT "customers_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "sizes" ADD CONSTRAINT "sizes_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "products"("id");
ALTER TABLE
    "user_verifications" ADD CONSTRAINT "user_verifications_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "images" ADD CONSTRAINT "images_id_foreign" FOREIGN KEY("id") REFERENCES "products"("id");
ALTER TABLE
    "carts" ADD CONSTRAINT "carts_seller_id_foreign" FOREIGN KEY("seller_id") REFERENCES "sellers"("id");
ALTER TABLE
    "addresses" ADD CONSTRAINT "addresses_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");