# uPay in Golang

Payment Gateway Microservice in Golang

## Feature

- SCA ready with [Stripe Payment Intents](https://stripe.com/docs/payments/payment-intents)
    - Off-session intents
    - Separation of auth and capture
        - New intent
        - Confirm intent
        - Capture/Delete intent
- No database infrastructure needed
- Stripe API keys configuration per currency

## Installation

```bash
cp config.json.dist config.json
vi config.json # Add your config values

ln -s ../../.github/hooks/pre-commit .git/hooks/pre-commit # If you want to contribute

cp config.json test/functional

go run main.go # go build main.go
```

## Tests

```bash
go test ./... # In another terminal
```

### APIs

- Swagger */swagger/index.html*
- [Postman collection](https://www.getpostman.com/collections/08908d8ba23942d002f6)

## TODO

See [*projects* section](https://github.com/lelledaniele/upaygo/projects)

### Goals

- Finish [Payment Intent project](https://github.com/lelledaniele/upaygo/projects/1) by the end of August 2019.
 **EU SCA law will be on duty after 14th September 2019**
