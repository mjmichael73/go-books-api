### Database:
    - docker compose up --build -d
    - docker compose exec -it db bash
    - psql -h localhost -p 5432 -U postgres -d readinglist
    - RUN SQL commands in init.sql