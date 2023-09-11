<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";
    import Button from "../common/Button.svelte";
    import Label from "../common/Label.svelte";
    import TextField from "../common/TextField.svelte";
    import PreviewPicker from "./PreviewPicker.svelte";
    import Header from "./Header.svelte";
    import Error from "../common/Error.svelte";

    export let title: string
    export let description: string
    export let preview: Blob
    export let video: Blob
    export let errorMessage: string | null
    export const validate: Function = () => {
        let titleErrorMessage = validateTitle();
        if (titleErrorMessage) return titleErrorMessage
        let descriptionErrorMessage = validateDescription();
        if (descriptionErrorMessage) return descriptionErrorMessage
        let previewErrorMessage = validatePreview();
        if (previewErrorMessage) return previewErrorMessage
        return null
    }
    export let onUpload: MouseEventHandler<HTMLButtonElement>
    export let onClose: MouseEventHandler<HTMLButtonElement>

    function validateTitle(): string | null {
        if (title.length == 0) return "Title is empty."
        if (title.length > 100) return "Title is too long."
        return null
    }

    function validateDescription(): string | null {
        if (description.length > 5000) {
            return "Description is too long. What are you describing the history of the universe?"
        }
        return null
    }

    function validatePreview(): string | null {
        if (preview.type.startsWith("image/")) {
            return "Please select an image for the preview or generate it."
        }
        return null
    }
</script>

<div class="flex flex-col gap-8 p-10 bg-white sm:rounded-2xl shadow-md w-full max-[640px]:h-full sm:w-3/4 xl:w-[960px]">
    <Header text="Upload new video" onClick={onClose} />
    <div class="flex flex-col w-full max-[900px]:gap-8">
        <div class="flex flex-row max-[900px]:flex-col-reverse w-full gap-8">
            <div class="flex-3">
                <Label label="Title">
                    <TextField name="title" placeholder="Name the video" bind:text={title} limit={100} />
                </Label>
            </div>
            <div class="flex-2">
                <Label label="Preview">
                    <PreviewPicker bind:preview={preview} {video} alt="Preview placeholder" />
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
    {#if errorMessage}
        <Error {errorMessage} />
    {/if}
    <div class="flex w-full justify-end">
        <Button text="Upload" isPrimary={true} onClick={onUpload} />
    </div>
</div>