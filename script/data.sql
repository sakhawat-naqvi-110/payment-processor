CREATE TABLE merchant (
   id SERIAL PRIMARY KEY,
   merchant_name VARCHAR(255) NOT NULL,
   merchant_code VARCHAR(50) UNIQUE NOT NULL,
   allowed_currency VARCHAR(10) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   created_by VARCHAR(255),
   last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   last_updated_by VARCHAR(255),
   is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE customer (
   id SERIAL PRIMARY KEY,
   customer_name VARCHAR(255) NOT NULL,
   customer_email VARCHAR(255) UNIQUE NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   created_by VARCHAR(255),
   last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   last_updated_by VARCHAR(255),
   is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE invoice (
  id SERIAL PRIMARY KEY,
  merchant_id INT NOT NULL REFERENCES merchant(id) ON DELETE CASCADE,
  customer_id INT NOT NULL REFERENCES customer(id) ON DELETE CASCADE,
  amount NUMERIC(18,2) NOT NULL CHECK (amount > 0),
  currency VARCHAR(10) NOT NULL,
  optional_description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_updated_by VARCHAR(255),
  is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE payment (
  id SERIAL PRIMARY KEY,
  invoice_id INT NOT NULL REFERENCES invoice(id) ON DELETE CASCADE,
  merchant_id INT NOT NULL REFERENCES merchant(id) ON DELETE CASCADE,
  customer_id INT NOT NULL REFERENCES customer(id) ON DELETE CASCADE,
  amount NUMERIC(18,2) NOT NULL CHECK (amount > 0),
  payment_status VARCHAR(50) NOT NULL,
  payment_method VARCHAR(100) NOT NULL,
  payment_source VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_updated_by VARCHAR(255),
  is_active BOOLEAN DEFAULT TRUE
);

-- inserts for validation purposes
-- Insert sample merchant
INSERT INTO merchant (merchant_name, merchant_code, allowed_currency, created_by, last_updated_by)
VALUES
    ('Amazon', 'AMZ123', ARRAY['USD', 'EUR'], 'admin', 'admin'),
    ('eBay', 'EBY456', ARRAY['USD', 'GBP'], 'admin', 'admin');

-- Insert sample customer
INSERT INTO customer (customer_name, customer_email, created_by, last_updated_by)
VALUES
    ('John Doe', 'john.doe@example.com', 'admin', 'admin'),
    ('Jane Smith', 'jane.smith@example.com', 'admin', 'admin');

