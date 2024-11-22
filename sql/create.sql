CREATE TABLE IF NOT EXISTS user_balance
(
    id                   BIGINT PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    email                VARCHAR(64)        NOT NULL, --user email
    chain                VARCHAR(16)        NOT NULL, -- chain name:ethereum,base,polygon
    coin_type            VARCHAR(16)        NOT NULL, -- coin type:eth,usdt,uni
    balance              DECIMAL(65, 18)    NOT NULL  --user balance
);
CREATE UNIQUE INDEX IF NOT EXISTS user_balance_unique_idx_user_balance ON user_balance (email,chain,coin_type);
CREATE INDEX IF NOT EXISTS user_balance_idx_email ON user_balance (email);

CREATE TABLE IF NOT EXISTS transaction_history
(
    id                   BIGINT PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    "from"               VARCHAR(64)        NOT NULL,
    "to"                 VARCHAR(64)        NOT NULL,
    value                DECIMAL(65, 18)    NOT NULL,
    chain                VARCHAR(16)        NOT NULL, -- chain name:eth,base,op
    coin_type            VARCHAR(16)        NOT NULL, -- coin type:eth,usdt,uni
    timestamp            BIGINT             NOT NULL
);
CREATE INDEX IF NOT EXISTS transaction_history_idx_from ON transaction_history ("from");
CREATE INDEX IF NOT EXISTS transaction_history_idx_to ON transaction_history ("to");

CREATE TABLE IF NOT EXISTS user_info
(
    id                   BIGINT PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    email                VARCHAR(64)        NOT NULL, --user email
    chain                VARCHAR(16)        NOT NULL,
    address              VARCHAR(64)        NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS user_info_unique_idx_user_info ON user_info (email,chain,address);
CREATE INDEX IF NOT EXISTS user_info_idx_user_info ON user_info (email,chain);

INSERT INTO user_info (id, email, chain, address) VALUES (1, 'a@gmail.com', 'ethereum', '0x70997970C51812dc3A010C7d01b50e0d17dc79C8');
INSERT INTO user_info (id, email, chain, address) VALUES (2, 'b@gmail.com', 'ethereum', '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC');
INSERT INTO user_info (id, email, chain, address) VALUES (3, 'c@gmail.com', 'ethereum', '0x90F79bf6EB2c4f870365E785982E1f101E93b906');

INSERT INTO user_balance (id, email, chain, coin_type, balance) VALUES (1, 'a@gmail.com', 'ethereum', 'eth', 2);
INSERT INTO user_balance (id, email, chain, coin_type, balance) VALUES (2, 'b@gmail.com', 'ethereum', 'eth', 3);
INSERT INTO user_balance (id, email, chain, coin_type, balance) VALUES (3, 'c@gmail.com', 'ethereum', 'eth', 4);