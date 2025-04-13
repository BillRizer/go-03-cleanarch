### update graphql code
go run github.com/99designs/gqlgen generate 

### update grpc code
protoc  --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

----



# how to run
```
git clone https://github.com/BillRizer/go-03-cleanarch.git

cd go-03-cleanarch

#run DB and Queue
docker compose up --build -d

cd cmd/ordersystem

# run application
go run main.go wire_gen.go 

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
    ListOrders{
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
./evans -r repl
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



