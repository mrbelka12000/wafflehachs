postgres:
	@docker run --name=myForum -e POSTGRES_PASSWORD='mrbelka12000' -p 5678:5432 -d postgres