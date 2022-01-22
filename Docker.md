## Docker

[Docker](https://www.docker.com) is a tool that enables packaging software as a [container image](https://www.docker.com/resources/what-container) that includes all of the dependencies to run a given application. In the context of this project, Docker enables building and running the `tbb` application across multiple operating systems.

### Installation

To get started, follow the official [Getting Started](https://docs.docker.com/get-docker/) guide to install Docker for your operating system. If you are using Linux, be sure to also follow [the guide for installing docker-compose](https://docs.docker.com/compose/install/#install-compose-on-linux-systems). This extra step is not necessary on Windows and MacOS.

Once all tools are installed, you will have two new executables that you can run in your terminal:
1. `docker` - This is the application used for building and running containers
2. `docker-compose` - [Docker Compose](https://docs.docker.com/compose/) is a tool for defining how to run any number of containers

You can check that both of these applications are installed by running the following commands in your terminal:

```
$ which docker
/usr/local/bin/docker

$ which docker-compose
/usr/local/bin/docker-compose
```

This uses the UNIX [which](https://www.tutorialspoint.com/unix_commands/which.htm) command to print the absolute path to the executable in question. As long as a path is printed for both `docker` and `docker-compose` then you are good to go - the paths do not need to be the same as those shown above.

### Building

You can build the Docker image for the `tbb` application by running `make image` in the root of this repository. This will take a minute or two to build and the output in your terminal will look something like this:

```
$ make image
docker build -t tbb:latest .
[+] Building 36.0s (9/9) FINISHED
 => [internal] load build definition from Dockerfile                               0.0s
 => => transferring dockerfile: 137B                                               0.0s
 => [internal] load .dockerignore                                                  0.0s
 => => transferring context: 2B                                                    0.0s
 => [internal] load metadata for docker.io/library/golang:1.16                     0.4s
 => [internal] load build context                                                  0.2s
 => => transferring context: 6.78MB                                                0.2s
 => [1/4] FROM docker.io/library/golang:1.16@sha256:8f29258b4b992b383d03290acde77  0.0s
 => CACHED [2/4] WORKDIR /build                                                    0.0s
 => [3/4] ADD ./ /build                                                            0.1s
 => [4/4] RUN make install                                                        33.0s
 => exporting to image                                                             2.3s
 => => exporting layers                                                            2.3s
 => => writing image sha256:b088ed645ed132544fe8257284a5ed03961da0801016a5544a5a6  0.0s
 => => naming to docker.io/library/tbb:latest                                      0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
```

You can verify that the image was built successfully by running the command `docker images` in your terminal:

```
$ docker images
REPOSITORY               TAG       IMAGE ID       CREATED          SIZE
tbb                      latest    b088ed645ed1   15 seconds ago   1.12GB
```

Now that you have the image available, you can run it like so:

```
$ docker run -ti tbb
The Blockchain Bar CLI

Usage:
  tbb [flags]
  tbb [command]

Available Commands:
  balances    Interacts with balances (list...).
  help        Help about any command
  run         Launches the TBB node and its HTTP API.
  version     Describes version.
  wallet      Manages blockchain accounts and keys.

Flags:
  -h, --help   help for tbb

Use "tbb [command] --help" for more information about a command.
```

### Docker for Local Development

Docker makes it much easier to package and distribute a runnable image of your application. However, Docker can be used for local development as well. This is useful in many situations, such as if your Operating System or CPU architecture is not supported by the libraries used by `tbb`. In this case, [Docker Compose](https://docs.docker.com/compose/) can be used to give you a running [shell](https://en.wikipedia.org/wiki/Shell_(computing)) within a `tbb` container. You can even mount this repository within the container so that the changes you make to the code can be immediately run within the container.

To try this out, run the command `make local` in the root of this repository:

```
$ make local
docker-compose build
Building dev
[+] Building 36.7s (9/9) FINISHED
 => [internal] load build definition from Dockerfile                                0.0s
 => => transferring dockerfile: 137B                                                0.0s
 => [internal] load .dockerignore                                                   0.0s
 => => transferring context: 2B                                                     0.0s
 => [internal] load metadata for docker.io/library/golang:1.16                      1.1s
 => [internal] load build context                                                   0.0s
 => => transferring context: 9.17kB                                                 0.0s
 => [1/4] FROM docker.io/library/golang:1.16@sha256:8f29258b4b992b383d03290acde77c  0.0s
 => CACHED [2/4] WORKDIR /build                                                     0.0s
 => [3/4] ADD ./ /build                                                             0.1s
 => [4/4] RUN make install                                                         33.1s
 => exporting to image                                                              2.3s
 => => exporting layers                                                             2.3s
 => => writing image sha256:f98072939f4707f5259495dc85dc2510febc8282372b3c2e2246fe  0.0s
 => => naming to docker.io/library/the-blockchain-bar_dev                           0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
docker-compose run \
                --rm -v "/Users/daniel/src/the-blockchain-bar":/build:consistent dev bash
Creating the-blockchain-bar_dev_run ... done

root@693ca11edb0f:/build#
```

Notice that when this command is completed, that your shell prompt looks slightly different:

```
root@693ca11edb0f:/build#
```

This is because you now have a shell _within_ the container. Let's try building `tbb` and seeing which version is running:

```
root@693ca11edb0f:/build# make install
go install -ldflags "-X main.GitCommit=c912bcf8a465bf81848b7789995a7f8e7261024c" ./...

root@693ca11edb0f:/build# tbb version
Version: 1.9.2-alpha c912bc TX Gas
```

That's pretty cool. Let's see how we can make a change to the code and see that change reflected inside the container.

First, open [cmd/tbb/version.go](./cmd/tbb/version.go) in your text editor of choice and change the following block:

```
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Describes version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Version: %s.%s.%s-alpha %s %s", Major, Minor, Fix, shortGitCommit(GitCommit), Verbal))
	},
}
```

to

```
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Describes version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("My Cool Version: %s.%s.%s-alpha %s %s", Major, Minor, Fix, shortGitCommit(GitCommit), Verbal))
	},
}

```

Now let's try building `tbb` again and checking the current version:

```
root@693ca11edb0f:/build# make install
go install -ldflags "-X main.GitCommit=c912bcf8a465bf81848b7789995a7f8e7261024c" ./...
root@693ca11edb0f:/build# tbb version
My Cool Version: 1.9.2-alpha c912bc TX Gas
```

Awesome! You can now write code using your preferred tooling and immediately see your changes reflected when you build and run the application within the container.

