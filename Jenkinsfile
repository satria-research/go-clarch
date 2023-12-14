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
                FILE_ENV = credentials('dev-env')
            }
            steps {
                checkout scm
                sh '''#!/bin/bash
                addgroup jenkins docker
                docker ps
                cp -rf $FILE_ENV .env
                '''
            }
        }
        stage('Build Image') {
            steps {
		         sh '''#!/bin/bash
                 docker build -t ubedev/brantas:$BUILD_NUMBER .
                 '''
            }
        }
        stage('Docker Login') {
            steps {
                sh 'echo $DOCKERHUB_CREDS_PSW | docker login -u $DOCKERHUB_CREDS_USR --password-stdin'                
            }
         }
        stage('Docker Push') {
            steps {  
                sh 'docker push ubedev/brantas:$BUILD_NUMBER'
            }
         }
        stage('Send Discord Notif') {
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New brantas pipeline triggered for $env.GIT_BRANCH", footer: 'Brantas Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
   }
    post {
		always {
			sh 'docker logout'
		}
	 }
    }