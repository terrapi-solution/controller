## Supported filter operators

The following filter operators are supported:

- `:`   The equality operator `filter=username:John` matches only when the username is exactly `John`
- `\>`  The greater than operator `filter=age>35` matches only when age is more than 35
- `\<`  The less than operator `filter=salary<80000` matches only when salary is less than 80,000
- `\>=` The greater than or equals to operator `filter=items>=100` matches only when items is at least 100
- `\<=` The less than or equals to operator `filter=score<=100000` matches when score is 100,000 or lower
- `\!=` The not equals to operator `state!=FAIL` matches when state has any value other than FAIL
- `\~`  The like operator `filter=lastName~illi` matches when lastName contains the substring `illi`

## Request example

```shell
curl -X GET http://localhost:8080/v1/modules?page=1&limit=10&order_by=username&order_direction=asc&filter="name:John"
```

## Model definition

```go
type UserModel struct {
    Username string `gorm:"uniqueIndex" filter:"param:login;searchable;filterable"`
    FullName string `filter:"searchable"`
    Role     string `filter:"filterable"`
}
```

The `param` tag defines a custom column name for the query parameter.
