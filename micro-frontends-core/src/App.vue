<template>
  <div>
    <div style="background:deepskyblue">微前端主框架</div>
    <nav-bar></nav-bar>
    <div id="vue"></div>
    <div id="ncda" style="position: relative;border: #008ac7"></div>
    <router-view/>
  </div>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import navBar from "@/components/nav-bar/index.vue";

  export default defineComponent({
    name: 'App',
    components:{navBar},
    props: {
      msg: String,
    },
  });
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
