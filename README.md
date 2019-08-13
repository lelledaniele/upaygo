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
go run *.go # go build ./...
```

## Doc

See [*wiki* section](https://github.com/lelledaniele/upaygo/wiki)

### APIs

- Swagger */swagger/index.html*
- [Postman collection](https://www.getpostman.com/collections/08908d8ba23942d002f6)

## TODO

See [*projects* section](https://github.com/lelledaniele/upaygo/projects)

