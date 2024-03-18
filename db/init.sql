CREATE TABLE invoices (
	id SERIAL PRIMARY KEY NOT NULL, 
	name VARCHAR(100),
	position INT DEFAULT 0,
	archived BOOLEAN,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	total         FLOAT(2),
	SubTotal      FLOAT(2),
	VatPercentage FLOAT(2),
	Description   VARCHAR(100)
);
