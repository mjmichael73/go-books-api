### Database:
    - docker compose up --build -d
    - docker compose exec -it db bash
    - psql -h localhost -p 5432 -U postgres
    - RUN SQL commands in init.sql
    - export READINGLIST_DB_DSN='postgres://readinglist:pa55w0rd@localhost/readinglist?sslmode=disable'