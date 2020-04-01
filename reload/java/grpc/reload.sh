export JAVA_HOME=$(/usr/libexec/java_home -v 1.8)
./gradlew build -x test

cp core/build/libs/grpc-core-1.10.1.jar ./grpc-core-1.12.0.jar
cp stub/build/libs/grpc-stub-1.10.1.jar ./grpc-stub-1.12.0.jar

cp -f grpc-stub-1.12.0.jar ~/.gradle/caches/modules-2/files-2.1/io.grpc/grpc-stub/1.12.0/fbd2bafe09a89442ab3d7a8d8b3e8bafbd59b4e0/
cp -f grpc-core-1.12.0.jar ~/.gradle/caches/modules-2/files-2.1/io.grpc/grpc-core/1.12.0/541a5c68ce85c03190e29bc9e0ec611d2b75ff24/
