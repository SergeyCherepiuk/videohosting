<script lang="ts">
    import Label from "../components/upload/Label.svelte";
    import TextField from "../components/upload/TextField.svelte";

    let title: string;
</script>

<Label label="Title">
    <TextField bind:text={title} name="title" limit={100} singleLine={false} />
</Label>