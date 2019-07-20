#!groovy
node {
    properties([
        parameters([
            gitTagVersionInput()
        ])
    ])

    runDefaultDockerPipeline currentBuild: currentBuild, imageName: "timoreymann/download_service"
}
