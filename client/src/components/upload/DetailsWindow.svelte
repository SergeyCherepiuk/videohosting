<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";
    import Button from "../common/Button.svelte";
    import Label from "../common/Label.svelte";
    import TextField from "../common/TextField.svelte";
    import PreviewPicker from "./PreviewPicker.svelte";
    import Header from "./Header.svelte";

    export let title: string
    export let description: string
    export let preview: FileList
    export let clearPreview: Function
    export let onUpload: MouseEventHandler<HTMLButtonElement>
    export let onClose: MouseEventHandler<HTMLButtonElement>
</script>

<div class="flex flex-col gap-8 p-10 bg-white sm:rounded-2xl shadow-md w-full max-[640px]:h-full sm:w-3/4 xl:w-[960px]">
    <Header text="Upload new video" onClick={onClose} />
    <div class="flex flex-col w-full max-[900px]:gap-8">
        <div class="flex flex-row max-[900px]:flex-col-reverse w-full gap-8">
            <div class="flex-3">
                <Label label="Title">
                    <TextField name="title" placeholder="Name the video"
                        bind:text={title} limit={100} />
                </Label>
            </div>
            <div class="flex-2">
                <Label label="Preview">
                    <PreviewPicker src="images/preview_placeholder.svg" alt="Preview placeholder"
                        bind:preview={preview} bind:clearPreview={clearPreview} />
                </Label>
            </div>
        </div>
        <div class="w-full min-h-[192px]">
            <Label label="Description">
                <TextField name="description" placeholder="Tell viewers about your video"
                    bind:text={description} singleLine={false} limit={5000} minHeightPx={192} />
            </Label>
        </div>
    </div>
    <div class="flex w-full justify-end">
        <Button text="Upload" onClick={onUpload} />
    </div>
</div>