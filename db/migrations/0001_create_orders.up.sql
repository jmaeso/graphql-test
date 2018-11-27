CREATE TABLE IF NOT EXISTS orders (
    id UUID NOT NULL,
    retailer_id UUID NOT NULL,
    num_packages INT NOT NULL,
    PRIMARY KEY (id)
);
