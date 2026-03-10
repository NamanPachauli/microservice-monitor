# Microservice Monitoring System (Go + Docker + Kubernetes + Jenkins + Terraform)

##  Project Overview

This project is a **Microservice Monitoring System** built using **Golang**, containerized with **Docker**, deployed on **Kubernetes**, automated with **Jenkins**, and infrastructure managed using **Terraform**.

The goal of this project is to simulate a **DevOps-based microservice architecture** where services can be monitored and managed efficiently. The system demonstrates how modern DevOps tools work together in a real-world workflow.

This repository is designed as a **learning and portfolio project for DevOps and Backend Development**, showing practical integration between application code, containerization, orchestration, CI/CD pipelines, and infrastructure automation.

---

#  Tech Stack

### Backend

* Golang (Go)
* REST API

### Containerization

* Docker
* Docker Compose

### Orchestration

* Kubernetes

### CI/CD

* Jenkins Pipeline

### Infrastructure as Code

* Terraform

### Database

* MongoDB

---

#  Project Architecture

The architecture consists of a monitoring service that tracks and manages microservices running inside containers.

Main components:

1. **Monitoring Service (Golang)**
2. **Health Service**
3. **MongoDB Database**
4. **Docker Containers**
5. **Kubernetes Deployments**
6. **Jenkins Pipeline**
7. **Terraform Infrastructure**

```
                +--------------------+
                |      Jenkins       |
                |   CI/CD Pipeline   |
                +---------+----------+
                          |
                          v
                +--------------------+
                |   Docker Build     |
                |  Container Images  |
                +---------+----------+
                          |
                          v
                +--------------------+
                |     Kubernetes     |
                |  Deploy Services   |
                +---------+----------+
                          |
                          v
                +--------------------+
                | Monitoring Service |
                |       (Go)         |
                +---------+----------+
                          |
                          v
                    +-----------+
                    |  MongoDB  |
                    +-----------+
```

---

#  Project Structure

```
microservice-monitor
│
├── checker.go
├── handlers.go
├── models.go
├── mongo.go
├── registry.go
├── server.go
│
├── health.go
├── health.Dockerfile
│
├── docker-compose.yml
├── monitor.Dockerfile
│
├── monitor-deploy.yaml
├── mongo-deployment.yaml
├── payment-service.yaml
│
├── Jenkinsfile
│
├── terraform
│   ├── main.tf
│   ├── provider.tf
│   ├── variables.tf
│   └── outputs.tf
│
├── go.mod
├── go.sum
└── .dockerignore
```

---

#  Features

###  Microservice Monitoring

* Register services
* List active services
* Monitor service health

###  REST API

API endpoints allow interaction with the monitoring system.

Example endpoints:

```
GET /services
POST /register
GET /health
```

###  Containerization with Docker

The application runs inside Docker containers.

Dockerfiles are provided for:

* Monitoring service
* Health service

###  Kubernetes Deployment

Kubernetes manifests deploy services as:

* Deployments
* Services
* Pods

###  Jenkins CI/CD Pipeline

Jenkins automates:

* Code build
* Docker image creation
* Deployment

###  Terraform Infrastructure

Terraform manages infrastructure configuration and automation.

---

# 🛠 Installation & Setup

##  Clone Repository

```
git clone https://github.com/NamanPachauli/microservice-monitor.git
cd microservice-monitor
```

---

# Run Locally (Go)

```
go run server.go
```

Server runs on:

```
http://localhost:5001
```

---

#  Run with Docker

Build image:

```
docker build -f monitor.Dockerfile -t monitor-image .
```

Run container:

```
docker run -p 5001:5001 monitor-image
```

---

#  Run with Docker Compose

```
docker-compose up
```

This will start:

* Monitoring service
* MongoDB

---

#  Deploy with Kubernetes

Apply deployments:

```
kubectl apply -f mongo-deployment.yaml
kubectl apply -f monitor-deploy.yaml
kubectl apply -f payment-service.yaml
```

Check services:

```
kubectl get pods
kubectl get svc
```

---

#  Jenkins CI/CD Pipeline

The Jenkins pipeline automates:

1. Code checkout
2. Build Golang service
3. Build Docker image
4. Push image
5. Deploy application

Pipeline defined inside:

```
Jenkinsfile
```

---

#  Terraform Infrastructure

Initialize terraform:

```
terraform init
```

Plan infrastructure:

```
terraform plan
```

Apply configuration:

```
terraform apply
```

---

#  API Example

### Get Services

```
GET /services
```

Response

```
[
 { "id": 1, "name": "payment-service" }
]
```

---

#  Learning Goals

This project demonstrates:

* Building REST APIs with Go
* Containerizing applications with Docker
* Deploying microservices on Kubernetes
* Implementing CI/CD pipelines with Jenkins
* Infrastructure automation using Terraform
* Managing microservice architecture

---

#  Use Case

This project can be used as:

* DevOps portfolio project
* Backend microservice example
* Kubernetes deployment practice
* Terraform infrastructure learning
* CI/CD pipeline demonstration

---

#  Future Improvements

Possible improvements:

* Service discovery
* Prometheus monitoring
* Grafana dashboards
* Kubernetes Helm charts
* GitHub Actions CI/CD
* Load balancing
* Authentication & security

---

#  Author

**Naman Pachauli**

DevOps & Backend Enthusiast
Focused on learning **Cloud, Kubernetes, and Microservices Architecture**

GitHub

https://github.com/NamanPachauli

---

# ⭐ If you like this project

Give this repository a **star ⭐ on GitHub**.
