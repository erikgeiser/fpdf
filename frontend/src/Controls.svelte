<script lang="ts">
  export let pdfFiles: File[] = [];

  let fileNames: string[] = [];

  function onDrop(e: DragEvent) {
    e.preventDefault();

    for (var i = 0; i < e.dataTransfer.files.length; i++) {
      let file = e.dataTransfer.files[i];
      fileNames = [...fileNames, file.name];
      pdfFiles = [...pdfFiles, file];
    }

    return false;
  }

  function onDragEnter(e: DragEvent) {
    e.dataTransfer.dropEffect = 'move';
    e.preventDefault();
    return false;
  }

  function onDragOver(e: DragEvent) {
    e.dataTransfer.dropEffect = 'move';
    e.preventDefault();
    return false;
  }
</script>

<style>
  main {
  }
  .dropzone {
    height: 293px;
    background: #181719;
    overflow: auto;
    margin: 10px;
    border-radius: 5px;
    
  }
  .items {
    margin: 5px;
    background: #2b2b2b;
    list-style: none;
    font-size: 20px;
    color: white;
    font-family: sans-serif;
    padding: 10px;
    z-index: 999999;
    position: relative;
  }
  .controls{
    height:50px;
    width:100%;
    background: #181719;
  }
  .dropzoneText{
    position: absolute;
    max-width: 100%;
    height: auto;
    top: 200px;
    transform: translate(-50%, -50%);
    left: 50%;
    color: white;
    font-size: 30px;
    font-family:sans-serif;
  }
</style>

<main>
  <div class="controls"></div>

  <div
    class="dropzone"
    on:drop={onDrop}
    on:dragenter={onDragEnter}
    on:dragover={onDragOver}>

    <div class="dropzoneText">DROP IT LIKE IT'S HOT!</div>

    <div class="itemsUL">
      {#each fileNames as fileName, i}
        <li class="items">{i} {fileName}</li>
      {/each}
    </div>

  </div>
</main>
