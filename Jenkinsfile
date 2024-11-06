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
            agent none 
            steps {
                sh '''
                    docker version
                    docker build -t myapp:${BUILD_NUMBER} .
                '''
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