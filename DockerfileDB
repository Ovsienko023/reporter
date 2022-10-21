FROM postgres:14
COPY init.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_PASSWORD=1234
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=postgres
ENV TZ=Europe/Moscow