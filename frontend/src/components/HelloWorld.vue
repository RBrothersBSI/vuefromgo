<template>
  <div class="hello">
    <button @click="sendtobackend">Press Me</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'HelloWorld',
  props: {
  },
  methods: {
    sendtobackend: function(){
      axios.get('http://localhost:3000/api/hello').then(function(res){
        console.log(res)
      })
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
