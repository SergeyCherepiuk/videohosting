<script lang="ts">
    import Button from "../common/Button.svelte";

    export let preview: Blob | null = null
    export let video: Blob
    export let alt: string 

    let previewFile: FileList
    function onSelect() {
        preview = previewFile.item(0) as Blob
    }

    function generatePreview() {
        let videoElement = document.getElementById("player") as HTMLVideoElement
        let canvasElement = document.getElementById("canvas") as HTMLCanvasElement
        let ctx = canvasElement.getContext("2d")
        ctx?.drawImage(videoElement, 0, 0, canvasElement.width, canvasElement.height)
        canvasElement.toBlob(blob => { if (blob) preview = blob })
    }

    function clearPreview() {
        preview = null
    }
</script>

{#if preview}
    <button on:click={clearPreview}><span class="text-4xl">&times;</span></button>
{/if}

<label for="preview" class="flex aspect-video rounded-xl overflow-clip bg-black justify-center items-center">
    {#if preview}
        <img id="image" src={URL.createObjectURL(preview)} {alt}
            class="w-full aspect-video rounded-xl bg-black object-contain" /> 
    {:else}
        <div class="flex flex-col items-center gap-4">
            <Button text="Generate preview" isPrimary={false} onClick={generatePreview} />
            <p class="text-white text-lg">Or click outside to <u>choose</u></p>
        </div>
    {/if} 

    <video id="player" src={URL.createObjectURL(video)} class="w-0">
        <track kind="captions" />
    </video>
    <canvas id="canvas" class="hidden" />
</label>
<input id="preview" type="file" class="hidden" 
    bind:files={previewFile}  on:change={onSelect} />