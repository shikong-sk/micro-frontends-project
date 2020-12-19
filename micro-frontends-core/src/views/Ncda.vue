<template>
	<div id="ncda" style="min-height: 100vh" ref="ncda"></div>
</template>

<script lang="ts">
    import {defineComponent, ref} from "vue";
    import useNavBar              from "@/components/nav-bar/useNavBar";

    export default defineComponent({
		name:"Ncda",
		setup(props,context){
			let ncda = ref();
			useNavBar(ncda);

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
