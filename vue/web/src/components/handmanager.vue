<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <input type="file" @change="fileChange">
    <button @click="upload">upload</button>
    <h1>{{formData}}</h1>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: 'handmanager',
  data () {
    return {
      msg: 'Welcome to PokerManager',
      formData: new FormData(),
    }
  },
  methods:{
    fileChange(e){
      this.formData.append('file', e.target.files[0])
      this.msg = this.formData.get('file')
    },
    upload() {
        axios.post('http://127.0.0.1', this.formData)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1{
  font-weight: normal;
}
</style>