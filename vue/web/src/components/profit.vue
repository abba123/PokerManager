<template>
  <ve-line :data="chartData" :extend="extend" width="80%" style="margin: auto;"></ve-line>
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
      }
    }
  },
  methods:{
    profit() {
        this.$http.get('http://'+this.$root.backIP+'/profit')
          .then( (response) => {
            response.data.forEach(element => {
              this.chartData.rows.push(element)
            });
          })
    },
  },
  mounted(){
    this.profit()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>