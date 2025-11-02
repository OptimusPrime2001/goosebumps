routes -> controllers -> services -> repositories -> Models


Config
1. Models
   - Define the data structures and their relationships.
2. Repositories (Models)
   - Handle database operations for the models.
3. Services (Repositories)
   - Implement business logic and rules.
4. Controllers (Services)
   - Handle HTTP requests and responses.
5. Routes (Controllers)
   - Define the API endpoints and their handlers.

Command DB : 
- import-db:
   docker exec -i postgres_db psql -U trung_db -d user_manage < user_manage.sql
- export-db:
   docker exec -i postgres_db psql -U trung_db -d user_manage > user_manage.sql