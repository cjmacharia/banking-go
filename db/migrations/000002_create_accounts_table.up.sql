CREATE TABLE accounts(
  account_id serial PRIMARY KEY,
  customer_id  VARCHAR (255) NOT NULL,
  account_type  VARCHAR (255) NOT NULL,
   amount VARCHAR (255) NOT NULL,
  opening_date DATE NOT NULL,
  status  BOOLEAN DEFAULT TRUE,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE
)