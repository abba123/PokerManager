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
            <select v-model=player1.Card[0].Num>
              <option v-for="num in nums" :key=num>
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model=player1.Card[0].Suit>
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td>
            <select v-model=player1.Card[1].Num>
              <option v-for="num in nums" :key=num>
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model=player1.Card[1].Suit>
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <img v-bind:src="'../../static/images/'+player1.Card[0].Num+player1.Card[0].Suit+'.png'">
          <img v-bind:src="'../../static/images/'+player1.Card[1].Num+player1.Card[1].Suit+'.png'">
          <td id="result">
            {{winRate1}}
          </td>
        </tr>
        <tr>
          <td>Player2</td>
          <td>
            <select v-model=player2.Card[0].Num>
              <option v-for="num in nums" :key=num>
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model=player2.Card[0].Suit>
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <td>
            <select v-model=player2.Card[1].Num>
              <option v-for="num in nums" :key= num>
                {{ num }}
              </option>
            </select>
          </td>
          <td>
            <select v-model=player2.Card[1].Suit>
              <option v-for="suit in suits" :key="suit">
                {{ suit }}
              </option>
            </select>
          </td>
          <img v-bind:src="'../../static/images/'+player2.Card[0].Num+player2.Card[0].Suit+'.png'">
          <img v-bind:src="'../../static/images/'+player2.Card[1].Num+player2.Card[1].Suit+'.png'">
          <td id="result">
            {{winRate2}}
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
      player:[],
      player1:{
        Name:"1",
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
        Name:"123",
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
      winRate1: 0,
      winRate2: 0,
    }
  },
  methods:{
    getWinRate: function(){
      this.player = [this.player1,this.player2]
      console.log(this.player)
      this.$http
        .post('http://'+this.$root.backIP+'/getwinrate',this.player)
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