<template>
  <div id="app">
    <img src="./assets/logo.png">
    <aside>
      <router-link to = "/">{{msg}}</router-link>
      <router-link to = "getwinrate">WinRate</router-link>
      <router-link to = "handmanager">handmanager</router-link>
    </aside>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      msg: 'Welcome to PokerManager',
    }
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
</style>
