<script lang="ts">
    let files: FileList
    
    async function uploadFile() {
        let formData = new FormData()
        formData.append("file", files.item(0) as Blob)
        let response = await fetch("http://localhost:3000/upload", {
            method: "POST",
            headers: {
                "Accept": "text/event-stream",
                "Cache-Control": "no-cache",
                "Connection": "keep-alive",
            },
            body: formData,
        })

        let reader = response.body?.getReader()
        while (true) {
            let bytes = await reader?.read();
            if (bytes?.done) break;
            console.log(bytes?.value.toString() ?? "");
        }
    }
</script>

<input type="file" bind:files />
<button on:click={uploadFile}>Submit</button>