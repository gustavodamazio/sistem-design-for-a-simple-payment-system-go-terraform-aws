# terraform-lambda-go

Sample project to build lambda API with Go and terraform

## How to

The following explain the support command in this project in Make file

### How to clean

Run the following command to clean up the build binary

```shell
make clean
```

### How to build

Run the following command to generate go binary

```shell
make build
```

### How to init terraform

Run the following command to initialize terraform

> You require to run one time only

```shell
make init
```

### How to review terraform plan

Run the following command to review the changes that terraform going to apply to your infrastructure

```shell
make plan
```

### How to apply terraform changes

Run the following command to deploy your changes to your aws account

```shell
make apply
```

### How to test your api

After terraform apply successful, you will see `api_url` output to your console, that will be your api to test

### How to clean up your infrastructure

After testing complete, you can run the following command to clean up your lambda and api gateway

```shell
make destroy
```
