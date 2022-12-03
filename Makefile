
gen:
	kitex idl/world.proto

svc:
	kitex -service world idl/world.proto

wire:
	cd cmd && wire

run:
	cd cmd && go run .

try:
	cd client && go run .