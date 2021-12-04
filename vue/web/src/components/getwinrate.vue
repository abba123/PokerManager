<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
      <button v-on:click="getWinRate">Calculate Win Rate</button>
      <table class="table">
        <thead>
          <th></th>
          <th colspan="2" >Card1</th>
          <th colspan="2" >Crad2</th>
          <th colspan="2" >Winrate</th>
        </thead>
        <tr>
          <td rowspan="2">Player1</td>
          <td>
            <select v-model="player1Card1Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player1Card1Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player1Card2Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player1Card2Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td id="result" colspan="2">
            {{winRate1}}
          </td>
        </tr>
        <tr>
          <td colspan="2"><img v-bind:src="'../../static/images/'+player1Card1Num+player1Card1Suit+'.png'"></td>
          <td colspan="2"><img v-bind:src="'../../static/images/'+player1Card2Num+player1Card2Suit+'.png'"></td>
        </tr>
        <tr>
          <td rowspan="2">Player2</td>
          <td>
            <select v-model="player2Card1Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player2Card1Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player2Card2Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player2Card2Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td id="result" colspan="2">
            {{winRate2}}
          </td>
        </tr>
        <tr>
          <td colspan="2"><img v-bind:src="'../../static/images/'+player2Card1Num+player2Card1Suit+'.png'"></td>
          <td colspan="2"><img v-bind:src="'../../static/images/'+player2Card2Num+player2Card2Suit+'.png'"></td>
        </tr>
      </table>
      
  </div>
</template>

<script>

export default {
  name: 'getwinrate',
  data () {
    return {
      counter : 1,
      msg: 'Welcome to PokerManager',
      suits: ["s","h","d","c"],
      nums:["1","2","3","4","5","6","7","8","9","10","11","12","13"],
      player1Card1Num:"",
      player1Card1Suit:"",
      player1Card2Num:"",
      player1Card2Suit:"",
      player2Card1Num:"",
      player2Card1Suit:"",
      player2Card2Num:"",
      player2Card2Suit:"",
      player1Card1: "",
      player1Card2: "",
      player2Card1: "",
      player2Card2: "",
      winRate1: 0,
      winRate2: 0,
    }
  },
  methods:{
    getWinRate: function(){
      this.$http
        .get('http://3.133.150.55/getwinrate/', {
          params: {
            name1: "player1",
            name2: "player2",
            p1Card1Num: this.player1Card1Num,
            p1Card1Suit: this.player1Card1Suit,
            p1Card2Num: this.player1Card2Num,
            p1Card2Suit: this.player1Card2Suit,
            p2Card1Num: this.player2Card1Num,
            p2Card1Suit: this.player2Card1Suit,
            p2Card2Num: this.player2Card2Num,
            p2Card2Suit: this.player2Card2Suit,
          }
        })
        .then( (response) => {
          this.winRate1 = response.data.player1
          this.winRate2 = response.data.player2
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
    width: 500px;
    margin:auto;
    margin-top: 30px;
}

select{
  width: 50px;
}

img{
  width: 50px;
}
</style>