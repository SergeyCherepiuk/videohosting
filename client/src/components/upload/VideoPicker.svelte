<script lang="ts">
    export let video: Blob
    export let fileName: string
    export let fileSize: string
    export let text: string
    
    let videoFile: FileList
    function onChange() {
        video = videoFile.item(0) as Blob
        if (isContentTypeValid(video)) {
            fileName = video.name ?? "unknown name"
            fileSize = video.size?.toString() ?? "unknown size"
        }
    }

    function isContentTypeValid(file: Blob): boolean {
        return file.type.startsWith("video/")
    }
</script>

<label for="video" class="text-lg text-white font-amazon-ember rounded-lg px-4 py-2 bg-blue-50 w-min">
    <p class="truncate select-none">{text}</p>
</label>
<input id="video" type="file" accept="video/*" class="hidden" bind:files={videoFile} on:change={onChange} />