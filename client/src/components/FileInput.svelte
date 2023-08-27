<script lang="ts">
    let files: FileList
    let progress: string = ""
    
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

        try {
            const reader = response.body?.getReader()
            if (!reader) return

            while (true) {
                const { done, value } = await reader.read()
                if (done) break

                if (value) {
                    console.log(value)
                    progress = new TextDecoder().decode(value); 
                }
            }
        } catch (error) {
            console.error("An error occurred:", error);
        }
    }
</script>

<input type="file" name="file" bind:files />
<button on:click={uploadFile}>Submit</button>
<p>{progress}</p>