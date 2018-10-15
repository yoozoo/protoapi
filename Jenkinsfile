node {
	ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}") {
		env.GOPATH=pwd()
		dir('src/github.com/yoozoo/protoapi') {
			stage('Checkout'){
				checkout scm
			}
			stage('Pre Build'){
				sh 'go get -u github.com/golang/dep/cmd/dep'
				sh '${GOPATH}/bin/dep ensure'
			}
			stage('Build'){
				sh """go build"""
			}
			stage('Test'){
				echo 'Testing'
					dir('test') {
						sh """mkdir result && ./protoapi.bats"""
					}
			}
		}
	}
}
