<template>
  <div class="hello">
    <form>
      帳號:<input type="text" v-model="username" />
      <br>
      密碼:<input type="password" v-model="password" />
      <br>
      <button class="btn" @click="submit">確認</button>
    </form>
  </div>
</template>


<script>


export default {
  name: 'login',
  data () {
    return {
      msg: 'Welcome to PokerManager',
      username:"456",
      password:"",
    }
  },
  methods:{
    submit(){
      this.$http
        .post('http://127.0.0.1/', {
          params: {
            username: this.username,
            password: this.password,
          }
        })
        .then( (response) => {
          if (response.data){
            this.$root.token = response.data
            this.$http.defaults.headers.common['Authorization'] = this.$root.token 
          }
        })
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
img{
  width: 30px;
}
</style>