<template>
  <div >
    <b-form-select v-on:change="analysis" v-model="selectplayer" :options="players"></b-form-select>  
    <table class="table">
      <thead>
        <th>PreFlop : {{this.preflopCount}}</th>
        <th>Flop : {{this.flopCount}}</th>
        <th>Turn : {{this.turnCount}}</th>
        <th>River : {{this.riverCount}}</th>
      </thead>
      <td>
        <ve-pie :data="preFlop" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
      <td>
        <ve-pie :data="Flop" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
      <td>
        <ve-pie :data="Turn" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
      <td>
        <ve-pie :data="River" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
    </table>
  </div>
</template>


<script>
export default {
  name: 'profit',
  data () {
    return {
      preFlop: {
        columns: ['action', 'count'],
        rows: [
          {
            'action':'Raise',
            'count':'0'
          },
          {
            'action':'Call',
            'count':'0'
          },
          {
            'action':'Fold',
            'count':'0'
          },
          {
            'action':'Check',
            'count':'0'
          },
          {
            'action':'Bet',
            'count':'0'
          }
        ],
      },
      Flop: {
        columns: ['action', 'count'],
        rows: [
          {
            'action':'Raise',
            'count':'0'
          },
          {
            'action':'Call',
            'count':'0'
          },
          {
            'action':'Fold',
            'count':'0'
          },
          {
            'action':'Check',
            'count':'0'
          },
          {
            'action':'Bet',
            'count':'0'
          }
        ],
      },
      Turn: {
        columns: ['action', 'count'],
        rows: [
          {
            'action':'Raise',
            'count':'0'
          },
          {
            'action':'Call',
            'count':'0'
          },
          {
            'action':'Fold',
            'count':'0'
          },
          {
            'action':'Check',
            'count':'0'
          },
          {
            'action':'Bet',
            'count':'0'
          }
        ],
      },
      River: {
        columns: ['action', 'count'],
        rows: [
          {
            'action':'Raise',
            'count':'0'
          },
          {
            'action':'Call',
            'count':'0'
          },
          {
            'action':'Fold',
            'count':'0'
          },
          {
            'action':'Check',
            'count':'0'
          },
          {
            'action':'Bet',
            'count':'0'
          }
        ],
      },
      charExtend:{
        legend:{
          left: 'center',
        },
        series:{
          color:[]
        }
      },
      preflopCount:0,
      flopCount:0,
      turnCount:0,
      riverCount:0,
      selectplayer:"Hero",
      players:[],
    }
  },
  methods:{
    analysis(){
      this.preflop()
      this.flop()
      this.turn()
      this.river()
    },
    preflop() {
      this.$http.get('http://'+this.$root.backIP+'/preflop',{params : {player: this.selectplayer}})
        .then( (response) => {
          this.preFlop.rows[0].count = response.data.Raise,
          this.preFlop.rows[1].count = response.data.Call,
          this.preFlop.rows[2].count = response.data.Fold,
          this.preFlop.rows[3].count = response.data.Check,
          this.preFlop.rows[4].count = response.data.Bet,
          this.preflopCount = parseInt(this.preFlop.rows[0].count) + parseInt(this.preFlop.rows[1].count) + parseInt(this.preFlop.rows[2].count) + parseInt(this.preFlop.rows[3].count) + parseInt(this.preFlop.rows[4].count)
        })
    },
    flop() {
      this.$http.get('http://'+this.$root.backIP+'/flop',{params : {player: this.selectplayer}})
        .then( (response) => {
          this.Flop.rows[0].count = response.data.Raise,
          this.Flop.rows[1].count = response.data.Call,
          this.Flop.rows[2].count = response.data.Fold,
          this.Flop.rows[3].count = response.data.Check,
          this.Flop.rows[4].count = response.data.Bet,
          this.flopCount = parseInt(this.Flop.rows[0].count) + parseInt(this.Flop.rows[1].count) + parseInt(this.Flop.rows[2].count) + parseInt(this.Flop.rows[3].count) + parseInt(this.Flop.rows[4].count)
        })
    },
    turn() {
      this.$http.get('http://'+this.$root.backIP+'/turn',{params : {player: this.selectplayer}})
        .then( (response) => {
          this.Turn.rows[0].count = response.data.Raise,
          this.Turn.rows[1].count = response.data.Call,
          this.Turn.rows[2].count = response.data.Fold,
          this.Turn.rows[3].count = response.data.Check,
          this.Turn.rows[4].count = response.data.Bet,
          this.turnCount = parseInt(this.Turn.rows[0].count) + parseInt(this.Turn.rows[1].count) + parseInt(this.Turn.rows[2].count) + parseInt(this.Turn.rows[3].count) + parseInt(this.Turn.rows[4].count)
        })
    },
    river() {
      this.$http.get('http://'+this.$root.backIP+'/river',{params : {player: this.selectplayer}})
        .then( (response) => {
          this.River.rows[0].count = response.data.Raise,
          this.River.rows[1].count = response.data.Call,
          this.River.rows[2].count = response.data.Fold,
          this.River.rows[3].count = response.data.Check,
          this.River.rows[4].count = response.data.Bet,
          this.riverCount = parseInt(this.River.rows[0].count) + parseInt(this.River.rows[1].count) + parseInt(this.River.rows[2].count) + parseInt(this.River.rows[3].count) + parseInt(this.River.rows[4].count)

        })
    },
    getPlayer(){
        this.$http.get('http://'+this.$root.backIP+'/player')
          .then( (response) => {
            response.data.forEach(element => {
              if (element == "Hero"){
                this.players.unshift(element)
              }else{
                this.players.push(element)
              }
            });
          })
    },
  },
  mounted(){
    this.getPlayer()
    this.analysis()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
table{
    table-layout : fixed;
    margin:auto;
}
</style>