run:
	git pull
	git reset --hard origin/main
	sudo docker stop $$(sudo docker ps -q --filter ancestor=reporter_database)
	sudo docker stop $$(sudo docker ps -q --filter ancestor=reporter_server)
	sudo docker system prune -a
	sudo docker-compose down --volumes
	sudo docker-compose up -d --build
