# Frontend Vendor API Contracts

---

## 1. Vendor Registration

__POST ```/auth/registration```__

__Body:__

```json
{
  "email": "string",
  "password": "string",
  "name": "string",
  "description": "string",
  "logo": "string",
  "address": "string",
  "website": "string"
}
```

__Response:__

```json
{
  "token": "jwt-token",
  "vendorId": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```

---

## 2. Vendor Sign In

__POST ```/auth/login```__

__Body:__

```json
{
  "email": "string",
  "password": "string"
}
```

__Response:__

```json
{
  "tokenSession": "jwt-token",
  "vendorSessionId": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```401 Unauthorized (authentication required or failed) ```

---

## 3. Check Vendor Profile Information

__GET ```/vendors/:vendorId```__

__Response:__

```json
{
  "id": "string",
  "email": "string",
  "name": "string",
  "description": "string",
  "logo": "string",
  "address": "string",
  "website": "string",
  "catalogId": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor account not found)```

---

## 4. Update Vendor Profile Information

__PUT ```/vendors/:vendorId```__

__Body:__

```json
{
  "id": "string",
  "email": "string",
  "password": "string",
  "name": "string",
  "description": "string",
  "logo": "string",
  "address": "string",
  "website": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor account not found)```

---


## 5. Check All Products Information

__GET ```/products```__

__Response:__

```json
[
  {
    "id": "string",
    "vendorId": "string",
    "name": "string",
    "price": "float64",
    "category": "string",
    "image": "string",
    "availability": "string"
  }
]
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor not found)```

---

## 6. Check One Product Information

__GET ```/products/:productId```__

__Response:__

```json
{
  "id": "string",
  "vendorId": "string",
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": ["string"],
  "tags": ["string"],
  "availability": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor/product not found)```

---

## 7. Add New Products

__POST ```/products```__

__Body:__

```json
{
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": ["string"],
  "tags": ["string"],
  "availability": "string"
}
```

__Response:__

```json
{
  "id": "string",
  "vendorId": "string",
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": ["string"],
  "tags": ["string"],
  "availability": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor not found)```

---


## 8. Update Product Information

__PUT ```/products/:productId```__

__Body:__

```json
{
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": ["string"],
  "tags": ["string"],
  "availability": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/product not found)```

---

## 9. Modify Product Information

__PATCH ```/products/:productId```__

__Body:__

```json
{
  "name": "string", (optional)
  "description": "string", (optional)
  "price": "float64", (optional)
  "category": "string", (optional)
  "images": ["string"], (optional)
  "tags": ["string"], (optional)
  "availability": "string", (optional)
}
```

_Note: (optional) tag means that it is not strictly required to include certain field in request body._

__Response:__

```json
{
  "id": "string",
  "vendorId": "string",
  "name": "string",
  "description": "string",
  "price": "float64",
  "category": "string",
  "images": ["string"],
  "tags": ["string"],
  "availability": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/product not found)```

---

## 10. Modify Products Information

__PATCH ```/products```__

__Body:__

```json
[
  {
    "name": "string", (optional)
    "description": "string", (optional)
    "price": "float64", (optional)
    "category": "string", (optional)
    "images": ["string"], (optional)
    "tags": ["string"], (optional)
    "availability": "string", (optional)
  }
]
```

_Note: (optional) tag means that it is not strictly required to include certain field in request body._

__Response:__

```json
[
  {
    "id": "string",
    "vendorId": "string",
    "name": "string",
    "description": "string",
    "price": "float64",
    "category": "string",
    "images": ["string"],
    "tags": ["string"],
    "availability": "string"
  }
]
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/products not found)```

---

## 11. Delete Product

__DELETE ```/products/:productId```__

__Response:__

```json
{
  "deletedId": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor/product not found)```

---

## 12. Delete Products

__DELETE ```/products```__

__Body:__

```json
{
  "idsToDelete": ["string"]
}
```

__Response:__

```json
{
  "deletedIds": ["string"]
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/products not found)```

---

## 13. Check All Products Information On Stock Page

__GET ```/products/stock```__

__Response:__

```json
[
  {
    "id": "string",
    "vendorId": "string",
    "name": "string",
    "image": "string",
    "availability": "string",
    "quantity": "int"
  }
]
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor not found)```

---

## 14. Check One Product Information From Stock Page

__GET ```/products/stock/:stockId```__

__Response:__

```json
{
  "id": "string",
  "name": "string",
  "image": "string",
  "availability": "string",
  "quantity": "int"
}
```

__Response:__

```json

[
  {
    "id": "string",
    "vendorId": "string",
    "name": "string",
    "image": "string",
    "availability": "string",
    "quantity": "int"
  }
]
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor/product not found)```

---

## 15. Update Product Quantity From Stock Page

__PUT ```/products/stock/:stockId```__

__Body:__

```json
{
  "quantity": "int"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/product not found)```

---

## 16. Check Reviews On Products

__GET ```/reviews```__

__Response:__

```json
[
  {
    "reviewId": "string",
    "productId": "string",
    "reviewerId": "string",
    "reviewerName": "string",
    "rating": "float64",
    "comment": "string",
    "date": "date"
  }
]
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor not found)```

---

## 17. Check Review On Product

__GET ```/reviews/:reviewId```__

__Response:__

```json
{
  "reviewId": "string",
  "productId": "string",
  "reviewerId": "string",
  "reviewerName": "string",
  "rating": "float64",
  "comment": "string",
  "date": "date",
  "replies": [
    {
      "replierId": "string",
      "name": "string",
      "comment": "string",
      "date": "date"
    }
  ]
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (vendor/review not found)```

---

## 18. Add a Reply On a Review

__POST ```/reviews/:reviewId/replies```__

__Body:__

```json
{
  "comment": "string"
}
```

__Response:__

```json
  {
  "replyId": "string",
  "reviewerId": "string",
  "comment": "string",
  "date": "date"
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/review not found)```

---

## 19. Updating Review Reply Comment

__PUT ```/reviews/:reviewId/replies/:replyId```__

__Body:__

```json
{
  "comment": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor/review/reply not found)```

---


## 20. Check Orders

__GET ```/orders```__

__Response:__

```json
{
  "orderId": "string",
  "customerId": "string",
  "vendorId": "string",
  "items": ["string"],
  "totalPrice": "float64",
  "status": "string",
  "date": "date",
  "vendorConfStatus": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```404 Not Found (vendor not found)```

---

## 21. Check One Order

__GET ```/orders/:orderId```__

__Response:__

```json
{
  "orderId": "string",
  "customerId": "string",
  "items": ["string"],
  "totalPrice": "float64",
  "status": "string",
  "date": "date",
  "vendorConfStatus": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```404 Not Found (vendor/order not found)```

---

## 22. Approve Orders

__PUT ```/orders/:orderId```__

__Body:__

```json
{
  "vendorConfStatus": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (vendor not found)```

---

## 23. Check Analytics

__GET ```/analytics```__

__Response:__

```json
{
  "to be determined": "to be determined"
}
```

__Status codes:__
- ```to be determined```








