<script lang="ts">
    import TextField from "../components/upload/TextField.svelte";

    let title: string;
</script>

<TextField name="title" label="Title" bind:text={title} limit={100} singleLine={false} />