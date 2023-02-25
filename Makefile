.PHONY: run
run:
	git pull
	git reset --hard origin/main
	sudo docker stop $$(sudo docker ps -q --filter ancestor=reporter_database)
	sudo docker stop $$(sudo docker ps -q --filter ancestor=reporter_server)
	sudo docker system prune -a
	sudo docker-compose up -d --build

.PHONY: cleanly
cleanly:
	git pull
	git reset --hard origin/main
	sudo docker stop $$(sudo docker ps -q --filter ancestor=reporter_database)
	sudo docker stop $$(sudo docker ps -q --filter ancestor=reporter_server)
	sudo docker system prune -a
	sudo docker-compose down --volumes
	sudo docker-compose up -d --build


.PHONY: build
build:
	sudo docker-compose up --build


.PHONY: deploy
deploy:
	sudo docker system prune -a
	sudo docker-compose --file docker-compose.deploy.yaml up --build
