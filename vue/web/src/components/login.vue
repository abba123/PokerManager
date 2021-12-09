<template>
  <div class="hello" size="10px">
    <b-form v-if="!this.$root.token" style="width:30%; margin: auto;">
      <b-form-group v-if="!this.$root.token" label="username">
        <b-form-input v-model="username" required></b-form-input>
      </b-form-group>

      <b-form-group v-if="!this.$root.token" label="password">
        <b-form-input v-model="password" required></b-form-input>
      </b-form-group>
      <b-button class="btn" @click="oAuth">github</b-button>
      <b-button variant="primary" class="btn" @click="login">登入</b-button>
      <b-button variant="primary" class="btn" @click="register">註冊</b-button>
    </b-form>
    <div v-if="this.$root.token">
      <b-button variant="danger" class="btn" @click="logout">登出</b-button>
    </div>
  </div>
</template>


<script>


export default {
  name: 'login',
  data () {
    return {
      msg: 'Welcome to PokerManager',
      username:"test",
      password:"test",
      timer: '',
      timecount :0,
    }
  },
  methods:{
    login(){
      this.$http
        .post('http://'+this.$root.backIP, {
          username: this.username,
          password: this.password,
          
        })
        .then( (response) => {
          if (response.data){
            this.$root.token = response.data
            this.$http.defaults.headers.common['Authorization'] = this.$root.token 
          }
        })
    },
    register(){
      this.$http
        .put('http://'+this.$root.backIP, {
          username: this.username,
          password: this.password,
          
        })
    },
    logout(){
      this.$http
        .delete('http://'+this.$root.backIP)
        
        this.$root.token = ""
        this.$http.defaults.headers.common['Authorization'] = this.$root.token
    },
    oAuth(){
      this.$http
        .get('http://'+this.$root.backIP+'/oauth/access')
        .then( (response) => {
           //console.log(this.$http.defaults.headers.common)
           //this.$http.get(response.data)
           window.open(response.data)
           this.timeInterval = setInterval(this.checkOAuth, 5000);
        })
    },
    checkOAuth(){
      if (this.timecount < 10){
        this.timecount += 1
        this.$http
          .get('http://'+this.$root.backIP+'/oauth/check')
          .then( (response) => {
            if(response.data){
             this.$root.token = response.data
             this.$http.defaults.headers.common['Authorization'] = this.$root.token
              clearInterval(this.timeInterval);
            }
          })
      }else{
        this.timecount = 0
        clearInterval(this.timeInterval);
      }
      
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>