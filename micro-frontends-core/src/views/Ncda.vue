<template>
	<div id="ncda" style="min-height: 100vh" ref="ncda"></div>
</template>

<script lang="ts">
    import {defineComponent, ref} from "vue";
    import UseNavBar              from "@/components/nav-bar/UseNavBar";

    export default defineComponent({
		name:"Ncda",
		setup(props,context){
			let ncda = ref();
			UseNavBar(ncda);

			return {
				ncda
			}
		},
	})
</script>

<style scoped>
	div#ncda > ::v-deep(div)  {
		min-height: 100vh;
	}
</style>
