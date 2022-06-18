postgres:
	@docker run --name=myForum -e POSTGRES_PASSWORD='mrbelka12000' -p 5678:5432 -d postgres

dbuild:
	@docker image build -f Dockerfile -t forumimage .
	@docker container run -p 8080:8080 --detach --name forum forumimage
