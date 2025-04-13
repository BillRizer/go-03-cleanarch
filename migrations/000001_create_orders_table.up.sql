CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		price float NOT NULL,
		tax float NOT NULL,
		final_price float NOT NULL
)