<template>
  <ve-line :data="chartData" :extend="extend" width="80%" style="margin: auto;"></ve-line>
</template>

<script>
export default {
  name: 'analysis',
  data () {
    return {
      extend:{
        series:{
          smooth:false
        }
      },
      chartData: {
        columns: ['Hands', 'Profit'],
        rows: [
        ]
      }
    }
  },
  methods:{
    analysis() {
        this.$http.get('http://'+this.$root.backIP+'/analysis')
          .then( (response) => {
            response.data.forEach(element => {
              this.chartData.rows.push(element)
            });
            console.log(response.data)
            console.log(this.chartData.rows)
          })
    },
  },
  mounted(){
    this.analysis()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>