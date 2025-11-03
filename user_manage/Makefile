importdb:
	docker exec -i postgres_db psql -U trung_db -d user_manage < user_manage.sql
exportdb:
	docker exec -i postgres_db pg_dump -U trung_db -d user_manage > user_manage.sql

