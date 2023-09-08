<script lang="ts">
    import { uploadWithProgress } from "../../upload/upload";
    import DetailsWindow from "./DetailsWindow.svelte";
    import ProgressWindow from "./ProgressWindow.svelte";

    export let video: FileList
    export let clearVideo: Function
    export let onClose: Function

    let title: string = ""
    let description: string = ""
    let preview: FileList
    let clearPreview: Function 
    let progress: number = 0

    enum Stage { Details, Progress }
    let stage: Stage = Stage.Details

    function upload() {
        stage = Stage.Progress
        uploadWithProgress(
            new Map<string, string | Blob>([
                ["title", title],
                ["description", description],
                ["video", video.item(0) as Blob],
                ["preview", preview.item(0) as Blob]
            ]),
            "http://localhost:3000/videos/upload",
            p => { progress = p }
        )
    }

    function close() {
        title = ""
        description = ""
        progress = 0 
        clearPreview()
        clearVideo()
        onClose()
    }
</script>

{#if stage == Stage.Details}
    <DetailsWindow bind:title={title} bind:description={description}
        bind:preview={preview} bind:clearPreview={clearPreview}
        onUpload={upload} onClose={close} />
{/if}

{#if stage == Stage.Progress}
    <ProgressWindow fileName={video.item(0)?.name ?? "unknown filename"} 
        fileSize={(video.item(0)?.size ?? 0).toString()} 
        {progress} onClose={close} />
{/if}