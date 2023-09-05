<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";
    import Button from "../common/Button.svelte";
    import Label from "../common/Label.svelte";
    import Preview from "../common/Preview.svelte";
    import TextField from "../common/TextField.svelte";

    export let title: string
    export let description: string
    export let onUpload: MouseEventHandler<HTMLButtonElement>
    export let onClose: MouseEventHandler<HTMLButtonElement>
</script>

<div class="flex flex-col gap-8 p-10 bg-white rounded-2xl shadow-md w-full sm:w-3/4 xl:w-[960px]">
    <div class="flex flex-row w-full justify-between items-center">
        <span class="text-3xl font-amazon-ember ">Upload new video</span>
        <button on:click={onClose}><span class="text-4xl">&times;</span></button>
    </div>
    <div class="flex flex-col w-full max-[900px]:gap-8">
        <div class="flex flex-row max-[900px]:flex-col-reverse w-full gap-8">
            <div class="flex-3">
                <Label label="Title">
                    <TextField 
                        name="title" placeholder="Name the video"
                        bind:text={title} limit={100} />
                </Label>
            </div>
            <div class="flex-2">
                <Label label="Preview">
                    <Preview src="images/preview_placeholder.svg" alt="Preview placeholder" />
                </Label>
            </div>
        </div>
        <div class="w-full min-h-[192px]">
            <Label label="Description">
                <TextField 
                    name="description" placeholder="Tell viewers about your video"
                    bind:text={description} singleLine={false} limit={5000} minHeightPx={192} />
            </Label>
        </div>
    </div>
    <div class="flex w-full justify-end">
        <Button text="Upload" onClick={onUpload} />
    </div>
</div>