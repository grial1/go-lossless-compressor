pipeline {
    
    agent any
    
    tools {
        
        go "Go 1.15" 

    }

    stages {
        
        stage ("build") {

            when {
                changeset "**/src/**"
            }
            
            steps {

                //git 'https://github.com/grial1/go-lossless-compressor'
                echo "Building project."
                sh "go build -o compressor.bin"
               
            }
            
        }
        stage ("test") {
        
            when {
                changeset "**/src/**"
                branch "master"
            }
        
            steps {
            
                echo "Testing project."
                sh "go test -v"
               
            }
            
        }        
        
    }
    post {
        
        success {
            
            echo "Build succeeded"
            slackSend (color: "#00FFFF",channel: "#instavot-cd", message: "Build Succeeded: ${env.JOB_NAME} ${env.BUILD_NUMBER}")
            archiveArtifacts artifacts:'*.bin', fingerprint: true

        }
        
        failure {
            
            echo "Build failed"
            slackSend (color: "#DC0000",channel: "#instavot-cd", message: "Build Failed: ${env.JOB_NAME} ${env.BUILD_NUMBER}")
        }
        
    }
    
}