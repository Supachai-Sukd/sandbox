
CREATE TABLE
    IF NOT EXISTS "true_money_wallet" (
        id SERIAL PRIMARY KEY,
        name TEXT,
        category TEXT,
        currency TEXT,
        balance float8 NOT NULL DEFAULT 0
    );

INSERT INTO "true_money_wallet" ("id", "name", "category", "currency", "balance") VALUES (1, 'test-name', 'test-category', 'test-currency', 9.1);