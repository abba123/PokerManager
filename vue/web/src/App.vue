<template>
  <div id="app">
    <img src="./assets/logo.png">
    <h1>{{msg}}</h1>
    <b-aside>
      <router-link to = "/">Home</router-link>
      <router-link to = "getwinrate">WinRate</router-link>
      <router-link to = "handmanager">Handmanager</router-link>
    </b-aside>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      msg: 'Welcome to PokerManager!!!!!!!!!??????',
    }
  },
  metaInfo: {
    meta: [
      { charset: 'utf-8' },
      { name: 'google-site-verification', content: '3kSBrI8vs9SEJ8nPH9KEUX-o53ffxs7TIwR4YPdkZ5I' },
    ]
  },
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
img{
  width: 100px;
}
</style>
