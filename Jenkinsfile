pipeline {
    agent {
        docker {
            image 'golang:1.23'
        }
    }
    
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
    }
    
    stages {
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        
        stage('Build') {
            steps {
                sh 'go build -o app ./cmd/app'
            }
        }
        
        stage('Build Docker Image') {
            agent any
            steps {
                script {
                    docker.build("myapp:${env.BUILD_ID}")
                }
            }
        }
        
        stage('Deploy') {
            agent any
            steps {
                echo 'Deploying...'
                // Здесь может быть деплой в k8s или на сервер
            }
        }
    }
    
    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}