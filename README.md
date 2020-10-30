# DataHow Challenge
DataHow Backend Interview - Coding Challenge

## Objectives achieved

* Creation of microservice that computes number of unique IP Addresses in a in-memory list of logs
* Insertion of new logs via correct port, simple JSON validation
* Retrieval of data from the correct port and address using the given standard

### Missing pieces

* Benchmark/Performance testing to produce optimized solution
* Usage of workers in Post request to enable async creating of metrics
* Performance improvements on the calculation of unique ip addresses (better space and time complexity)

### Major difficulties

* Adaption to new language and frameworks in short amount of time
* Instalation of benchmark performace tools

### Main Requirements

* Time limit : 3 hours
* Listen on ports :5000 and :9102
* Receive JSON logs on :5000/logs
* Server Prometheus metrics on :9102/metrics
* Compute number of unique IP addresses in logs since service start
* Create custom Prometheus metric "unique_ip_addresses"
* Publish in public Git Repository

### Not required

* Persistence
* Validate inputs

### Bonus

* Benchmark API using siege or gobench
* Preferred languages Go, Rust, Typescript
* Development using logical increments and document via commit messages
* Test driven development

### Evaluation Criteria

* Clean code
* Memory usage
* Performance
* Unit tests
* Git history and commit messages
