<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button>Calculate the winrate</button>
      <table class="table">
        <thead>
          <th></th>
          <th colspan="2" >Card1</th>
          <th colspan="2" >Crad2</th>
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
        </tr>
        <tr>
          <td colspan="2"><img v-bind:src="player1Card1"></td>
          <td colspan="2"><img v-bind:src="player1Card2"></td>
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
        </tr>
        <tr>
          <td colspan="2"><img v-bind:src="player2Card1"></td>
          <td colspan="2"><img v-bind:src="player2Card2"></td>
        </tr>
      </table>
      
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
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
    }
  },
  watch:{
    player1Card1Num: function(){
      this.player1Card1 = "../../static/images/"+this.player1Card1Num+this.player1Card1Suit+".png"
    },
    player1Card1Suit: function(){
      this.player1Card1 = "../../static/images/"+this.player1Card1Num+this.player1Card1Suit+".png"
    },
    player1Card2Num: function(){
      this.player1Card2 = "../../static/images/"+this.player1Card2Num+this.player1Card2Suit+".png"
    },
    player1Card2Suit: function(){
      this.player1Card2 = "../../static/images/"+this.player1Card2Num+this.player1Card2Suit+".png"
    },
    player2Card1Num: function(){
      this.player2Card1 = "../../static/images/"+this.player2Card1Num+this.player2Card1Suit+".png"
    },
    player2Card1Suit: function(){
      this.player2Card1 = "../../static/images/"+this.player2Card1Num+this.player2Card1Suit+".png"
    },
    player2Card2Num: function(){
      this.player2Card2 = "../../static/images/"+this.player2Card2Num+this.player2Card2Suit+".png"
    },
    player2Card2Suit: function(){
      this.player2Card2 = "../../static/images/"+this.player2Card2Num+this.player2Card2Suit+".png"
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
    width: 300px;
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