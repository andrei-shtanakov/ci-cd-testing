pipeline {
    agent {
        label 'node-1' // Замените 'node-1' на метку вашей ноды
    }
    
    environment {
        PATH = "$PATH:/usr/local/go/bin"
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
            steps {
                sh 'echo "Testing shell access"'
                sh '''
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