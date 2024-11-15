## The RProxy Server

A Simple Proxy Server written in Go.

## It's part of a bigger project. Stay Tuned.


This project is part of a bigger project in undertaking. We are building an "Orchestration + Service Discovery" platform that combines the best of both orchestrators like "Kubernetes" and "Hashicorp Nomad" and Service Registries like "Hashicorp Consul".

This project is completely a fun project which aims to evolve into an all encompassing platform. We aim to reduce number of dependencies a Platform like Kubernetes needs to implement SOA (Service Oriented Architecture) and Microservice Architecture.

I and a few enthusiasts [@HarshSharma0801](https://github.com/HarshSharma0801) are building it. Stay Tuned.

Why have two separate platforms for "Container Orchestration" and a "Service Mesh" for more sophisticated Service Orchestration ? 
This is what we intend to solve

Also while we are container first in our approach we want to build this platform to include "VM Workloads" and have special constructs for Stateful Applications like databases.


## What is yet to be implemented here

- [X] Add a rudimentary Reverse Proxy. [Boilerplate for further development. It's mostly copied code with enough understanding].
- [ ] Add Support to run as a SystemD Unit.
- [ ] Add Host based Routing.
- [ ] Implement L4/7 Load Balancing.
- [ ] Add SSL/TLS Termination Support.
- [ ] Integrate with a Service Registry [WIP.]
- [ ] Add gRPC and Websocket.

## Running RProxy

### Step1: Configure
- Create a configuration file `config.yml`.
- Copy this file either to `$(pwd)/data/config` or to `/etc/RPServer/data`
- The Config file:  
```yaml
gateway:
  host: localhost
  listen_port: "8080"
  scheme: http

resources:
  - name: Serv1
    endpoint: /server1
    upstream_url: "http://localhost:9001"
  - name: Serv2
    endpoint: /server2
    upstream_url: "http://localhost:9002"
  - name: Serv3
    endpoint: /server3
    upstream_url: "http://localhost:9003"
```

### Running without build
```
make run-server
```

### Building and Running
- Build the executable
```
OS_ARCH=amd64 #change to arm64 for ARM processors
make GOARCH=amd64  all # defaults to amd64 if not set
cd build/
```  
- In the next step setup configuration as shown above. Remember if configuration is not in `/etc/RPServer/data` then it should reside in same directory as the executable.

- Run the executable

__For Mac and Linux__

```
cd build/
export OS_NAME=linux #for mac change to darwin
export OS_ARCH=amd64 #for m1 mac set to arm64 else set to amd64
./rpserve-$OS_NAME-$OS_ARCH
```  
  
__For Windows__ [[@harsh098](https://github.com/harsh098) **has no intention to support it.**]  

```
cd build
start rpserve-win64-amd64.exe #for arm64 change to rpserve-win64-arm64.exe
```

### Cleanup After build
```
  make clean
```

# Setup a Dev Environment
- Clone this repo
- Navigate to the root of the cloned repo
- run:
  ```
    make run-dev
  ```
- To stop containers
  ```
    make stop
  ```

# Acknowledgement
The boilerplate is heavily borrowed from [https://prabeshthapa.medium.com/learn-reverse-proxy-by-creating-one-yourself-using-go-87be2a29d1e](https://prabeshthapa.medium.com/learn-reverse-proxy-by-creating-one-yourself-using-go-87be2a29d1e) and can be found at [pgaijin66/tinyrp](https://github.com/pgaijin66/tinyrp).
