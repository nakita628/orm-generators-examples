# Gradle with JDK 21 and the dependencies of build.gradle.kts in its cache: the Exposed harness
# builds offline at run time, with the generated code in src/main/kotlin/models.
FROM gradle:8.14-jdk21
WORKDIR /harness
COPY settings.gradle.kts build.gradle.kts ./
RUN mkdir -p src/main/kotlin/compat && printf 'package compat\nfun main() {}\n' > src/main/kotlin/compat/Main.kt && gradle installDist --no-daemon -q && rm -rf build src
