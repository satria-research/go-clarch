pipeline {
    agent any
    environment{
        DOCKERHUB_CREDS = credentials('Dockerhub')
    }
    stages {
        stage('Clone Repo Voyager') {
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
         stage('Clone Repo Pacific') {
             when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/main/pacific' }
                }
            }
            environment {
                FILE_ENV = credentials('dev-env-pacific')
            }
            steps {
                checkout scm
                sh '''#!/bin/bash
                addgroup jenkins docker
                docker ps
                echo "$FILE_ENV"
                cp -rf $FILE_ENV .env
                '''
            }
        }
        stage('Build Image Voyager') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
		         sh '''#!/bin/bash
                 docker build -t ubedev/brantas:12 .
                 '''
            }
        }
         stage('Build Image Pacific') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/main/pacific' }
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
                sh 'docker push ubedev/brantas:12'
            }
         }
         stage('Docker Push Pacific') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/main/pacific' }
                }
            }
            steps {  
                sh 'docker push ubedev/brantas:14'
            }
         }
        stage('Send Discord Notif Voyager') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New brantas VOYAGER pipeline triggered for $env.GIT_BRANCH", footer: 'Brantas Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
        stage('Send Discord Notif Pacific') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/main/pacific' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New brantas PACIFIC pipeline triggered for $env.GIT_BRANCH", footer: 'Brantas Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
   }
    post {
		always {
			sh 'docker logout'
		}
	 }
    }