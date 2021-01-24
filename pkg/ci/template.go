package ci

// Base Jenkinsfile template
const jenkinsfileTemplate = `
#!/usr/bin/env groovy

pipeline {
	agent {
		label 'TO-DO'
	}

	options {

	}

	environment {
		VERSION = readFile('VERSION').trim()
	}

	stages {
		stage('TO-DO') {
			steps {

			}
		}

		stage('TO-DO') {
			steps {
				
			}
		}

		stage('TO-DO') {
			steps {
				
			}
		}

		stage('TO-DO') {
			steps {
				
			}
		}
	}

	post {
		always {
			sendNotifications(currentBuild.result, "TO-DO")
		}
	}
}
`
