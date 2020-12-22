<template>
	<div ref="frame" style="height: 100vh">
		<iframe class="frame" src="http://192.168.3.2:9080"></iframe>
	</div>
</template>

<script lang="ts">
	import {defineComponent, ref} from "vue";
	import useNavBar              from "@/components/nav-bar/UseNavBar";

	export default defineComponent({
		name: "DockerRepositories",
		setup: (prop, context) => {
			let frame = ref(null)
			console.log(frame)

			useNavBar(frame)

			return {
				frame
			}
		}
	});
</script>

<style scoped>
	.frame {
		border: 0;
		width: 100%;
		height: 100%;
		overflow: hidden;
	}
</style>