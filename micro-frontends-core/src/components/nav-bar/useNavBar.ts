import {inject, ref, watch, onMounted, Ref} from "vue";

function useNavBar(containerRef: Ref, offset: number = 0) {
	onMounted(() => {
		childAppOnMounted(containerRef, offset);
	});
}

function childAppOnMounted(containerRef: Ref, offset: number = 0) {
	let collapseNav = ref(inject("collapseNav"));
	containerRef.value.style.position = "absolute";
	collapseNavChange(containerRef, collapseNav.value,offset);

	watch(collapseNav, (val) => {
		collapseNavChange(containerRef, val);
	});
}

let collapseNavChange = (containerRef: Ref, val: unknown, offset: number = 0) => {
	console.log("containerRef", containerRef);
	if (val) {
		let wid = 65 + offset + "px";
		containerRef.value.style.width = `calc(100vw - ${wid})`;
		containerRef.value.style.left = "65px";
	} else {
		let wid = 200 + offset + "px";
		containerRef.value.style.width = `calc(100vw - ${wid})`;
		containerRef.value.style.left = "200px";
	}
};
export default useNavBar;
export {childAppOnMounted};