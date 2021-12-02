<template>
  <div class="hello">
    <form v-if="!this.$root.token">
      帳號:<input type="text" v-model="username"/>
      <br>
      密碼:<input type="password" v-model="password"/>
      <br>
      <button class="btn" @click="login">登入</button>
      <button class="btn" @click="register">註冊</button>
    </form>
    <div v-if="this.$root.token">
      登入成功
      <button class="btn" @click="logout">登出</button>
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
    }
  },
  methods:{
    login(){
      this.$http
        .post('http://127.0.0.1/', {
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
        .put('http://127.0.0.1/', {
          username: this.username,
          password: this.password,
          
        })
    },
    logout(){
      this.$http
        .delete('http://127.0.0.1/')
        
        this.$root.token = ""
        this.$http.defaults.headers.common['Authorization'] = this.$root.token
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1{
  font-weight: normal;
}
table{
    table-layout : fixed;
    margin:auto;
    margin-top: 30px;
}
</style>