pipeline {
  agent none
  stages {
    stage('pipeline') {
      steps {
        script {
          node {
            stage('delete old file'){
              dir("/home/ec2-user"){
                if(fileExists('Poker')){
                  sh "rm Poker"
                }
              }
            }
            stage('get clone'){
              dir('/home/ec2-user/Poker'){
                sh "git clone https://github.com/abba123/PokerManager"
              }
            }
            stage('build go'){
              dir('/home/ec2-user/Poker/PokerManager'){
                if(fileExists('output')){
                  sh "rm output"
                }
                sh "go build -o output"
              }
            }
            stage('run ansible'){
              dir('/home/ec2-user'){
                ansiblePlaybook(playbook: '/home/ec2-user/playbook.yml')
              }
            }
            stage('delete tmp'){
              dir('/home/ec2-user/Poker@tmp'){
                deleteDir()
              }
              dir('/home/ec2-user/Poker/PokerManager@tmp'){
                deleteDir()
              }
              dir('/home/ec2-user/Poker/PokerManager/vue/web@tmp'){
                deleteDir()
              }
            }
          }
        }

      }
    }

  }
}