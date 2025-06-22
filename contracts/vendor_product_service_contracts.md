# Vendor Products Service API Contracts

---

__Swagger API Documentation: https://marketplace-vendor-products-service.onrender.com/swagger/index.html__

## 1. Get All Products Information

__GET ```api/products```__

__Query parameters:__

| Parameter | Type    | Description                                                      |
|-----------|---------|------------------------------------------------------------------|
| category  | string  | Filter products by category                                      |
| minPrice  | float64 | Filter products by price greater than or equal to this value     |
| maxPrice  | float64 | Filter products by price less than or equal to this value        |
| search    | string  | Search products by name                                          |
| sortBy    | string  | Sort products by fields: price, name, quantity                   |
| sortOrder | string  | Sort order: asc/desc                                             |
| limit     | int     | Max amount of products to return                                 |
| offset    | int     | Number of products to skip before starting collecting the result |
| page      | int     | Page number to retrieve                                          |
| size      | int     | Number of products to return per page                            |

__Response:__

```json
[
  {
    "id": "uuid",
    "vendor_id": "string",
    "name": "string",
    "price": "float64",
    "category": "string",
    "image": {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    },
    "quantity": "int"
  }
]
```

__Status codes:__

- ```200 OK (success)```
- ```500 Internal Server Error```

---

## 2. Get One Product Information

__GET ```api/products/:productId```__

__Response:__

```json
{
  "id": "uuid",
  "vendor_id": "string",
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": [
    {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    }
  ],
  "tags": [
    {
      "id": "uuid",
      "tag_name": "string"
    }
  ],
  "quantity": "int"
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid productId)```
- ```404 Not Found (product not found)```
- ```500 Internal Server Error```

---

## 3. Add New Product

__POST ```api/products```__

__Body:__

```json
{
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": [
    "string"
  ],
  "tags": [
    "string"
  ]
}
```

__Response:__

```json
{
  "id": "uuid",
  "vendor_id": "string",
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": [
    {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    }
  ],
  "tags": [
    {
      "id": "uuid",
      "tag_name": "string"
    }
  ],
  "quantity": "int"
}
```

__Status codes:__

- ```201 Created (success)```
- ```400 Bad Request (invalid product data)```
- ```500 Internal Server Error```

---

## 4. Update Product Information

__PUT ```api/products/:productId```__

__Body:__

```json
{
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": [
    {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    }
  ],
  "tags": [
    {
      "id": "uuid",
      "tag_name": "string"
    }
  ]
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid productId/product data)```
- ```404 Not Found (product not found)```
- ```500 Internal Server Error```

---

## 5. Modify Product Information

__PATCH ```api/products/:productId```__

__Body:__

```json
{
  "name": "string",
  (optional)
  "description": "string",
  (optional)
  "price": "float64",
  (optional)
  "category": "string",
  (optional)
  "images": [
    {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    }
  ],
  (optional)
  "tags": [
    {
      "id": "uuid",
      "tag_name": "string"
    }
  ]
  (optional)
}
```

_Note: (optional) tag means that it is not strictly required to include certain field in request body._

__Response:__

```json
{
  "id": "uuid",
  "vendor_id": "string",
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": [
    {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    }
  ],
  "tags": [
    {
      "id": "uuid",
      "tag_name": "string"
    }
  ],
  "quantity": "int"
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid productId/product data)```
- ```404 Not Found (product not found)```
- ```500 Internal Server Error```

---

## 6. Delete Product

