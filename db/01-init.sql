
CREATE TABLE
    IF NOT EXISTS "true_money_wallet" (
        id SERIAL PRIMARY KEY,
        name TEXT,
        category TEXT,
        currency TEXT,
        balance float8 NOT NULL DEFAULT 0
    );

