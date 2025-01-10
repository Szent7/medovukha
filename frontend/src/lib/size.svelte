<script lang="ts" context="module">
	export function SizeConvertToMegabyte(bytes: number): string {
		return (bytes / 1048576).toFixed(2);
	}
</script>
