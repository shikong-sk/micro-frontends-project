<template>
  <div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link> |
    <router-link to="/vue">子应用</router-link>
  </div>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
  </div>
</template>

<script lang="ts">
import { Options, Vue }     from 'vue-class-component';
import HelloWorld           from '@/components/HelloWorld.vue';
import {globalStateActions} from "@/main"; // @ is an alias to /src

@Options({
  components: {
    HelloWorld,
  },
})
export default class Home extends Vue {
  mounted(){
    globalStateActions.setGlobalState({"navIndex":"home"});
  }
}
</script>
