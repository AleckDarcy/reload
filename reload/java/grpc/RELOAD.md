version: 1.10.1

export JAVA_HOME=$(/usr/libexec/java_home -v 1.8)


conda create --name proto351 protobuf=3.5.1
conda activate proto351

// rebuild jars
./gradlew build -x test

// rebuild service implementation
./gradlew installDist
