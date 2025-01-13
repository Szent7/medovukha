<script lang="ts" context="module">
	export function SizeConvert(bytes: number): string {
		if (bytes === 0) return '0 Bytes';

		const k = 1024;
		const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
		const i = Math.floor(Math.log(bytes) / Math.log(k));
		console.log('Bytes:' + bytes);
		console.log('Math log bytes:' + Math.log(bytes));
		console.log('Math log k(1024):' + Math.log(k));
		console.log('i:' + i);
		console.log('Math pow k i:' + Math.pow(k, i));
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	}
</script>
