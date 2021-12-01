<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <input type="file" @change="fileChange">
    <button @click="upload">upload</button>
    <form >
      比數:
      <select v-model="num">
              <option>1</option>
              <option>10</option>
              <option>100</option>
      </select>
      <button class="btn" @click="gethand">搜尋</button>
    </form>
    <table >
      <tr>
        <td>Time</td>
        <td>Player</td>
        <td>Seat</td>
        <td>Gain</td>
        <td>Card</td>
        <td>Preflop</td>
        <td>Flop</td>
        <td>Turn</td>
        <td>River</td>
      </tr>
      <tr v-for="t in table">
        <td>{{t.Time}}</td>
        <td>{{t.Player[0].Name}}</td>
        <td>{{t.Player[0].Seat}}</td>
        <td>{{t.Player[0].Gain}}</td>
        <img v-bind:src= "imgsrc + t.Player[0].Card[0].Num + t.Player[0].Card[0].Suit + '.png'">
        <img v-bind:src= "imgsrc + t.Player[0].Card[1].Num + t.Player[0].Card[1].Suit + '.png'">
        <!--td>{{t.Player[0].Card[0].Num + t.Player[0].Card[0].Suit +" "+ t.Player[0].Card[1].Num + t.Player[0].Card[1].Suit}}</td-->
        <td>{{t.Player[0].Action.Preflop}}</td>
        <td>{{t.Player[0].Action.Flop}}</td>
        <td>{{t.Player[0].Action.Turn}}</td>
        <td>{{t.Player[0].Action.River}}</td>
      </tr>
    </table>
  </div>
</template>

<script>

export default {
  name: 'handmanager',
  data () {
    return {
      msg: 'Welcome to PokerManager',
      formData: new FormData(),
      table: [],
      imgsrc :"../../static/images/",
      num:1
    }
  },
  methods:{
    fileChange(e){
      this.formData.append('file', e.target.files[0])
      
    },
    upload() {
        this.$http.put('http://127.0.0.1/hand/', this.formData)
          .then( (response) => {
            this.table = response.data
          })
    },

    gethand(){
      this.$http.get('http://127.0.0.1/hand/', {
        params: {
            num: this.num,
          }
      })
        .then((response) => {
            this.table = response.data
            console.log(this.table)
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