<template>
		<nav-bar v-model:collapse="collapse"
		         :collapse="collapse"
		         style="z-index: 999999">
		</nav-bar>

		<router-view/>
</template>

<script lang="ts">
	import {defineComponent, ref,watch,provide} from 'vue';
	import navBar                 from "@/components/nav-bar/index.vue";
	import {useStore}             from "vuex";

	export default defineComponent({
		name: 'App',
		components: {navBar},
		props: {
			msg: String,
		},
		setup(){
			let collapse = ref<boolean>(true);
			provide("collapseNav",collapse);
			const collapseChange = (collapse: boolean) => {
				console.log("collapseChange",collapse)
			};
			return {
				collapse,
				$store:useStore(),
				collapseChange,
			}
		},
		mounted() {
			console.log("MainApp", this);
			this.collapseChange(this.collapse);
			console.log("route", this.$route);
		},
		watch: {
			collapse(v: boolean) {
				this.collapseChange(v);
			},
		},
	});
</script>

<style>
	html,body{
		margin: 0;
		padding: 0;
		border: 0;
	}

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
