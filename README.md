# Code Challenge: Go(lang)

## Goal of this project

The goal of this project is provide a transfer API between internal accounts of a digital bank.

## Requirements

- [Go 1.18](https://go.dev/)
- [Docker](https://www.docker.com/)

## Run Application

```
make run
```

## Run Tests

```
make test
```

## Routes:

### Create an account:

* **Method**: `POST`
* **URL**: `/accounts`
* **Body**:
  ```json
  {
    "name": "account-test",
    "cpf": "97014963023",
    "secret": "u69UFv$*ETH4",
    "balance": 100
  }
  ```
* **Responses**:
    * **Success**:
        * Status code: `201 - Created`
        * Example body:
      ```json
      {
        "account_id": 53
      }
      ```
    * **Errors**:
        * Request with invalid syntax
            * Status code: `400 - Bad Request`
            * Example body:
          ```json
          {
            "error": "secret must be at least 8 characters long"
          }
          ```
        * Request with CPF already registered
            * Status code: `409 - Conflict`
            * Example body:
          ```json
          {
            "error": "CPF already exists"
          }
          ```
        * Internal Server Error
            * Status code: `500 - Internal Server Error`

### Get account list:

* **Method**: `GET`
* **URL**: `/accounts`
* **Responses**:
    * **Success**:
        * Status code: `200 - OK`
        * Example body:
      ```json
      [
        {
          "id": 53,
          "name": "account-test",
          "cpf": "97014963023",
          "secret": "$2a$10$L61LEOytm.hUNSeq897S0eBhPwx14hMogi9wZ7EsHCsK115e2Oh7G",
          "balance": 100,
          "created_at": "2022-07-24T18:25:20.21256-03:00"
        }
      ]
      ```
    * **Errors**:
        * Internal Server Error
            * Status code: `500 - Internal Server Error`

### Get account balance:

* **Method**: `GET`
* **URL**: `/accounts/{account_id}/balance`
* **Responses**:
    * **Success**:
        * Status code: `200 - OK`
        * Example body:
      ```json
      {
        "balance": 100
      }
      ```
    * **Errors**:
        * Account not fount
            * Status code: `404 - Not Found`
            * Example body:
          ```json
          {
            "error": "record not found"
          }
          ```
        * Internal Server Error
            * Status code: `500 - Internal Server Error`

### Authenticate user:

* **Method**: `POST`
* **URL**: `/login`
* **Body**:
  ```json
  {
    "cpf": "97014963023",
    "secret": "u69UFv$*ETH4"
  }
  ```
* **Responses**:
    * **Success**:
        * Status code: `200 - OK`
        * Example body:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjo1MiwiZXhwIjoxNjU5MjI2MzkyfQ.xdoh-98tbj04-aYp7AtIiucI3YKvWZ5DvjvMoj9FZEo"
      }
      ```
    * **Errors**:
        * Account not fount
            * Status code: `404 - Not Found`
            * Example body:
          ```json
          {
            "error": "account not found"
          }
          ```
        * Unauthorized login
            * Status code: `401 - Unauthorized`
        * Request with invalid syntax
            * Status code: `400 - Bad Request`
        * Internal Server Error
            * Status code: `500 - Internal Server Error`

### Transfer between accounts:

* **Method**: `POST`
* **URL**: `/transfers`
* **Authorization**: `Bearer Token`
* **Body**:
  ```json
  {
    "account_destination_id": 54,
    "amount": 50
  }
  ```
* **Responses**:
    * **Success**:
        * Status code: `200 - OK`
        * Example body:
      ```json
      {
        "transaction_id": 1
      }
      ```
    * **Errors**:
        * Insufficient balance for transfer
            * Status code: `412 - Precondition Failed`
            * Example body:
          ```json
          {
            "error": "insufficient balance for transfer"
          }
          ```
        * Unauthorized user
            * Status code: `401 - Unauthorized`
        * Request with invalid syntax
            * Status code: `400 - Bad Request`
        * Internal Server Error
            * Status code: `500 - Internal Server Error`

### Get transfer list from authenticated user:

* **Method**: `GET`
* **URL**: `/transfers`
* **Responses**:
    * **Success**:
        * Status code: `200 - OK`
        * Example body:
      ```json
      [
        {
          "id": 1,
          "account_origin_id": 53,
          "account_destination_id": 54,
          "amount": 50,
          "created_at": "2022-07-24T18:25:20.21256-03:00"
        }
      ]
      ```
    * **Errors**:
        * Unauthorized user
            * Status code: `401 - Unauthorized`
        * Internal Server Error
            * Status code: `500 - Internal Server Error`

## Rules:

* ### Account:
    * `name`, `cpf`, and `secret` are required
    * `cpf` must contain 11 characters
    * `secret` must contain at least 8 characters
    * `secret` must be saved as hash
    * the default `balance` amount is 0
    * only one account per CPF

* ### Login:
    * `cpf` and `secret` are required
    * returns the token to be used in authenticated routes
    * the generated token is valid for 1 hour

* ### Transfer:
    * `account_destination_id` and `amount` are required
    * use bearer token authentication
    * the transfer amount must be greater than 0 and the account balance must be >= the amount to be transferred
