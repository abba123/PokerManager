<template>
  <div >
    <table class="table">
      <thead>
        <th>PreFlop</th>
        <th>Flop</th>
        <th>Turn</th>
        <th>Raise</th>
      </thead>
      <td>
        <ve-pie :data="preFlop" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
      <td>
        <ve-pie :data="preFlop" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
      <td>
        <ve-pie :data="preFlop" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
      </td>
      <td>
        <ve-pie :data="preFlop" width="100%" :extend="charExtend" style="margin: auto;"></ve-pie>
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
        columns: ['title', 'count'],
        rows: [
          {
            'title':'total',
            'count':'1'
          },
          {
            'title':'3Bet',
            'count':'1'
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
    threeBet() {
      this.$http.get('http://'+this.$root.backIP+'/threebet')
        .then( (response) => {
          this.preFlop.rows[0].count = response.data.Total,
          this.preFlop.rows[1].count = response.data.ThreeBet
        })
    },
  },
  mounted(){
    this.threeBet()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
td{
    width: 300px;
    margin:auto;
}
</style>