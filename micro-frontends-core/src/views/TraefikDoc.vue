<template>
	<div ref="frame" style="height: 100vh">
		<iframe class="frame" src="https://traefik.tech/"></iframe>
	</div>
</template>

<script lang="ts">
	import {defineComponent, ref} from "vue";
	import useNavBar              from "@/components/nav-bar/UseNavBar";

	export default defineComponent({
		name: "TraefikDoc",
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