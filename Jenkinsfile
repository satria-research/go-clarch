pipeline {
    agent any
    environment{
        DOCKERHUB_CREDS = credentials('Dockerhub')
    }
    stages {
         stage('Clone Repo') {
             when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            environment {
                FILE_ENV_FROM_JENKINS = credentials('dev-env-satria')
            }
            steps {
                checkout scm
                sh '''#!/bin/bash
                addgroup jenkins docker
                docker ps
                cp -rf $FILE_ENV_FROM_JENKINS .env
                '''
            }
        }
        stage('Sonarqube Analysis') {
            steps {
                script {
                    scannerHome = tool 'jenkinsSonarScanner'
                }
                withSonarQubeEnv('brantasdua') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }
        stage("Quality Gate") {
            steps {
                timeout(time: 1, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }
        stage('Build Image') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
		         sh '''#!/bin/bash
                 docker build -t ubedev/brantas:14 .
                 '''
            }
        }
        stage('Docker Login') {
            steps {
                sh 'echo $DOCKERHUB_CREDS_PSW | docker login -u $DOCKERHUB_CREDS_USR --password-stdin'                
            }
         }
        stage('Docker Push Voyager') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {  
                sh 'docker push ubedev/brantas:14'
            }
         }
        stage('Send Discord Notif') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New brantas SATRIA pipeline triggered for $env.GIT_BRANCH", footer: 'Brantas Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
   }
    post {
		always {
			sh 'docker logout'
		}
	 }
    }