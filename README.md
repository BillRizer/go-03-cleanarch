### update graphql code
go run github.com/99designs/gqlgen generate 

### update grpc code
protoc  --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

----



# how to run
```
git clone https://github.com/BillRizer/go-03-cleanarch.git

cd go-03-cleanarch

#run all 
docker compose up --build -d

>> o docker compose aguarda o mysql estar disponivel, isso pode demorar alguns segundos

```

## WEB
`http://localhost:8000/`

**create order**
```
POST - localhost:8000/order
use file  /api/create_order.http to run POST
```
   
**list orders**
```
get - localhost:8000/order
use file  /api/list_orders.http to run POST
```



## GraphQL  
`http://localhost:8080/`

**create order:**
```
mutation createOrder{
  createOrder(input: {
    id: "123",
    Price: 100.0,
    Tax: 10.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**list orders:**
```
query listOrders{
    listOrders{
        id
        Tax
        Price
        FinalPrice
    }
}
```
 

## GRPC 
`http://localhost:50051/`

Use evans to interact with GRPC resource
https://github.com/ktr0731/evans


```bash 
./evans -r repl --host 0.0.0.0 --port 3000
package pb
service OrderService
```
**create order:**
```
call CreateOrder

...
id (TYPE_STRING) =>
price (TYPE_FLOAT) =>  
tax (TYPE_FLOAT) => 
```

**list orders:**
```
call ListOrders
```



