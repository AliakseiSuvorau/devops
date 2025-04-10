FROM gradle:jdk21-alpine

# Copy all app data into container
COPY ../ /app/

# Setting workdir for convenience
WORKDIR /app/

# Creating user and building a gradle app
RUN gradle clean build && \
        rm -rf gradle gradle.properties && \
        addgroup -S app_user_group && \
        adduser -S app_user -G app_user_group && \
        chown -R app_user:app_user_group .
USER app_user

# Running the app
ENTRYPOINT ["gradle", ":run"]
