# Use the official PostgreSQL image from the Docker Hub
FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=myuser
ENV POSTGRES_DB=gator

# Expose the default PostgreSQL port
EXPOSE 5432

# Add any custom initialization scripts to the Docker image
COPY init_db.sh /docker-entrypoint-initdb.d/

# Command to run the PostgreSQL server
CMD ["postgres"]