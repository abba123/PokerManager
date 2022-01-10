<template>
  <div class="hello">
    <input type="file" @change="fileChange">
    <b-button @click="upload">upload</b-button>
    <b-form>
      筆數:
      <select v-model="num">
              <option>1</option>
              <option>10</option>
              <option>100</option>
      </select>
      Win/Lose:
      <select v-model="gain">
              <option>all</option>
              <option>>1.0</option>
              <option>>0.5</option>
              <option>>0.0</option>
              <option>>-0.5</option>
              <option>>-1.0</option>
      </select>
      位置:
      <select v-model="seat">
              <option>all</option>
              <option>HJ</option>
              <option>LJ</option>
              <option>CO</option>
              <option>BTN</option>
              <option>SB</option>
              <option>BB</option>
      </select>
      <b-button class="btn" @click="gethand">搜尋</b-button>
    </b-form>
    <table >
      <tr>
        <td width = "200px">Time</td>
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
        <td>{{t.Player[user].Seat}}</td>
        <td>{{t.Player[user].Gain}}</td>
        <img v-bind:src= "imgsrc + t.Player[user].Card[0].Num + t.Player[user].Card[0].Suit + '.png'">
        <img v-bind:src= "imgsrc + t.Player[user].Card[1].Num + t.Player[user].Card[1].Suit + '.png'">
        <!--td>{{t.Player[0].Card[0].Num + t.Player[0].Card[0].Suit +" "+ t.Player[0].Card[1].Num + t.Player[0].Card[1].Suit}}</td-->
        <td>{{t.Player[user].Action.Preflop}}</td>
        <td>{{t.Player[user].Action.Flop}}</td>
        <td>{{t.Player[user].Action.Turn}}</td>
        <td>{{t.Player[user].Action.River}}</td>
      </tr>
    </table>
    <b-modal v-model="insertHandShow">Insert Hand success</b-modal>
  </div>
</template>

<script>

export default {
  name: 'handmanager',
  data () {
    return {
      user:this.$root.user,
      msg: 'Welcome to PokerManager',
      formData: new FormData(),
      table: [],
      imgsrc :"../../static/images/",
      num:1,
      gain:"all",
      seat:"all",
      insertHandShow: false,
    }
  },
  methods:{
    fileChange(e){
      this.formData.append('file', e.target.files[0])
      
    },
    upload() {
        this.$http.post('http://'+this.$root.backIP+'/hand', this.formData)
          .then( (response) => {
            this.insertHandShow = true
            //this.gethand()
          })
    },

    gethand(){
      this.$http.get('http://'+this.$root.backIP+'/hand', {
        params: {
            num: this.num,
            gain: this.gain,
            seat: this.seat
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
    width: 800px;
    margin:auto;
    margin-top: 30px;
}
td{
   border-bottom:1px solid #915957;
}
img{
  width: 30px;
}
form, input{
  margin-top: 20px;
}
</style>