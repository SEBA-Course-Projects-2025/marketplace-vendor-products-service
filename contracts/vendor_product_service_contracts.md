# Vendor Products Service API Contracts 

---

## 1. Get All Products Information

__GET ```api/products```__

__Response:__

```json
[
  {
    "id": "uuid",
    "vendorId": "string",
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
  "vendorId": "string",
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
  "images": ["string"],
  "tags": ["string"]
}
```

__Response:__

```json
{
  "id": "uuid",
  "vendorId": "string",
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
  "name": "string", (optional)
  "description": "string", (optional)
  "price": "float64", (optional)
  "category": "string", (optional)
  "images": [
    {
      "id": "uuid",
      "image_url": "string",
      "product_id": "uuid"
    }
  ], (optional)
  "tags": [
    {
      "id": "uuid",
      "tag_name": "string"
    }
  ] (optional)
}
```

_Note: (optional) tag means that it is not strictly required to include certain field in request body._

__Response:__

```json
{
  "id": "uuid",
  "vendorId": "string",
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
  "idsToDelete": ["uuid"]
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

__Response:__

```json
[
  {
    "id": "uuid",
    "dateSupplied": "date",
    "locationId": "uuid"
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
  "vendorId": "uuid",
  "dateSupplied": "date",
  "location": {
    "id": "uuid",
    "city": "string",
    "address": "string"
  },
  "products": [
    {
      "productId": "uuid",
      "name": "string",
      "quantity": "int",
      "unitCost": "float64",
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

## 10.  Add A New Stock 

__POST ```api/stocks```__

__Body:__

```json
{
  "dateSupplied": "date",
  "locationId": "uuid",
  "products": [
    {
      "productId": "uuid",
      "quantity": "int",
      "unitCost": "float64"
    }
  ]
}
```

__Response:__

```json
{
  "id": "uuid",
  "dateSupplied": "date",
  "locationId": "uuid"
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
  "dateSupplied": "date",
  "locationId": "uuid"
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
  "unitCost": "float64"
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
  "dateSupplied": "date", (optional)
  "locationId": "uuid" (optional)
}
```

__Response:__

```json
{
  "id": "uuid",
  "dateSupplied": "date",
  "location": {
    "id": "uuid",
    "city": "string",
    "address": "string"
  },
  "products": [
    {
      "productId": "uuid",
      "name": "string",
      "quantity": "int",
      "unitCost": "float64",
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
    "productId": "uuid",
    "quantity": "int", (optional)
    "unitCost": "float64" (optional)
  }
]
```
__Response:__

```json
[
  {
    "productId": "uuid",
    "name": "string",
    "quantity": "int",
    "unitCost": "float64",
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
  "quantity": "int", (optional)
  "unitCost": "float64" (optional)
}

```
__Response:__

```json
{
  "productId": "uuid",
  "name": "string",
  "quantity": "int",
  "unitCost": "float64",
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
  "idsToDelete": ["uuid"]
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
  "idsToDelete": ["uuid"]
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid stockId)```
- ```404 Not Found (stock products not found)```
- ```500 Internal Server Error```

---
