module github.com/fadilrifqi/orchestra

go 1.24.1

require (
	github.com/fadilrifqi/order-service v0.0.0-00010101000000-000000000000
	github.com/fadilrifqi/payment-service v0.0.0-00010101000000-000000000000
	github.com/fadilrifqi/shipping-service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)

replace github.com/fadilrifqi/payment-service => ../payment-service

replace github.com/fadilrifqi/order-service => ../order-service

replace github.com/fadilrifqi/shipping-service => ../shipping-service
