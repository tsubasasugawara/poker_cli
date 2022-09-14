start:
	cd ./client && make start && cd ../server && make start && cd ..

stop:
	cd ./client && make stop && cd ../server && make stop && cd ..
	
web:
	cd ./server && make web && cd ..
	
db:
	cd ./server && make db && cd ..

go:
	cd ./client && make go && cd ..
