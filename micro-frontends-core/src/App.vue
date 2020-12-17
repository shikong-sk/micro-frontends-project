<template>
		<nav-bar :collapse="collapse" style="z-index: 999999"></nav-bar>
		<router-view/>
</template>

<script lang="ts">
	import {defineComponent, ref} from 'vue';
	import navBar                 from "@/components/nav-bar/index.vue";
	import {useStore}             from "vuex";

	export default defineComponent({
		name: 'App',
		components: {navBar},
		props: {
			msg: String,
		},
		setup(){
			const collapseChange = (collapse: boolean) => {
				console.log("collapseChange",collapse)
			};
			return {
				$store:useStore(),
				collapseChange,
			}
		},
		data() {
			return {
				collapse: true,
			};
		},
		mounted() {
			console.log("MainApp", this);
			this.collapseChange(this.collapse);
			console.log("route", this.$route);
		},
		watch: {
			collapse(v: boolean) {
				console.log("collapse",v);
				this.collapseChange(v);
			},
		},
		methods: {}
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
