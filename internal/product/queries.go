package product

var (
	selectAllProducts = "SELECT * FROM my_db.products;"
	selectProductById = "SELECT * FROM my_db.products WHERE id = ?;"
	insertProduct     = `INSERT INTO my_db.products(name, quantity, code_value, is_published, expiration, price) 
		VALUES(?, ?, ?, ?, ?, ?);`
	deleteProduct = `DELETE FROM my_db.products WHERE id = ?;`
	updateProduct = `UPDATE my_db.products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?`
)
