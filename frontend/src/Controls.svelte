<script lang="ts">
  let fileNames: string[] = [];

  function onDrop(e: DragEvent) {
    e.preventDefault();

    for (var i = 0; i < e.dataTransfer.files.length; i++) {
      let file = e.dataTransfer.files[i];
      fileNames = [...fileNames, `${file.name} (${file.type})`];
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

  $: foo = fileNames;
</script>

<style>
  main {
    height: 100%;
    background: green;
  }
  .dropzone {
    height: 100%;
    background: red;
  }
</style>

<main>
  <div>
    <p>controls</p>
  </div>
  <div
    class="dropzone"
    on:drop={onDrop}
    on:dragenter={onDragEnter}
    on:dragover={onDragOver}>

    <ul>
      {#each fileNames as fileName, i}
        <li>{i} {fileName}</li>
      {/each}
    </ul>
  </div>
</main>
