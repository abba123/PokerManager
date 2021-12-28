<template>
  <div>
    <b-form-select v-on:change="profit" v-model="selectplayer" :options="players"></b-form-select>  
    <ve-line :data="chartData" :extend="extend" width="80%" style="margin: auto;"></ve-line> 
  </div>

    

</template>

<script>
export default {
  name: 'profit',
  data () {
    return {
      extend:{
        series:{
          smooth:false
        }
      },
      chartData: {
        columns: ['Hand', 'Gain'],
        rows: [
        ]
      },
      selectplayer:"Hero",
      players:[],
    }
  },
  methods:{
    profit() {
        this.chartData.rows = []
        this.$http.get('http://'+this.$root.backIP+'/profit',{params : {player: this.selectplayer}})
          .then( (response) => {
            response.data.forEach(element => {
              this.chartData.rows.push(element)
            });
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
    this.profit()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>