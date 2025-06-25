up:
	@echo "Bhai ruk... docker build kar raha hoon "
	docker-compose up --build -d
	@echo "Sab kuch ready hai, backend udan pe hai "

down:
	@echo "Bhai sab kuch band kar raha hoon... shaanti chahiye "
	docker-compose down

dev:
	@echo "Development mode ON "
	go run main.go

db:
	@echo "Sirf Database start ho raha hai... chai bana tab tak "
	docker-compose up -d db

start-dev:
	@echo "Pehle DB chalu karte hain... bina uske kuch nahi hota "
	docker-compose up -d db
	@echo "AIR se hawa bhar raha hoon project mein... live reload ON.."
	air
