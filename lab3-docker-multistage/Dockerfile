# Stage 1
FROM maven:3.9-eclipse-temurin-17-alpine AS builder

# Setting working directory and copying all source files
WORKDIR /app
COPY ./spring-boot ./

# Build maven app
RUN mvn clean install

# Stage 2
FROM maven:3.9-eclipse-temurin-17-alpine

# Copying compiled jar file
COPY --from=builder /app/target/*.jar app.jar

# Creating user
RUN addgroup -S app_user_group && \
        adduser -S app_user -G app_user_group && \
        chown -R app_user:app_user_group \
            /app/target/*.jar \
            app.jar
USER app_user

# Run precompiled app
CMD ["java", "-jar", "app.jar"]
