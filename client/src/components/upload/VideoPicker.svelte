<script lang="ts">
    export let video: Blob
    export let fileName: string
    export let fileSize: string
    export let text: string
    
    let videoFile: FileList
    function onSelect() {
        video = videoFile.item(0) as Blob
        fileName = videoFile.item(0)?.name ?? "unknown name"
        fileSize = videoFile.item(0)?.size?.toString() ?? "unknown size"
    }
</script>

<label for="video"
    class="text-lg text-white font-amazon-ember
        rounded-lg px-4 py-2 bg-blue-50 w-min">
    <p class="truncate">{text}</p>
</label>
<input id="video" type="file" class="hidden"
    bind:files={videoFile} on:change={onSelect} />