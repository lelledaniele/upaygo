# uPay in Golang

Payment Gateway Microservice in Golang

## PSD2 SCA

**EU SCA law will be on duty after 14th September 2019**

### Updates
 
- 13/08/2019 - For **UK cards** the [SCA implementation deadline is March 2021](https://www.fca.org.uk/news/press-releases/fca-agrees-plan-phased-implementation-strong-customer-authentication)

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

# API doc
swag init

# If you want to contribute
cp .github/hooks/pre-commit .git/hooks/pre-commit
# Open and change absolute config path

go run main.go # go build main.go
```

## Tests

```bash
go test ./... -failfast -tags=unit
go test ./... -failfast -tags=stripe -config=ABS_PATH/config.json
```

### APIs

- Swagger */swagger/index.html*
- [Postman collection](https://www.getpostman.com/collections/08908d8ba23942d002f6)

## TODO

See [*projects* section](https://github.com/lelledaniele/upaygo/projects)
