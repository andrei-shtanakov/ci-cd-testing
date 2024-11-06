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
        stage('Build') {
            steps {
                // Показываем текущую директорию и её содержимое для диагностики
                sh 'pwd && ls -la'
                
                // Загружаем зависимости
                sh 'go mod tidy'
                
                // Собираем приложение
                sh 'go build -v -o app ./cmd/app'
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test ./...'
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