<script lang="ts">
    import Label from "../common/Label.svelte";
    import Preview from "../common/Preview.svelte";
    import TextField from "../common/TextField.svelte";

    let title: string
    let description: string
    $: console.log(title, description)
</script>

<div class="flex flex-col">
    <div class="flex flex-row gap-8">
        <div class="flex-3">
            <Label label="Title">
                <TextField name="title" bind:text={title} limit={100} />
            </Label>
        </div>
        <div class="flex-2">
            <Label label="Preview">
                <Preview src="images/preview_placeholder.svg" alt="Preview placeholder" />
            </Label>
        </div>
    </div>
    <Label label="Description">
        <TextField name="description" bind:text={description} singleLine={false} limit={5000} minHeightPx={192} />
    </Label>
</div>