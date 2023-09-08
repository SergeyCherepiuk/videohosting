<script lang="ts">
    import { uploadWithProgress } from "../../upload/upload";
    import DetailsWindow from "./DetailsWindow.svelte";
    import ProgressWindow from "./ProgressWindow.svelte";

    export let video: Blob
    export let fileName: string
    export let fileSize: string
    export let onClose: Function

    let title: string = ""
    let description: string = ""
    let preview: Blob
    let progress: number = 0

    // TODO: Complete preview generation from video blob
    let reader = new FileReader();
    reader.readAsDataURL(video);
    reader.onload = e => {
        if (e.target) {
            preview = new Blob([e.target.result as ArrayBuffer])
            console.log(preview)
        }
    }    

    enum Stage { Details, Progress }
    let stage: Stage = Stage.Details

    function upload() {
        stage = Stage.Progress
        uploadWithProgress(
            new Map<string, string | Blob>([
                ["title", title],
                ["description", description],
                ["video", video],
                ["preview", preview]
            ]),
            "http://localhost:3000/videos/upload",
            p => { progress = p }
        )
    }

    function close() {
        title = ""
        description = ""
        progress = 0
        onClose()
    }
</script>

{#if stage == Stage.Details}
    <DetailsWindow bind:title={title} bind:description={description}
        bind:preview={preview} onUpload={upload} onClose={close} />
{/if}

{#if stage == Stage.Progress}
    <ProgressWindow {fileName} {fileSize} {progress} onClose={close} />
{/if}