pipeline {
    agent {
        docker {
            image 'golang:1.23'
            args '-v ${HOME}/.cache:/go/cache'
        }
    }
    
    environment {
        GOCACHE = '/go/cache'
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
    }
    
    stages {
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
        
        stage('Build') {
            steps {
                sh 'go build -v -o app ./cmd/app'
            }
        }
        
        stage('Build Docker Image') {
            agent any  // Переключаемся на хост Jenkins для работы с Docker
            steps {
                // Используем credentials для Docker registry если нужно
                // withCredentials([usernamePassword(credentialsId: 'docker-hub', passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
                sh 'docker build -t myapp:${BUILD_NUMBER} .'
                // Если нужно запушить образ
                // sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
                // sh 'docker push myapp:${BUILD_NUMBER}'
                // }
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}