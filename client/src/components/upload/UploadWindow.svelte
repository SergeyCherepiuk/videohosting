<script lang="ts">
    import Button from "../common/Button.svelte";
    import Label from "../common/Label.svelte";
    import Preview from "../common/Preview.svelte";
    import TextField from "../common/TextField.svelte";

    let title: string
    let description: string
</script>

<div class="flex flex-col items-end gap-4 p-10 rounded-2xl shadow-md w-3/4 xl:w-[960px]">
    <div class="flex flex-row max-[900px]:flex-col-reverse gap-12 w-full">
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
    <Button text="Upload" onClick={() => {}} />
</div>