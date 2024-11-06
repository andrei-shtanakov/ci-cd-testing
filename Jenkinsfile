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
        DOCKER_BUILDKIT = '1'
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
            steps {
                script {
                    // Используем встроенный Docker
                    def customImage = docker.build("myapp:${env.BUILD_ID}", ".")
                    
                    // Если нужно, можно запушить образ
                    // customImage.push()
                }
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