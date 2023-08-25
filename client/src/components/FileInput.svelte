<script lang="ts">
    let files: FileList
    let onSubmitPromise: Promise<Response>

    async function onSubmit() {
        const formData = new FormData()
        formData.append("file", files.item(0) as Blob)
        return await fetch("http://localhost:3000/upload", {
            "method": "POST",
            "body": formData 
        })
    }
</script>

<input type="file" bind:files />
<button on:click={() => onSubmitPromise = onSubmit()}>Submit</button>
{#if onSubmitPromise}
    {#await onSubmitPromise}
        <p>Sending file...</p> 
    {:then response} 
        {#if response.ok}
            <p>File has been sent successfully!</p>
        {:else}
            <p>The error ocurred.</p>
        {/if}
    {:catch}
        <p>File is too large!</p>
    {/await}
{/if}