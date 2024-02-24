# Simple GoLang REST API with Gin Framework and GORM

This repository contains a simple GoLang REST API built using the Gin web framework and GORM package for interacting with the database. The API provides endpoints to retrieve JSON data for products.

## Requirements

- Go (v1.16 or higher)
- Gin Framework
- GORM
- MySQL (or any other supported database)

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/mfhmiii/go-rest-api-1.git
```

2. Depedency installation
```bash
go mod tidy
```

3. Set up your database connection in config/config.go.

4. Run the application:
```bash
go run main.go
```

## Usage
Once the application is running, you can access the following endpoints:

**GET /products**: Retrieve JSON data for all products.<br>
**GET /products/:id**: Retrieve JSON data for a specific product by ID.<br>
**POST /product**: Create a new product. Send JSON data in the request body with the product details.<br>
**PUT /products/:id**: Update an existing product by ID. Send JSON data in the request body with the updated product details.<br>
**DELETE /product**: Delete a product by ID.<br>

## Example

Get All Products
```bash
curl http://localhost:8080/products
```

Get Product by ID
```bash
curl http://localhost:8080/products/1
```

Create Product
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"New Product","price":9.99}' http://localhost:8080/products
```

Update Product
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Product","price":19.99}' http://localhost:8080/products/1
```
Delete Product
```bash
curl -X DELETE http://localhost:8080/product
```

## Contributing
Contributions are welcome! Please feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License.
```csharp
This Markdown content will render as a structured and formatted README when viewed in a Markdown renderer, such as on GitHub.
```
