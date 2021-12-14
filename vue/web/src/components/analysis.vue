<template>
  <div >
    <table class="table">
      <thead>
        <th>PreFlop</th>
        <th>Flop</th>
        <th>Turn</th>
        <th>River</th>
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
        ]
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
        ]
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
        ]
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
        ]
      },
      charExtend:{
        legend:{
          bottom: '10%',
          left: 'center',
        },
        series:{
          color:[]
        }
      },
    }
  },
  methods:{
    preflop() {
      this.$http.get('http://'+this.$root.backIP+'/preflop')
        .then( (response) => {
          this.preFlop.rows[0].count = response.data.Raise,
          this.preFlop.rows[1].count = response.data.Call,
          this.preFlop.rows[2].count = response.data.Fold
        })
    },
    flop() {
      this.$http.get('http://'+this.$root.backIP+'/flop')
        .then( (response) => {
          this.Flop.rows[0].count = response.data.Raise,
          this.Flop.rows[1].count = response.data.Call,
          this.Flop.rows[2].count = response.data.Fold
        })
    },
    turn() {
      this.$http.get('http://'+this.$root.backIP+'/turn')
        .then( (response) => {
          this.Turn.rows[0].count = response.data.Raise,
          this.Turn.rows[1].count = response.data.Call,
          this.Turn.rows[2].count = response.data.Fold
        })
    },
    river() {
      this.$http.get('http://'+this.$root.backIP+'/river')
        .then( (response) => {
          this.River.rows[0].count = response.data.Raise,
          this.River.rows[1].count = response.data.Call,
          this.River.rows[2].count = response.data.Fold
        })
    },
  },
  mounted(){
    this.preflop(),
    this.flop(),
    this.turn(),
    this.river()
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