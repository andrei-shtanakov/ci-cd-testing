pipeline {
    agent any
    
    stages {
        stage('Go Build') {
            agent {
                docker {
                    image 'golang:1.23'
                    args '-v ${HOME}/.cache:/go/cache'
                }
            }
            steps {
                sh 'go test -v ./...'
                sh 'go build -v -o app ./cmd/app'
            }
        }
        
        stage('Docker Build') {
            steps {
                sh "docker build -t myapp:${BUILD_NUMBER} ."
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
    }
}