How to Create a Release of GRPC Java (for Maintainers Only)
===============================================================

Build Environments
------------------
We deploy GRPC to Maven Central under the following systems:
- Ubuntu 14.04 with Docker 1.6.1 that runs CentOS 6.6
- Windows 7 64-bit with MSYS2 with mingw32 and mingw64
- Mac OS X 10.9.5

Other systems may also work, but we haven't verified them.

Prerequisites
-------------

### Setup OSSRH and Signing

If you haven't deployed artifacts to Maven Central before, you need to setup
your OSSRH (OSS Repository Hosting) account and signing keys.
- Follow the instructions on [this
  page](http://central.sonatype.org/pages/ossrh-guide.html) to set up an
  account with OSSRH.
  - You only need to create the account, not set up a new project
  - Contact a gRPC maintainer to add your account after you have created it.
- (For release deployment only) Install GnuPG and [generate your key
  pair](https://www.gnupg.org/documentation/howtos.html). You'll also
  need to [publish your public key](https://www.gnupg.org/gph/en/manual.html#AEN464)
  to make it visible to the Sonatype servers
  (e.g. `gpg --keyserver pgp.mit.edu --send-key <key ID>`).
- Put your GnuPG key password and OSSRH account information in
  `<your-home-directory>/.gradle/gradle.properties`.

```
# You need the signing properties only if you are making release deployment
signing.keyId=<8-character-public-key-id>
signing.password=<key-password>
signing.secretKeyRingFile=<your-home-directory>/.gnupg/secring.gpg

ossrhUsername=<ossrh-username>
ossrhPassword=<ossrh-password>
checkstyle.ignoreFailures=false
```

### Build Protobuf
Protobuf libraries are needed for compiling the GRPC codegen. Despite that you
may have installed Protobuf on your system, you may want to build Protobuf
separately and install it under your personal directory, because

1. The Protobuf version installed on your system may be different from what
   GRPC requires. You may not want to pollute your system installation.
2. We will deploy both 32-bit and 64-bit versions of the codegen, thus require
   both variants of Protobuf libraries. You don't want to mix them in your
   system paths.

Please see the [Main Readme](README.md) for details on building protobuf.

Common Variables
----------------
Many of the following commands expect release-specific variables to be set. Set
them before continuing, and set them again when resuming.

```bash
$ MAJOR=1 MINOR=7 PATCH=0 # Set appropriately for new release
$ VERSION_FILES=(
  build.gradle
  android-interop-testing/app/build.gradle
  core/src/main/java/io/grpc/internal/GrpcUtil.java
  cronet/build.gradle
  examples/build.gradle
  examples/pom.xml
  examples/android/helloworld/app/build.gradle
  examples/android/routeguide/app/build.gradle
  )
```


Branching the Release
---------------------
The first step in the release process is to create a release branch and bump
the SNAPSHOT version. Our release branches follow the naming
convention of `v<major>.<minor>.x`, while the tags include the patch version
`v<major>.<minor>.<patch>`. For example, the same branch `v1.7.x`
would be used to create all `v1.7` tags (e.g. `v1.7.0`, `v1.7.1`).

1. For `master`, change root build files to the next minor snapshot (e.g.
   ``1.8.0-SNAPSHOT``).

   ```bash
   $ git checkout -b bump-version master
   # Change version to next minor (and keep -SNAPSHOT)
   $ sed -i 's/[0-9]\+\.[0-9]\+\.[0-9]\+\(.*CURRENT_GRPC_VERSION\)/'$MAJOR.$((MINOR+1)).0'\1/' \
     "${VERSION_FILES[@]}"
   $ ./gradlew build
   $ git commit -a -m "Start $MAJOR.$((MINOR+1)).0 development cycle"
   ```
2. Go through PR review and submit.
3. Create the release branch starting just before your commit and push it to GitHub:

   ```bash
   $ git fetch upstream
   $ git checkout -b v$MAJOR.$MINOR.x \
     $(git log --pretty=format:%H --grep "^Start $MAJOR.$((MINOR+1)).0 development cycle$" upstream/master)^
   $ git push upstream v$MAJOR.$MINOR.x
   ```
4. Go to [Travis CI settings](https://travis-ci.org/grpc/grpc-java/settings) and
   add a _Cron Job_:
   * Branch: `v$MAJOR.$MINOR.x`
   * Interval: `weekly`
   * Options: `Do not run if there has been a build in the last 24h`
   * Click _Add_ button
5. Continue with Google-internal steps at go/grpc/java/releasing.
6. Create a milestone for the next release.
7. Move items out of the release milestone that didn't make the cut. Issues that
   may be backported should stay in the release milestone. Treat issues with the
   'release blocker' label with special care.

Tagging the Release
-------------------

1. Verify there are no open issues in the release milestone. Open issues should
   either be deferred or resolved and the fix backported.
2. For vMajor.Minor.x branch, change `README.md` to refer to the next release
   version. _Also_ update the version numbers for protoc if the protobuf library
   version was updated since the last release.

   ```bash
   $ git checkout -b release v$MAJOR.$MINOR.x
   # Bump documented versions. Don't forget protobuf version
   $ ${EDITOR:-nano -w} README.md
   $ git commit -a -m "Update README to reference $MAJOR.$MINOR.$PATCH"
   ```
3. Change root build files to remove "-SNAPSHOT" for the next release version
   (e.g. `0.7.0`). Commit the result and make a tag:

   ```bash
   # Change version to remove -SNAPSHOT
   $ sed -i 's/-SNAPSHOT\(.*CURRENT_GRPC_VERSION\)/\1/' "${VERSION_FILES[@]}"
   $ sed -i s/-SNAPSHOT// compiler/src/test{,Lite,Nano}/golden/TestService.java.txt
   $ ./gradlew build
   $ git commit -a -m "Bump version to $MAJOR.$MINOR.$PATCH"
   $ git tag -a v$MAJOR.$MINOR.$PATCH -m "Version $MAJOR.$MINOR.$PATCH"
   ```
4. Change root build files to the next snapshot version (e.g. `0.7.1-SNAPSHOT`).
   Commit the result:

   ```bash
   # Change version to next patch and add -SNAPSHOT
   $ sed -i 's/[0-9]\+\.[0-9]\+\.[0-9]\+\(.*CURRENT_GRPC_VERSION\)/'$MAJOR.$MINOR.$((PATCH+1))-SNAPSHOT'\1/' \
     "${VERSION_FILES[@]}"
   $ sed -i s/$MAJOR.$MINOR.$PATCH/$MAJOR.$MINOR.$((PATCH+1))-SNAPSHOT/ compiler/src/test{,Lite,Nano}/golden/TestService.java.txt
   $ ./gradlew build
   $ git commit -a -m "Bump version to $MAJOR.$MINOR.$((PATCH+1))-SNAPSHOT"
   ```
5. Go through PR review and push the release tag and updated release branch to
   GitHub:

   ```bash
   $ git checkout v$MAJOR.$MINOR.x
   $ git merge --ff-only release
   $ git push upstream v$MAJOR.$MINOR.$PATCH
   $ git push upstream v$MAJOR.$MINOR.x
   ```
6. Close the release milestone.

Setup Build Environment
---------------------------

### Linux
The deployment for Linux uses [Docker](https://www.docker.com/) running
CentOS 6.6 in order to ensure that we have a consistent deployment environment
on Linux. You'll first need to install Docker if not already installed on your
system.  Make sure to have at least version 1.7.1 or later.

1. Under the [Protobuf source directory](https://github.com/google/protobuf), 
   build the `protoc-artifacts` image:
   
   ```bash
   protobuf$ docker build -t protoc-artifacts protoc-artifacts
   ```
2. Under the grpc-java source directory, build the `grpc-java-deploy` image:

   ```bash
   grpc-java$ docker build -t grpc-java-deploy compiler
   ```
3. Start a Docker container that has the deploy environment set up for you. The
   GRPC source is cloned into `/grpc-java`.
   
   ```bash
   $ docker run -it --rm=true grpc-java-deploy
   ```
   
   Note that the container will be deleted after you exit. Any changes you have
   made (e.g., copied configuration files) will be lost. If you want to keep the
   container, remove `--rm=true` from the command line.
4. Next, you'll need to copy your OSSRH credentials and GnuPG keys to your docker container.
   In Docker:
   ```
   # mkdir /root/.gradle
   ```
   Find the container ID in your bash prompt, which is shown as `[root@<container-ID> ...]`.
   In host:
   ```
   $ docker cp ~/.gnupg <container-ID>:/root/
   $ docker cp ~/.gradle/gradle.properties <container-ID>:/root/.gradle/
   ```
   
   You'll also need to update `signing.secretKeyRingFile` in
   `/root/.gradle/gradle.properties` to point to `/root/.gnupg/secring.gpg`.

### Windows

#### Windows 64-bit with MSYS2 (Recommended for Windows)
Because the gcc shipped with MSYS2 doesn't support multilib, you have to
compile and deploy 32-bit and 64-bit binaries in separate steps.

##### Under MinGW-w64 Win32 Shell
1. Compile and install 32-bit protobuf:

   ```bash
   protobuf$ ./configure --disable-shared --prefix=$HOME/protobuf-32
   protobuf$ make clean && make && make install
   ```
2. Configure CXXFLAGS needed by the protoc plugin when building.

   ```bash
   grpc-java$ export CXXFLAGS="-I$HOME/protobuf-32/include" \
     LDFLAGS="-L$HOME/protobuf-32/lib"
   ```

##### Under MinGW-w64 Win64 Shell
1. Compile and install 64-bit protobuf:

   ```bash
   protobuf$ ./configure --disable-shared --prefix=$HOME/protobuf-64
   protobuf$ make clean && make && make install
   ```
2. Configure CXXFLAGS needed by the protoc plugin when building.

   ```bash
   grpc-java$ export CXXFLAGS="-I$HOME/protobuf-64/include" \
     LDFLAGS="-L$HOME/protobuf-64/lib"
   ```

#### Windows 64-bit with Cygwin64 (TODO: incomplete)
Because the MinGW gcc shipped with Cygwin64 doesn't support multilib, you have
to compile and deploy 32-bit and 64-bit binaries in separate steps.

1. Compile and install 32-bit protobuf. `-static-libgcc -static-libstdc++` are
   needed for `protoc` to be successfully run in the unit test.
   
   ```bash
   protobuf$ LDFLAGS="-static-libgcc -static-libstdc++" ./configure --host=i686-w64-mingw32 --disable-shared --prefix=$HOME/protobuf-32
   protobuf$ make clean && make && make install
   ```

2. Compile and install 64-bit protobuf:

   ```bash
   protobuf$ ./configure --host=x86_64-w64-mingw32 --disable-shared --prefix=$HOME/protobuf-64
   protobuf$ make clean && make && make install
   ```

### Mac
Please refer to [Protobuf
README](https://github.com/google/protobuf/blob/master/README.md) for how to
set up GCC and Unix tools on Mac.

Mac OS X has been 64-bit-only since 10.7 and we are compiling for 10.7 and up.
We only build 64-bit artifact for Mac.

1. Compile and install protobuf:

   ```bash
   protobuf$ CXXFLAGS="-m64" ./configure --disable-shared --prefix=$HOME/protobuf
   protobuf$ make clean && make && make install
   ```
2. Configure CXXFLAGS needed by the protoc plugin when building.

   ```bash
   grpc-java$ export CXXFLAGS="-I$HOME/protobuf/include" \
     LDFLAGS="$HOME/protobuf/lib/libprotobuf.a $HOME/protobuf/lib/libprotoc.a"
   ```

Build and Deploy
----------------
We currently distribute the following OSes and architectures:

| OS | x86_32 | x86_64 |
| --- | --- | --- |
| Linux | X | X |
| Windows | X | X |
| Mac |  | X |

Deployment to Maven Central (or the snapshot repo) is a two-step process. The only
artifact that is platform-specific is codegen, so we only need to deploy the other
jars once. So the first deployment is for all of the artifacts from one of the selected
OS/architectures. After that, we then deploy the codegen artifacts for the remaining
OS/architectures.

**NOTE: _Before building/deploying, be sure to switch to the appropriate branch or tag in
the grpc-java source directory._**

### First Deployment

As stated above, this only needs to be done once for one of the selected OS/architectures.
The following command will build the whole project and upload it to Maven
Central. Parallel building [is not safe during
uploadArchives](https://issues.gradle.org/browse/GRADLE-3420).
```bash
grpc-java$ ./gradlew clean build && ./gradlew -Dorg.gradle.parallel=false uploadArchives
```

If the version has the `-SNAPSHOT` suffix, the artifacts will automatically
go to the snapshot repository. Otherwise it's a release deployment and the
artifacts will go to a freshly created staging repository.

### Deploy GRPC Codegen for Additional Platforms
The previous step will only deploy the codegen artifacts for the OS you run on
it and the architecture of your JVM. For a fully fledged deployment, you will
need to deploy the codegen for all other supported OSes and architectures.

To deploy the codegen for an OS and architecture, you must run the following
commands on that OS and specify the architecture by the flag `-PtargetArch=<arch>`.

If you are doing a snapshot deployment:

```bash
grpc-java$ ./gradlew clean grpc-compiler:build grpc-compiler:uploadArchives \
    -PtargetArch=<arch> -Dorg.gradle.parallel=false
```

When deploying a Release, the first deployment will create
[a new staging repository](https://oss.sonatype.org/#stagingRepositories). You'll need
to look up the ID in the OSSRH UI (usually in the form of `iogrpc-*`). Codegen
deployment commands should include `-PrepositoryId=<repository-id>` in order to
ensure that the artifacts are pushed to the same staging repository.

```bash
grpc-java$ ./gradlew clean grpc-compiler:build grpc-compiler:uploadArchives -PtargetArch=<arch> \
    -PrepositoryId=<repository-id> -Dorg.gradle.parallel=false
```

Releasing on Maven Central
--------------------------
Once all of the artifacts have been pushed to the staging repository, the
repository must first be `closed`, which will trigger several sanity checks
on the repository. If this completes successfully, the repository can then
be `released`, which will begin the process of pushing the new artifacts to
Maven Central (the staging repository will be destroyed in the process). You can
see the complete process for releasing to Maven Central on the [OSSRH
site](http://central.sonatype.org/pages/releasing-the-deployment.html).

Build interop container image
-----------------------------

We have containers for each release to detect compatibility regressions with old
releases. Generate one for the new release by following the
[GCR image generation instructions](https://github.com/grpc/grpc/blob/master/tools/interop_matrix/README.md#step-by-step-instructions-for-adding-a-gcr-image-for-a-new-release-for-compatibility-test).

Update README.md
----------------
After waiting ~1 day and verifying that the release appears on [Maven
Central](http://mvnrepository.com/), cherry-pick the commit that updated the
README into the master branch and go through review process.

```
$ git checkout -b bump-readme master
$ git cherry-pick v$MAJOR.$MINOR.$PATCH^
```

Update version referenced by tutorials
--------------------------------------

Update the `grpc_java_release_tag` in
[\_data/config.yml](https://github.com/grpc/grpc.github.io/blob/master/_data/config.yml)
of the grpc.github.io repository.

Notify the Community
--------------------
Finally, document and publicize the release.

1. Add [Release Notes](https://github.com/grpc/grpc-java/releases) for the new tag.
   The description should include any major fixes or features since the last release.
   You may choose to add links to bugs, PRs, or commits if appropriate.
2. Post a release announcement to [grpc-io](https://groups.google.com/forum/#!forum/grpc-io)
   (`grpc-io@googlegroups.com`). The title should be something that clearly identifies
   the release (e.g.`GRPC-Java <tag> Released`).

Update Hosted Javadoc
---------------------

Now we need to update gh-pages with the new Javadoc:

```bash
git checkout gh-pages
rm -r javadoc/
wget -O grpc-all-javadoc.jar "http://search.maven.org/remotecontent?filepath=io/grpc/grpc-all/$MAJOR.$MINOR.$PATCH/grpc-all-$MAJOR.$MINOR.$PATCH-javadoc.jar"
unzip -d javadoc grpc-all-javadoc.jar
patch -p1 < ga.patch
rm grpc-all-javadoc.jar
rm -r javadoc/META-INF/
git add -A javadoc
git commit -m "Javadoc for $MAJOR.$MINOR.$PATCH"
```

Push gh-pages to the main repository and verify the current version is [live
on grpc.io](https://grpc.io/grpc-java/javadoc/).
