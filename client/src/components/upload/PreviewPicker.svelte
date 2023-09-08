<script lang="ts">
    export let src: string
    export let alt: string
    export let preview: FileList
    export let clearPreview: Function
    let previewSrc: string | null = null

    clearPreview = () => {
        let el: HTMLInputElement | null = document.getElementById("preview") as HTMLInputElement
        if (el) { el.value = "" }
    }

    function onSelect() { 
        let reader = new FileReader();
        reader.readAsDataURL(preview.item(0) as File)
        reader.onload = e => {
            previewSrc = e.target?.result as string
        }
    }
</script>

<label for="preview">
    <img id="image" src={previewSrc ? previewSrc : src} {alt}
        class="w-full aspect-video rounded-xl bg-black object-contain" />
</label>
<input id="preview" type="file" class="hidden" 
    bind:files={preview}  on:change={onSelect} />