# Terraform Provider for an Example Expense API

For the original API, see
[here](https://github.com/joatmon08/dotnet-service-mesh-example).

## Tests

Start by running the following to check acceptance tests. Make sure
your API is running first!

```shell
EXPENSE_URL=http://localhost:5001 make testacc
```

You'll need to implement the following in order for the test to pass:

- `resourceExpenseUpdate`
- `resourceExpenseDelete`

