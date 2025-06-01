# Frontend  Storefront API Contract

## 1. Customer Registration (probably will be reworked with an external provider)

**POST** `/auth/registration`

**Body**:
```json
{
  "email": "string",
  "password": "string",
  "shippingAddress": "string",
  "name": "string"
}
```

**Response**:
```json
{
  "token": "jwt-token",
  "customerId": "string"
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```

---

## 2. Customer Sign In

**POST** `/auth/login`

**Body**:
```json
{
  "email": "string",
  "password": "string"
}
```

**Response**:
```json
{
  "token": "jwt-token",
  "customerId": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```401 Unauthorized (authentication required or failed) ```

---

## 3. Check Customer Profile

**GET** `/customers/:customerId`

**Response**:
```json
{
  "customerId": "string",
  "email": "string",
  "shippingAddress": "string",
  "orders": ["order1", "order2"]
}
```
__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (customer not found)```
---

## 4. Update Customer Profile Information

__PUT ```/customers/:customerId```__

__Body:__

```json
{
  "email": "string",
  "password": "string",
  "name": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (customer not found)```

---

## 5. Add to Cart

**POST** `/cart/items`

**Body**:
```json
{
  "productId": "string",
  "quantity": "int"
}
```

**Response**:
```json
{
  "cartId": "string",
  "items": [
    {
      "productId": "string",
      "name": "string",
      "price": "float64",
      "quantity": "int"
    }
  ]
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (cart/product not found)```

---

## 6. Get Cart

**GET** `/cart`

**Response**:
```json
{
  "cartId": "string",
  "items": [
    {
      "productId": "string",
      "name": "string",
      "price": "float64",
      "quantity": "int"
    }
  ]
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (cart/product not found)```

---

## 7. Update Cart Item

**PUT** `/cart/items/:productId`

**Body**:
```json
{
  "quantity": 2
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (product not found)```

---

## 8. Remove from Cart

**DELETE** `/cart/items/:productId`

__Response:__

```json
{
  "deletedProductId": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (product not found)```

---

## 9. Checkout Initiation

**POST** `/checkout`

**Body**:
```json
{
  "customerId": "string",
  "shippingAddress": "string"
}
```

**Response**:
```json
{
  "checkoutId": "string",
  "paymentUrl": "https://external-gateway.com/pay?checkoutId=..."
}
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (customer not found)```

---

## 10. Check Liked Products

__GET ```/liked-products```__

**Response**:
```json
[
  {
    "productId": "string",
    "name": "string",
    "price": "float64",
    "availability": "string",
    "image": "string"
  }
]
```

__Status codes:__
- ```201 Created (success)```
- ```400 Bad Request (invalid input)```

---

## 11. Update Liked Products

**PUT** `/liked-products/:likedProductId`

**Body**:
```json
{
  "liked": "bool"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request (invalid input)```
- ```404 Not Found (product not found)```

---

## 12. Remove Liked Products

**DELETE** `/liked-products`

**Body**:

```json
{
  "idsToDelete": ["string"] 
}
```

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (products not found)```

---

## 13. Remove Liked Product

**DELETE** `/liked-products/:likedProductId`

__Status codes:__
- ```200 OK (success)```
- ```404 Not Found (product not found)```
