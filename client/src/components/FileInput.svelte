<script lang="ts">
    import { parseEvent, type ServerSentEvent } from "../sse/sse";

    let files: FileList
    let progress: string = ""
    
    async function uploadFile() {
        let formData = new FormData()
        formData.append("file", files.item(0) as Blob)
        fetch("http://localhost:3000/upload", {
            method: "POST",
            headers: {
                "Accept": "text/event-stream",
                "Cache-Control": "no-cache",
                "Connection": "keep-alive",
            },
            body: formData,
        }).then(async (response) => {
            if (response.ok && response.body) {
                let reader = response.body.getReader()
                let decoder = new TextDecoder("utf-8");

                while (true) {
                    let { done, value } = await reader.read()
                    if (done) return
                    
                    let event: ServerSentEvent = parseEvent(decoder.decode(value))
                    let data: { progress: number } = JSON.parse(event.data)
                    progress = `${data.progress * 100}%`
                }   
            } 
        })
    }
</script>

<input type="file" name="file" bind:files />
<button on:click={uploadFile}>Submit</button>
<p>{progress}</p>