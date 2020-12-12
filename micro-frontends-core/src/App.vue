<template>
	<div>
		<nav-bar v-model:collapse="collapse" style="z-index: 999999"></nav-bar>
		<router-view/>

		<div id="vue"></div>
		<div id="ncda" style="position: relative;border: #008ac7;min-height:100vh" ref="ncda"></div>
	</div>
</template>

<script lang="ts">
    import {defineComponent} from 'vue';
    import navBar            from "@/components/nav-bar/index.vue";

    export default defineComponent({
        name      : 'App',
        components: {navBar},
        props     : {
            msg: String,
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
        watch     : {
            collapse(v: boolean) {
                this.collapseChange(v);
            }
        },
        methods   : {
            collapseChange(collapse: boolean) {
                if (collapse) {
                    // @ts-ignore
                    this.$refs.ncda.style.width = "calc(100vw - 82px)";
                    // @ts-ignore
                    this.$refs.ncda.style.left = "56px";
                } else {
                    // @ts-ignore
                    this.$refs.ncda.style.width = "calc(100vw - 218px)";
                    // @ts-ignore
                    this.$refs.ncda.style.left = "200px";
                }
            }
        }
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
