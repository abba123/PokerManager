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
    <b-modal v-model="registerModalShow">{{registerModalMsg}}</b-modal>
    <b-modal v-model="loginModalShow">{{loginModalMsg}}</b-modal>
    <b-modal v-model="logoutModalShow">logout success</b-modal>
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
      registerModalShow: false,
      registerModalMsg: "",
      loginModalShow: false,
      loginModalMsg: "",
      logoutModalShow: false
    }
  },
  methods:{
    login(){
      this.$http
        .post('http://'+this.$root.backIP+'/login', {
          username: this.username,
          password: this.password,
        })
        .then( (response) => {
          if (response.data){
            this.$root.user = this.username
            this.$root.token = response.data
            this.$http.defaults.headers.common['Authorization'] = this.$root.token 
            this.loginModalShow = true
            this.loginModalMsg = "login success"
          }
        })
        .catch( (error) =>{ // 请求失败处理
          this.loginModalShow = true
          this.loginModalMsg = "login fail"
        })
    },
    register(){
      this.$http
        .post('http://'+this.$root.backIP+'/register', {
          username: this.username,
          password: this.password,
        })
        .then( (response) => {
          this.registerModalShow = true
          this.registerModalMsg = "register success"
        })
        .catch( (error) => {
          this.registerModalShow = true
          this.registerModalMsg = "register fail : same user name"
        })
    },
    logout(){
      this.$http
        .delete('http://'+this.$root.backIP+'/logout')
        
      this.$root.user = ""
      this.$root.token = ""
      this.$http.defaults.headers.common['Authorization'] = this.$root.token
      this.logoutModalShow = true
    },
    oAuth(){
      this.$http
        .get('http://'+this.$root.backIP+'/oauth/access')
        .then( (response) => {
           window.open(response.data)
           this.checkOAuth()
        })
    },
    checkOAuth(){
      this.$http
        .get('http://'+this.$root.backIP+'/oauth/check')
        .then( (response) => {
          if(response.data){
            this.$root.user = response.data
            this.$root.token = response.data
            this.$http.defaults.headers.common['Authorization'] = this.$root.token
            this.loginModalShow = true
            this.loginModalMsg = "login success"
          }
        })
        .catch( (error) => {
          this.loginModalShow = true
          this.loginModalMsg = "login failed"
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>