dev:
	rm -rf postgres-data/
	docker build -t buddytracker -f Dockerfile . && docker compose -f docker-compose.yml up
