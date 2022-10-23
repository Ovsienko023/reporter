run:
	git fetch
	git reset --hard HEAD
	sudo docker stop $(docker ps -a -q)
	sudo docker-compose down --volumes
	sudo docker-compose up -d --build
