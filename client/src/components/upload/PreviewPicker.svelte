<script lang="ts">
    export let preview: Blob
    export let alt: string

    let previewFile: FileList
    function onSelect() {
        preview = previewFile.item(0) as Blob
        let el: HTMLImageElement = document.getElementById("image") as HTMLImageElement
        el.src = URL.createObjectURL(preview)
    }
</script>

<label for="preview">
    <img id="image" {alt}
        class="w-full aspect-video rounded-xl bg-black object-contain" />
</label>
<input id="preview" type="file" class="hidden" 
    bind:files={previewFile}  on:change={onSelect} />