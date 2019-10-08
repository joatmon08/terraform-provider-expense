# Terraform Provider for an Example Expense API


## Start

Download [this example](https://github.com/joatmon08/dotnet-service-mesh-example).

Run:

```shell
docker-compose up -d
```

This will create the API stack.

## Tests

Start by running the following to check acceptance tests. Make sure
your API is running first!

```shell
EXPENSE_URL=http://localhost:5001 make testacc
```

You'll need to implement the following in order for the test to pass:

- `resourceExpenseUpdate`
- `resourceExpenseDelete`


## Build

```shell
make plugin
```

## Running the Provider

- Go to `examples/`.
- Make sure you've added it to your plugins directory.
- Run `terraform init`.
- Run `EXPENSE_URL=http://localhost:5001 terraform plan`. You'll need the URL.
- Run `EXPENSE_URL=http://localhost:5001 terraform apply`. You'll need the URL.

