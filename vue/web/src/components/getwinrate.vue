<template>
  <div class="hello">
      <b-button v-on:click="getWinRate">Calculate Win Rate</b-button>
      <table class="table">
        <thead>
          <th></th>
          <th colspan="2" >Card1</th>
          <th colspan="2" >Crad2</th>
          <th>Hand</th>
          <th>Winrate</th>
        </thead>
        <tr>
          <td>Player1</td>
          <td>
            <select v-model="player1.Card[0].Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player1.Card[0].Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player1.Card[1].Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player1.Card[1].Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <img v-bind:src="'../../static/images/'+player1.Card[0].Num+player1.Card[0].Suit+'.png'">
          <img v-bind:src="'../../static/images/'+player1.Card[1].Num+player1.Card[1].Suit+'.png'">
          <td id="result">
            {{winRate.Player1}}
          </td>
        </tr>
        <tr>
          <td>Player2</td>
          <td>
            <select v-model="player2.Card[0].Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player2.Card[0].Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player2.Card[1].Num">
              <option v-for="num in nums" :key="num">
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model="player2.Card[1].Suit">
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <img v-bind:src="'../../static/images/'+player2.Card[0].Num+player2.Card[0].Suit+'.png'">
          <img v-bind:src="'../../static/images/'+player2.Card[1].Num+player2.Card[1].Suit+'.png'">
          <td id="result">
            {{winRate.Player2}}
          </td>
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
      nums:[1,2,3,4,5,6,7,8,9,10,11,12,13],
      player:{
        player1:{},
        player2:{}
      },
      player1:{
        Name:"player1",
        Card:[
          {
            Num:0,
            Suit:""
          },
          {
            Num:0,
            Suit:""
          },
        ]
      },
      player2:{
        Name:"player2",
        Card:[
          {
            Num:0,
            Suit:""
          },
          {
            Num:0,
            Suit:""
          },
        ]
      },
      winRate:{
        Player1:0,
        Player2:0,
      }
    }
  },
  methods:{
    getWinRate: function(){
      this.player.player1 = this.player1
      this.player.player2 = this.player2

      this.player.player1.Card.forEach( value =>
        value.Num = parseInt(value.Num)
      )
      this.player.player2.Card.forEach( value =>
        value.Num = parseInt(value.Num)
      )
      
      this.$http
        .post('http://'+this.$root.backIP+'/getwinrate',this.player)
        .then( (response) => {
          this.winRate.Player1 = response.data.player1
          this.winRate.Player2 = response.data.player2
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
    width: 750px;
    margin:auto;
    margin-top: 30px;
}

select{
  width: 50px;
}

img{
  width: 40px;
}
</style>