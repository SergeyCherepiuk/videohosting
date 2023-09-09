<script lang="ts">
    import { onMount } from "svelte";

    export let preview: Blob
    export let video: Blob
    export let alt: string

    let videoSrc = URL.createObjectURL(video)
    let previewSrc: string | null = null
    
    onMount(() => {
        let canvas = document.getElementById("canvas") as HTMLCanvasElement
        let player = document.getElementById("player") as HTMLVideoElement
        canvas.getContext("2d")?.drawImage(player, 0, 0, player.videoWidth, player.videoHeight)
        canvas.toBlob(blob => { if (blob) preview = blob }) // ISSUE: Returns empty blob
    })

    let previewFile: FileList
    function onSelect() {
        preview = previewFile.item(0) as Blob
        setPreviewSrc(preview)
    }

    function setPreviewSrc(image: Blob) {
        previewSrc = URL.createObjectURL(image)
    }
</script>

<canvas id="canvas" class="hidden" />
<label for="preview" class="aspect-video rounded-xl overflow-clip bg-black">
    {#if previewSrc}
        <img id="image" src={previewSrc} {alt}
            class="w-full aspect-video rounded-xl bg-black object-contain" />
    {:else}
        <video id="player" src={videoSrc} />
    {/if}
</label>
<input id="preview" type="file" class="hidden" 
    bind:files={previewFile}  on:change={onSelect} />