__DELETE ```api/products/:productId```__

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid productId)```
- ```404 Not Found (product not found)```
- ```500 Internal Server Error```

---

## 7. Delete Multiple Products

__DELETE ```api/products```__

__Body:__

```json
{
  "ids": [
    "uuid"
  ]
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid product ids)```
- ```404 Not Found (product not found)```
- ```500 Internal Server Error```

---

## 8. Get All Stocks Information

__GET ```api/stocks```__

__Query parameters:__

| Parameter   | Type   | Description                                                    |
|-------------|--------|----------------------------------------------------------------|
| sortBy      | string | Sort stocks by date_supplied field                             |
| sortOrder   | string | Sort order: asc/desc                                           |
| limit       | int    | Max amount of stocks to return                                 |
| offset      | int    | Number of stocks to skip before starting collecting the result |
| page        | int    | Page number to retrieve                                        |
| size        | int    | Number of stocks to return per page                            |
| location_id | string | Filter stocks by location_id                                   |

__Response:__

```json
[
  {
    "id": "uuid",
    "date_supplied": "date",
    "location_id": "uuid"
  }
]
```

__Status codes:__

- ```200 OK (success)```
- ```500 Internal Server Error```

---

## 9. Get One Stock Information

__GET ```api/stocks/:stockId```__

__Response:__

```json
{
  "id": "uuid",
  "vendor_id": "uuid",
  "date_supplied": "date",
  "location": {
    "id": "uuid",
    "city": "string",
    "address": "string"
  },
  "products": [
    {
      "product_id": "uuid",
      "name": "string",
      "quantity": "int",
      "unit_cost": "float64",
      "image": "string"
    }
  ]
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

## 10. Add A New Stock

__POST ```api/stocks```__

__Body:__

```json
{
  "date_supplied": "date",
  "location_id": "uuid",
  "products": [
    {
      "product_id": "uuid",
      "quantity": "int",
      "unit_cost": "float64"
    }
  ]
}
```

__Response:__

```json
{
  "id": "uuid",
  "date_supplied": "date",
  "location_id": "uuid"
}
```

__Status codes:__

- ```201 Created (success)```
- ```400 Bad Request (invalid stockId)```
- ```500 Internal Server Error```

---

## 11. Update Stock Information

__PUT ```api/stocks/:stockId```__

__Body:__

```json
{
  "date_supplied": "date",
  "location_id": "uuid"
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

## 12. Update Product Information From Stock

__PUT ```api/stocks/:stockId/products/:productId```__

__Body:__

```json
{
  "quantity": "int",
  "unit_cost": "float64"
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId/productId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

## 13. Modify Stock Information

__PATCH ```api/stocks/:stockId```__

__Body:__

```json
{
  "date_supplied": "date",
  (optional)
  "location_id": "uuid"
  (optional)
}
```

__Response:__

```json
{
  "id": "uuid",
  "date_supplied": "date",
  "location": {
    "id": "uuid",
    "city": "string",
    "address": "string"
  },
  "products": [
    {
      "product_id": "uuid",
      "name": "string",
      "quantity": "int",
      "unit_cost": "float64",
      "image": "string"
    }
  ]
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

## 14. Modify Products Information From Stock

__PATCH ```api/stocks/:stockId/products/```__

__Body:__

```json
[
  {
    "product_id": "uuid",
    "quantity": "int",
    (optional)
    "unit_cost": "float64"
    (optional)
  }
]
```

__Response:__

```json
[
  {
    "product_id": "uuid",
    "name": "string",
    "quantity": "int",
    "unit_cost": "float64",
    "image": "string"
  }
]
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

## 15. Modify Product Information From Stock

__PATCH ```api/stocks/:stockId/products/:productId```__

__Body:__

```json

{
  "quantity": "int",
  (optional)
  "unit_cost": "float64"
  (optional)
}

```

__Response:__

```json
{
  "product_id": "uuid",
  "name": "string",
  "quantity": "int",
  "unit_cost": "float64",
  "image": "string"
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId/productId/request data)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

# 16. Delete Stock

__DELETE ```api/stocks/:stockId```__

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

## 17. Delete Supplies From The Stock

__DELETE ```api/stocks```__

__Body:__

```json
{
  "ids": [
    "uuid"
  ]
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock not found)```
- ```500 Internal Server Error```

---

# 18. Delete Product From Supply From The Stock

__DELETE ```api/stocks/:stockId/products/:productId```__

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId/productId)```
- ```404 Not Found (stock product not found)```
- ```500 Internal Server Error```

---

## 19. Delete Products From Supply From The Stock

__DELETE ```api/stocks/:stockId/products```__

__Body:__

```json
{
  "ids": [
    "uuid"
  ]
}
```

__Status codes:__

- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock products not found)```
- ```500 Internal Server Error```

---
