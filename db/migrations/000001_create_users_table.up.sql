CREATE TABLE customers(
  customer_id  serial PRIMARY KEY,
  name VARCHAR (255) NOT NULL,
  city VARCHAR (255) NOT NULL,
  zipcode INT,
  date_of_birth DATE NOT NULL,
  status  BOOLEAN DEFAULT TRUE,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE
)