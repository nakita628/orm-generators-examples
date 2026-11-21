// Runs the compat plan through the tables and entities hekireki-exposed wrote (copied to
// src/main/kotlin/models at run time). Exposed and the driver are the versions of
// checks/exposed/build.gradle.kts.
plugins {
    kotlin("jvm") version "2.3.20"
    application
}

repositories {
    mavenCentral()
}

dependencies {
    val exposed = "1.5.0"
    implementation("org.jetbrains.exposed:exposed-core:$exposed")
    implementation("org.jetbrains.exposed:exposed-jdbc:$exposed")
    implementation("org.jetbrains.exposed:exposed-dao:$exposed")
    implementation("org.jetbrains.exposed:exposed-json:$exposed")
    implementation("org.jetbrains.exposed:exposed-java-time:$exposed")
    implementation("org.postgresql:postgresql:42.7.13")
    implementation("io.github.thibaultmeyer:cuid:2.0.5")
    implementation("com.github.f4b6a3:ulid-creator:5.2.4")
    implementation("com.fasterxml.jackson.core:jackson-databind:2.20.1")
    implementation(kotlin("reflect"))
    runtimeOnly("org.slf4j:slf4j-nop:2.0.17")
}

kotlin {
    jvmToolchain(21)
}

application {
    mainClass.set("compat.MainKt")
}
