FROM gradle:jdk21-alpine

# Copy all app data into container
COPY ../ /app/

# Setting workdir for convenience
WORKDIR /app/

# Building a gradle app
RUN gradle clean build

# Running the app
ENTRYPOINT ["gradle", ":run"]