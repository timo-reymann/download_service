#!groovy

node {
    properties([
        parameters([
             stringParam(name: 'tag', defaultValue: 'latest', description: 'Tag for build'),        
             booleanParam(name: 'release', defaultValue: true, description: 'Also tag as latest')
        ])
    ])

    def imageName = "timoreymann/download_service"
    def credentials = 'timoreymann-docker'
    def registry = 'https://registry.hub.docker.com'
    def app

    stage('Checkout') {
        checkout scm
    }

    stage('Build') {
        app = docker.build(imageName)
    }        
   
    stage("Publish") {
        docker.withRegistry(registry, credentials) {
            if(params.tag != 'latest' || !params.release) {
                app.push(params.tag)
            }
    
            if(params.release) {
                app.push("latest")
            }
        }
    }
}
