<script lang="ts">
    import TextField from "../components/upload/TextField.svelte";

    let title: string;
</script>

<TextField name="title" bind:text={title} limit={100} />