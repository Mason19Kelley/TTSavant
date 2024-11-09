# Use the official PostgreSQL image from the Docker Hub
FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=postgres

# Expose the PostgreSQL port
EXPOSE 5432

# Add a custom initialization script if needed
COPY ./schema/init.sql /docker-entrypoint-initdb.d/

# Command to run the PostgreSQL server
CMD ["postgres"]