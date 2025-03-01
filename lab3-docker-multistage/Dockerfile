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

# Run precompiled app
CMD ["java", "-jar", "app.jar"]
