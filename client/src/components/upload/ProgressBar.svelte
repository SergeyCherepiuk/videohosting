<script lang="ts">
    export let progress: number
</script>

<div class="flex flex-row w-full h-1.5 rounded-full overflow-hidden bg-gray-70">
    <div class="bg-blue-50" style="width: {progress}%"></div>
</div>