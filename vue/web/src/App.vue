<template>
  <b-card title="Card Title" no-body id="app">
    <b-card-header header-tag="nav">
      <b-nav card-header tabs>
        <!-- <b-nav-item>'s with child routes. Note the trailing slash on the first <b-nav-item> -->
        <b-nav-item to="/" exact exact-active-class="active">Home</b-nav-item>
        <b-nav-item to="/getwinrate" exact exact-active-class="active">getwinrate</b-nav-item>
        <b-nav-item to="/handmanager" exact exact-active-class="active">handmanager</b-nav-item>
        <b-nav-item to="/analysis" exact exact-active-class="active">analysis</b-nav-item>
      </b-nav>
    </b-card-header>

    <b-card-body>
      <router-view></router-view>
    </b-card-body>
  </b-card>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      msg: 'Welcome to PokerManager!',
    }
  }
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
