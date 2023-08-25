<script lang="ts">
    let files: FileList

    function onSubmit() {
        let fileReader = new FileReader()
        fileReader.readAsArrayBuffer(files.item(0) as Blob)
        fileReader.onload = async () => {
            let file = fileReader.result

            // TODO: Split into chunks
            await fetch("http://localhost:3000/upload", {
                "method": "POST",
                "headers": {
                    "content-type": "application/octet-stream",
                    "content-length": (file as ArrayBuffer).byteLength.toString(),
                },
                "body": file
            })
        }
    }
</script>

<input type="file" bind:files />
<button on:click={onSubmit}>Submit</button>