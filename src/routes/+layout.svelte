<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<div class="w-full p-8">
	{@render children()}
</div>
